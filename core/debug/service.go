package debug

import (
	"archive/zip"
	"context"
	"fmt"
	"github.com/anytypeio/any-sync/app"
	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/space"
	"github.com/gogo/protobuf/jsonpb"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const CName = "debug"

var logger = logging.Logger("anytype-debug")

func New() Debug {
	return new(debug)
}

type Debug interface {
	app.Component
	DumpTree(blockId, path string, anonymize bool, withSvg bool) (filename string, err error)
	DumpLocalstore(objectIds []string, path string) (filename string, err error)
	SpaceSummary() (summary SpaceSummary, err error)
	TreeHeads(id string) (info TreeInfo, err error)
}

type debug struct {
	block         *block.Service
	store         objectstore.ObjectStore
	clientService space.Service
}

func (d *debug) Init(a *app.App) (err error) {
	d.store = a.MustComponent(objectstore.CName).(objectstore.ObjectStore)
	d.clientService = a.MustComponent(space.CName).(space.Service)
	d.block = a.MustComponent(block.CName).(*block.Service)
	return nil
}

func (d *debug) Name() (name string) {
	return CName
}

type TreeInfo struct {
	Heads   []string
	Id      string
	SpaceId string
}

type SpaceSummary struct {
	SpaceId   string
	TreeInfos []TreeInfo
}

func (d *debug) SpaceSummary() (summary SpaceSummary, err error) {
	spc, err := d.clientService.AccountSpace(context.Background())
	if err != nil {
		return
	}
	var (
		debugTrees = spc.DebugAllHeads()
		treeInfos  = make([]TreeInfo, 0, len(debugTrees))
	)
	for _, t := range debugTrees {
		treeInfos = append(treeInfos, TreeInfo{
			Heads: t.Heads,
			Id:    t.Id,
		})
	}
	summary.SpaceId = spc.Id()
	return
}

func (d *debug) TreeHeads(id string) (info TreeInfo, err error) {
	spc, err := d.clientService.AccountSpace(context.Background())
	if err != nil {
		return
	}
	tree, err := d.block.GetTree(context.Background(), spc.Id(), id)
	if err != nil {
		return
	}
	info = TreeInfo{
		Id:      id,
		Heads:   tree.Heads(),
		SpaceId: spc.Id(),
	}
	return
}

func (d *debug) DumpTree(blockId, path string, anonymize bool, withSvg bool) (filename string, err error) {
	// 0 - get first block
	tree, err := d.block.GetAccountTree(context.Background(), blockId)
	if err != nil {
		return
	}

	// 1 - create ZIP file
	// <path>/at.dbg.bafkudtugh626rrqzah3kam4yj4lqbaw4bjayn2rz4ah4n5fpayppbvmq.20220322.121049.23.zip
	exporter := &treeExporter{s: d.store, anonymized: anonymize, id: blockId}
	zipFilename, err := exporter.Export(path, tree)
	if err != nil {
		logger.Fatal("build tree error:", err)
		return "", err
	}

	// if client never asked for SVG generation -> return
	if !withSvg {
		return zipFilename, err
	}

	// 2 - create SVG file near ZIP
	// <path>/at.dbg.bafkudtugh626rrqzah3kam4yj4lqbaw4bjayn2rz4ah4n5fpayppbvmq.20220322.121049.23.svg
	//
	// this will return "graphviz is not supported on the current platform" error if no graphviz!
	// generate a filename just like zip file had
	maxReplacements := 1
	svgFilename := strings.Replace(zipFilename, ".zip", ".svg", maxReplacements)
	debugInfo, err := tree.Debug(state.ChangeParser{})
	if err != nil {
		return
	}

	file, err := os.Create(svgFilename)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.Write([]byte(debugInfo.Graphviz))
	if err != nil {
		return
	}

	// return zip filename, but not svgFilename
	return zipFilename, nil
}

func (d *debug) DumpLocalstore(objIds []string, path string) (filename string, err error) {
	if len(objIds) == 0 {
		objIds, err = d.store.ListIds()
		if err != nil {
			return "", err
		}
	}

	filename = filepath.Join(path, fmt.Sprintf("at.store.dbg.%s.zip", time.Now().Format("20060102.150405.99")))
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	zw := zip.NewWriter(f)
	defer zw.Close()

	var wr io.Writer
	m := jsonpb.Marshaler{Indent: " "}

	for _, objId := range objIds {
		doc, err := d.store.GetWithLinksInfoByID(objId)
		if err != nil {
			var err2 error
			wr, err2 = zw.Create(fmt.Sprintf("%s.txt", objId))
			if err2 != nil {
				return "", err
			}

			wr.Write([]byte(err.Error()))
			continue
		}
		wr, err = zw.Create(fmt.Sprintf("%s.json", objId))
		if err != nil {
			return "", err
		}

		err = m.Marshal(wr, doc)
		if err != nil {
			return "", err
		}
	}
	return filename, nil
}
