package files

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/cafe"
	cafepb "github.com/anytypeio/go-anytype-middleware/pkg/lib/cafe/pb"
	symmetric "github.com/anytypeio/go-anytype-middleware/pkg/lib/crypto/symmetric"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/crypto/symmetric/cfb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/crypto/symmetric/gcm"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/ipfs"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/ipfs/helpers"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	m "github.com/anytypeio/go-anytype-middleware/pkg/lib/mill"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/mill/schema"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/mill/schema/anytype"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/storage"
	"github.com/gogo/protobuf/proto"
	"github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
	uio "github.com/ipfs/go-unixfs/io"
	"github.com/multiformats/go-base32"
	mh "github.com/multiformats/go-multihash"
)

var log = logging.Logger("anytype-files")

type Service struct {
	store localstore.FileStore
	ipfs  ipfs.IPFS
	cafe  cafe.Client
}

func New(store localstore.FileStore, ipfs ipfs.IPFS, cafe cafe.Client) *Service {
	return &Service{
		store: store,
		ipfs:  ipfs,
		cafe:  cafe,
	}
}

var ErrMissingMetaLink = fmt.Errorf("meta link not in node")
var ErrMissingContentLink = fmt.Errorf("content link not in node")

const MetaLinkName = "meta"
const ContentLinkName = "content"

var ValidMetaLinkNames = []string{"meta"}
var ValidContentLinkNames = []string{"content"}

var cidBuilder = cid.V1Builder{Codec: cid.DagProtobuf, MhType: mh.SHA2_256}

func (s *Service) FileAdd(ctx context.Context, opts AddOptions) (string, *storage.FileInfo, error) {
	fileInfo, err := s.FileAddWithConfig(ctx, &m.Blob{}, opts)
	if err != nil {
		return "", nil, err
	}

	node, keys, err := s.fileAddNodeFromFiles(ctx, []*storage.FileInfo{fileInfo})
	if err != nil {
		return "", nil, err
	}

	nodeHash := node.Cid().String()

	if err = s.fileIndexData(ctx, node, nodeHash); err != nil {
		return "", nil, err
	}

	if err = s.store.AddFileKeys(localstore.FileKeys{
		Hash: nodeHash,
		Keys: keys.KeysByPath,
	}); err != nil {
		return "", nil, err
	}

	if s.cafe != nil {
		go func() {
			for i := 0; i <= 10; i++ {
				_, err := s.cafe.FilePin(context.Background(), &cafepb.FilePinRequest{Cid: nodeHash})
				if err != nil {
					log.Errorf("failed to pin file %s on the cafe: %s", nodeHash, err.Error())
					time.Sleep(time.Minute * time.Duration(i+1))
					continue
				}

				log.Debugf("pinning file %s started on the cafe", nodeHash)
				break
			}
		}()
	}

	return nodeHash, fileInfo, nil
}

