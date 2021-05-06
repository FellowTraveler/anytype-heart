package core

import (
	"context"
	"fmt"
	"io"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/files"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/filestore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
)

var ErrFileNotFound = fmt.Errorf("file not found")

func (a *Anytype) FileGetKeys(hash string) (*files.FileKeys, error) {
	return a.files.FileGetKeys(hash)
}

func (a *Anytype) FileStoreKeys(fileKeys ...files.FileKeys) error {
	var fks []filestore.FileKeys

	for _, fk := range fileKeys {
		fks = append(fks, filestore.FileKeys{
			Hash: fk.Hash,
			Keys: fk.Keys,
		})
	}

	return a.fileStore.AddFileKeys(fks...)
}

func (a *Anytype) FileByHash(ctx context.Context, hash string) (File, error) {
	fileList, err := a.fileStore.ListByTarget(hash)
	if err != nil {
		return nil, err
	}

	if len(fileList) == 0 || fileList[0].MetaHash == "" {
		// info from ipfs
		fileList, err = a.files.FileIndexInfo(ctx, hash, false)
		if err != nil {
			log.Errorf("FileByHash: failed to retrieve from IPFS: %s", err.Error())
			return nil, ErrFileNotFound
		}
	}

	fileIndex := fileList[0]
	return &file{
		hash: hash,
		info: fileIndex,
		node: a.files,
	}, nil
}

func (a *Anytype) FileAdd(ctx context.Context, options ...files.AddOption) (File, error) {
	opts := files.AddOptions{}
	for _, opt := range options {
		opt(&opts)
	}

	err := a.files.NormalizeOptions(ctx, &opts)
	if err != nil {
		return nil, err
	}

	hash, info, err := a.files.FileAdd(ctx, opts)
	if err != nil {
		return nil, err
	}

	f := &file{
		hash: hash,
		info: info,
		node: a.files,
	}

	details, err := f.Details()
	if err != nil {
		return nil, err
	}

	err = a.objectStore.UpdateObjectDetails(f.hash, details, &model.Relations{Relations: bundle.MustGetType(bundle.TypeKeyFile).Relations})
	if err != nil {
		return nil, err
	}
	err = a.objectStore.AddToIndexQueue(f.hash)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (a *Anytype) FileAddWithReader(ctx context.Context, content io.ReadSeeker, filename string) (File, error) {
	return a.FileAdd(ctx, files.WithReader(content), files.WithName(filename))
}

func (a *Anytype) FileAddWithBytes(ctx context.Context, content []byte, filename string) (File, error) {
	return a.FileAdd(ctx, files.WithBytes(content), files.WithName(filename))
}
