package debug

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/anytypeio/any-sync/commonspace/object/tree/exporter"
	"github.com/anytypeio/any-sync/commonspace/object/tree/objecttree"
	"github.com/anytypeio/go-anytype-middleware/core/debug/ziparchive"
	"github.com/anytypeio/go-anytype-middleware/util/anonymize"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
)

type treeBuilder struct {
	log        *log.Logger
	s          objectstore.ObjectStore
	anonymized bool
	id         string
	zw         *zip.Writer
}

func (b *treeBuilder) Build(path string, tree objecttree.ReadableObjectTree) (filename string, err error) {
	filename = filepath.Join(path, fmt.Sprintf("at.dbg.%s.%s.zip", b.id, time.Now().Format("20060102.150405.99")))
	exp, err := ziparchive.NewExporter(filename)
	if err != nil {
		return
	}
	defer exp.Close()

	b.zw = exp.Writer()
	params := exporter.TreeExporterParams{
		ListStorageExporter: exp,
		TreeStorageExporter: exp,
		DataConverter:       &changeDataConverter{anonymize: b.anonymized},
	}
	logBuf := bytes.NewBuffer(nil)
	b.log = log.New(io.MultiWriter(logBuf, os.Stderr), "", log.LstdFlags)

	st := time.Now()
	treeExporter := exporter.NewTreeExporter(params)
	b.log.Printf("exporting tree and acl")
	err = treeExporter.ExportUnencrypted(tree)
	if err != nil {
		b.log.Printf("export tree in zip error: %v", err)
		return
	}

	b.log.Printf("exported tree for a %v", time.Since(st))
	data, err := b.s.GetByIDs(b.id)

	if err != nil {
		b.log.Printf("can't fetch localstore info: %v", err)
	} else {
		if len(data) > 0 {
			data[0].Details = anonymize.Struct(data[0].Details)
			data[0].Snippet = anonymize.Text(data[0].Snippet)
			for i, r := range data[0].Relations {
				data[0].Relations[i] = anonymize.Relation(r)
			}
			osData := pbtypes.Sprint(data[0])
			lsWr, er := b.zw.Create("localstore.json")
			if er != nil {
				b.log.Printf("create file in zip error: %v", er)
			} else {
				if _, err := lsWr.Write([]byte(osData)); err != nil {
					b.log.Printf("localstore.json write error: %v", err)
				} else {
					b.log.Printf("localstore.json wrote")
				}
			}
		} else {
			b.log.Printf("not data in objectstore")
		}
	}
	logW, err := b.zw.Create("creation.log")
	if err != nil {
		return
	}
	io.Copy(logW, logBuf)
	return
}