// fileRestoreKeys restores file path=>key map from the IPFS DAG using the keys in the localStore
func (s *Service) FileRestoreKeys(ctx context.Context, hash string) (map[string]string, error) {
	links, err := helpers.LinksAtCid(ctx, s.ipfs, hash)
	if err != nil {
		return nil, err
	}

	var fileKeys = make(map[string]string)
	for _, index := range links {
		node, err := helpers.NodeAtLink(ctx, s.ipfs, index)
		if err != nil {
			return nil, err
		}

		if looksLikeFileNode(node) {
			l := schema.LinkByName(node.Links(), ValidContentLinkNames)
			info, err := s.store.GetByHash(l.Cid.String())
			if err == nil {
				fileKeys["/"+index.Name+"/"] = info.Key
			} else {
				log.Warnf("fileRestoreKeys not found in db %s(%s)", node.Cid().String(), hash+"/"+index.Name)
			}
		} else {
			for _, link := range node.Links() {
				innerLinks, err := helpers.LinksAtCid(ctx, s.ipfs, link.Cid.String())
				if err != nil {
					return nil, err
				}

				l := schema.LinkByName(innerLinks, ValidContentLinkNames)
				if l == nil {
					log.Errorf("con")
					continue
				}

				info, err := s.store.GetByHash(l.Cid.String())

				if err == nil {
					fileKeys["/"+index.Name+"/"+link.Name+"/"] = info.Key
				} else {
					log.Warnf("fileRestoreKeys not found in db %s(%s)", node.Cid().String(), "/"+index.Name+"/"+link.Name+"/")
				}
			}
		}
	}

	err = s.store.AddFileKeys(localstore.FileKeys{
		Hash: hash,
		Keys: fileKeys,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to save file keys: %w", err)
	}

	return fileKeys, nil
}

func (s *Service) fileAddNodeFromDirs(ctx context.Context, dirs *storage.DirectoryList) (ipld.Node, *storage.FileKeys, error) {
	keys := &storage.FileKeys{KeysByPath: make(map[string]string)}
	outer := uio.NewDirectory(s.ipfs)
	outer.SetCidBuilder(cidBuilder)

	for i, dir := range dirs.Items {
		inner := uio.NewDirectory(s.ipfs)
		inner.SetCidBuilder(cidBuilder)
		olink := strconv.Itoa(i)

		var err error
		for link, file := range dir.Files {
			err = s.fileNode(ctx, file, inner, link)
			if err != nil {
				return nil, nil, err
			}
			keys.KeysByPath["/"+olink+"/"+link+"/"] = file.Key
		}

		node, err := inner.GetNode()
		if err != nil {
			return nil, nil, err
		}
		// todo: pin?
		err = s.ipfs.Add(ctx, node)
		if err != nil {
			return nil, nil, err
		}

		id := node.Cid().String()
		err = helpers.AddLinkToDirectory(ctx, s.ipfs, outer, olink, id)
		if err != nil {
			return nil, nil, err
		}
	}

	node, err := outer.GetNode()
	if err != nil {
		return nil, nil, err
	}
	// todo: pin?
	err = s.ipfs.Add(ctx, node)
	if err != nil {
		return nil, nil, err
	}
	return node, keys, nil
}

func (s *Service) fileAddNodeFromFiles(ctx context.Context, files []*storage.FileInfo) (ipld.Node, *storage.FileKeys, error) {
	keys := &storage.FileKeys{KeysByPath: make(map[string]string)}
	outer := uio.NewDirectory(s.ipfs)
	outer.SetCidBuilder(cidBuilder)

	var err error
	for i, file := range files {
		link := strconv.Itoa(i)
		err = s.fileNode(ctx, file, outer, link)
		if err != nil {
			return nil, nil, err
		}
		keys.KeysByPath["/"+link+"/"] = file.Key
	}

	node, err := outer.GetNode()
	if err != nil {
		return nil, nil, err
	}

	err = s.ipfs.Add(ctx, node)
	if err != nil {
		return nil, nil, err
	}

	/*err = helpers.PinNode(s.node, node, false)
	if err != nil {
		return nil, nil, err
	}*/
	return node, keys, nil
}

func (s *Service) FileGetInfoForPath(pth string) (*storage.FileInfo, error) {
	return nil, fmt.Errorf("not implemented")
}

// fileIndexData walks a file data node, indexing file links
func (s *Service) fileIndexData(ctx context.Context, inode ipld.Node, data string) error {
	for _, link := range inode.Links() {
		nd, err := helpers.NodeAtLink(ctx, s.ipfs, link)
		if err != nil {
			return err
		}
		err = s.fileIndexNode(ctx, nd, data)
		if err != nil {
			return err
		}
	}

	return nil
}

// fileIndexNode walks a file node, indexing file links
func (s *Service) fileIndexNode(ctx context.Context, inode ipld.Node, data string) error {
	links := inode.Links()

	if looksLikeFileNode(inode) {
		return s.fileIndexLink(ctx, inode, data)
	}

	for _, link := range links {
		n, err := helpers.NodeAtLink(ctx, s.ipfs, link)
		if err != nil {
			return err
		}

		err = s.fileIndexLink(ctx, n, data)
		if err != nil {
			return err
		}
	}

	return nil
}

// fileIndexLink indexes a file link
func (s *Service) fileIndexLink(ctx context.Context, inode ipld.Node, data string) error {
	dlink := schema.LinkByName(inode.Links(), ValidContentLinkNames)
	if dlink == nil {
		return ErrMissingContentLink
	}

	return s.store.AddTarget(dlink.Cid.String(), data)
}

func (s *Service) fileInfoFromPath(target string, path string, key string) (*storage.FileInfo, error) {
	r, err := helpers.DataAtPath(context.TODO(), s.ipfs, path+"/"+MetaLinkName)
	if err != nil {
		return nil, err
	}

	var file storage.FileInfo

	if key != "" {
		key, err := symmetric.FromString(key)
		if err != nil {
			return nil, err
		}

		modes := []storage.FileInfoEncryptionMode{storage.FileInfo_AES_CFB, storage.FileInfo_AES_GCM}
		for i, mode := range modes {
			if i > 0 {
				_, err = r.Seek(0, io.SeekStart)
				if err != nil {
					return nil, fmt.Errorf("failed to seek ciphertext after enc mode try")
				}
			}
			ed, err := getEncryptorDecryptor(key, mode)
			if err != nil {
				return nil, err
			}
			decryptedReader, err := ed.DecryptReader(r)
			if err != nil {
				return nil, err
			}
			b, err := ioutil.ReadAll(decryptedReader)
			if err != nil {
				if i == len(modes)-1 {
					return nil, fmt.Errorf("failed to unmarshal file info proto with all encryption modes: %w", err)
				}

				continue
			}
			err = proto.Unmarshal(b, &file)
			if err != nil || file.Hash == "" {
				if i == len(modes)-1 {
					return nil, fmt.Errorf("failed to unmarshal file info proto with all encryption modes: %w", err)
				}
				continue
			}
			// save successful enc mode so it will be cached in the DB
			file.EncMode = mode
			break
		}
	}
	if file.Hash == "" {
		return nil, fmt.Errorf("failed to read file info proto with all encryption modes")
	}

	file.Targets = []string{target}
	return &file, nil
}

func (s *Service) fileContent(ctx context.Context, hash string) (io.ReadSeeker, *storage.FileInfo, error) {
	var err error
	var file *storage.FileInfo
	var reader io.ReadSeeker
	file, err = s.store.GetByHash(hash)
	if err != nil {
		return nil, nil, err
	}
	reader, err = s.FileContentReader(ctx, file)
	return reader, file, err
}

func (s *Service) FileContentReader(ctx context.Context, file *storage.FileInfo) (symmetric.ReadSeekCloser, error) {
	fileCid, err := cid.Parse(file.Hash)
	if err != nil {
		return nil, err
	}
	fd, err := s.ipfs.GetFile(ctx, fileCid)
	if err != nil {
		return nil, err
	}
	if file.Key == "" {
		return fd, nil
	}

	key, err := symmetric.FromString(file.Key)
	if err != nil {
		return nil, err
	}

	dec, err := getEncryptorDecryptor(key, file.EncMode)
	if err != nil {
		return nil, err
	}

	return dec.DecryptReader(fd)
}

func (s *Service) FileAddWithConfig(ctx context.Context, mill m.Mill, conf AddOptions) (*storage.FileInfo, error) {
	var source string

	if conf.Use != "" {
		source = conf.Use
	} else {
		var err error
		source, err = checksum(conf.Reader, conf.Plaintext)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate checksum: %w", err)
		}
		_, err = conf.Reader.Seek(0, io.SeekStart)
		if err != nil {
			return nil, fmt.Errorf("failed to seek reader: %w", err)
		}
	}

	opts, err := mill.Options(map[string]interface{}{
		"plaintext": conf.Plaintext,
	})
	if err != nil {
		return nil, err
	}

	if efile, _ := s.store.GetBySource(mill.ID(), source, opts); efile != nil {
		return efile, nil
	}

	res, err := mill.Mill(conf.Reader, conf.Name)
	if err != nil {
		return nil, err
	}

	check, err := checksum(res.File, conf.Plaintext)
	if err != nil {
		return nil, err
	}
	conf.Reader.Seek(0, io.SeekStart)

	// todo: temp solution to read the same data again
	res, err = mill.Mill(conf.Reader, conf.Name)
	if err != nil {
		return nil, err
	}

	if efile, _ := s.store.GetByChecksum(mill.ID(), check); efile != nil {
		return efile, nil
	}
	model := &storage.FileInfo{
		Mill:     mill.ID(),
		Checksum: check,
		Source:   source,
		Opts:     opts,
		Media:    conf.Media,
		Name:     conf.Name,
		Added:    time.Now().Unix(),
		Meta:     pb.ToStruct(res.Meta),
	}

	var r io.Reader
	if mill.Encrypt() && !conf.Plaintext {
		key, err := symmetric.NewRandom()
		if err != nil {
			return nil, err
		}
		enc := cfb.New(key, [aes.BlockSize]byte{})

		r, err = enc.EncryptReader(res.File)
		if err != nil {
			return nil, err
		}

		model.Key = key.String()
		model.EncMode = storage.FileInfo_AES_CFB
	} else {
		r = res.File
	}

	node, err := s.ipfs.AddFile(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	stat, err := node.Stat()
	if err != nil {
		return nil, err
	}

	model.Hash = node.Cid().String()
	model.Size_ = int64(stat.CumulativeSize)
	err = s.store.Add(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) fileNode(ctx context.Context, file *storage.FileInfo, dir uio.Directory, link string) error {
	file, err := s.store.GetByHash(file.Hash)
	if err != nil {
		return err
	}

	// remove locally indexed targets
	file.Targets = nil

	plaintext, err := proto.Marshal(file)
	if err != nil {
		return err
	}

	var reader io.Reader
	if file.Key != "" {
		key, err := symmetric.FromString(file.Key)
		if err != nil {
			return err
		}

		ed, err := getEncryptorDecryptor(key, file.EncMode)
		if err != nil {
			return err
		}

		reader, err = ed.EncryptReader(bytes.NewReader(plaintext))
		if err != nil {
			return err
		}
	} else {
		reader = bytes.NewReader(plaintext)
	}

	pair := uio.NewDirectory(s.ipfs)
	pair.SetCidBuilder(cidBuilder)

	_, err = helpers.AddDataToDirectory(ctx, s.ipfs, pair, MetaLinkName, reader)
	if err != nil {
		return err
	}

	err = helpers.AddLinkToDirectory(ctx, s.ipfs, pair, ContentLinkName, file.Hash)
	if err != nil {
		return err
	}

	node, err := pair.GetNode()
	if err != nil {
		return err
	}
	err = s.ipfs.Add(ctx, node)
	if err != nil {
		return err
	}

	/*err = helpers.PinNode(s.node, node, false)
	if err != nil {
		return err
	}*/

	return helpers.AddLinkToDirectory(ctx, s.ipfs, dir, link, node.Cid().String())
}

func (s *Service) fileBuildDirectory(ctx context.Context, content []byte, filename string, plaintext bool, sch *storage.Node) (*storage.Directory, error) {
	dir := &storage.Directory{
		Files: make(map[string]*storage.FileInfo),
	}

	reader := bytes.NewReader(content)
	mil, err := anytype.GetMill(sch.Mill, sch.Opts)
	if err != nil {
		return nil, err
	}
	if mil != nil {
		opts := AddOptions{
			Reader:    reader,
			Use:       "",
			Media:     "",
			Name:      filename,
			Plaintext: sch.Plaintext || plaintext,
		}
		err := s.NormalizeOptions(ctx, &opts)
		if err != nil {
			return nil, err
		}

		added, err := s.FileAddWithConfig(ctx, mil, opts)
		if err != nil {
			return nil, err
		}
		dir.Files[schema.SingleFileTag] = added

	} else if len(sch.Links) > 0 {
		// determine order
		steps, err := schema.Steps(sch.Links)
		if err != nil {
			return nil, err
		}

		// send each link
		for _, step := range steps {
			stepMill, err := anytype.GetMill(step.Link.Mill, step.Link.Opts)
			if err != nil {
				return nil, err
			}
			var opts *AddOptions
			if step.Link.Use == schema.FileTag {
				opts = &AddOptions{
					Reader:    reader,
					Use:       "",
					Media:     "",
					Name:      filename,
					Plaintext: step.Link.Plaintext || plaintext,
				}
				err = s.NormalizeOptions(ctx, opts)
				if err != nil {
					return nil, err
				}

			} else {
				if dir.Files[step.Link.Use] == nil {
					return nil, fmt.Errorf(step.Link.Use + " not found")
				}

				opts = &AddOptions{
					Reader:    nil,
					Use:       dir.Files[step.Link.Use].Hash,
					Media:     "",
					Name:      filename,
					Plaintext: step.Link.Plaintext || plaintext,
				}

				err = s.NormalizeOptions(ctx, opts)
				if err != nil {
					return nil, err
				}
			}

			added, err := s.FileAddWithConfig(ctx, stepMill, *opts)
			if err != nil {
				return nil, err
			}
			dir.Files[step.Name] = added
			reader.Seek(0, 0)
		}
	} else {
		return nil, schema.ErrEmptySchema
	}

	return dir, nil
}

func (s *Service) FileIndexInfo(ctx context.Context, hash string) ([]*storage.FileInfo, error) {
	links, err := helpers.LinksAtCid(ctx, s.ipfs, hash)
	if err != nil {
		return nil, err
	}

	keys, err := s.store.GetFileKeys(hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get file keys from cache: %w", err)
	}

	var files []*storage.FileInfo
	for _, index := range links {
		node, err := helpers.NodeAtLink(ctx, s.ipfs, index)
		if err != nil {
			return nil, err
		}

		if looksLikeFileNode(node) {
			var key string
			if keys != nil {
				key = keys["/"+index.Name+"/"]
			}

			fileIndex, err := s.fileInfoFromPath(hash, hash+"/"+index.Name, key)
			if err != nil {
				return nil, fmt.Errorf("fileInfoFromPath error: %s", err.Error())
			}
			files = append(files, fileIndex)
		} else {
			for _, link := range node.Links() {
				var key string
				if keys != nil {
					key = keys["/"+index.Name+"/"+link.Name+"/"]
				}

				fileIndex, err := s.fileInfoFromPath(hash, hash+"/"+index.Name+"/"+link.Name, key)
				if err != nil {
					return nil, fmt.Errorf("fileInfoFromPath error: %s", err.Error())
				}
				files = append(files, fileIndex)
			}
		}
	}

	err = s.store.AddMulti(files...)
	if err != nil {
		return nil, fmt.Errorf("failed to add files to store: %w", err)
	}

	return files, nil
}

// looksLikeFileNode returns whether or not a node appears to
// be a textile node. It doesn't inspect the actual data.
func looksLikeFileNode(node ipld.Node) bool {
	links := node.Links()
	if len(links) != 2 {
		return false
	}
	if schema.LinkByName(links, ValidMetaLinkNames) == nil ||
		schema.LinkByName(links, ValidContentLinkNames) == nil {
		return false
	}
	return true
}

func checksum(r io.Reader, wontEncrypt bool) (string, error) {
	var add int
	if wontEncrypt {
		add = 1
	}
	h := sha256.New()
	_, err := io.Copy(h, r)
	if err != nil {
		return "", err
	}

	_, err = h.Write([]byte{byte(add)})
	if err != nil {
		return "", err
	}
	checksum := h.Sum(nil)
	return base32.RawHexEncoding.EncodeToString(checksum[:]), nil
}

func getEncryptorDecryptor(key symmetric.Key, mode storage.FileInfoEncryptionMode) (symmetric.EncryptorDecryptor, error) {
	switch mode {
	case storage.FileInfo_AES_GCM:
		return gcm.New(key), nil
	case storage.FileInfo_AES_CFB:
		return cfb.New(key, [aes.BlockSize]byte{}), nil
	default:
		return nil, fmt.Errorf("unsupported encryption mode")
	}
}
