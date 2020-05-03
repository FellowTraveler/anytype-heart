package _import

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-middleware/anymark"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/gogo/protobuf/types"
	"github.com/google/uuid"
)

var (
	linkRegexp = regexp.MustCompile(`\[([\s\S]*?)\]\((.*?)\)`)
)

type Import interface {
	ImportMarkdown(ctx *state.Context, req pb.RpcBlockImportMarkdownRequest) (rootLinks []*model.Block, err error)
}

func NewImport(sb smartblock.SmartBlock, ctrl Services) Import {
	return &importImpl{SmartBlock: sb, ctrl: ctrl}
}

type importImpl struct {
	smartblock.SmartBlock
	ctrl Services
}

type Services interface {
	CreateSmartBlock(req pb.RpcBlockCreatePageRequest) (pageId string, err error)
	Paste(ctx *state.Context, req pb.RpcBlockPasteRequest) (blockIds []string, uploadArr []pb.RpcBlockUploadRequest, caretPosition int32, err error)
	UploadBlockFile(ctx *state.Context, req pb.RpcBlockUploadRequest) error
}

func (imp *importImpl) ImportMarkdown(ctx *state.Context, req pb.RpcBlockImportMarkdownRequest) (rootLinks []*model.Block, err error) {
	s := imp.NewStateCtx(ctx)

	nameToBlocks, isPageLinked, filesCount, err := imp.DirWithMarkdownToBlocks(req.ImportPath)
	nameToId := make(map[string]string)

	for name := range nameToBlocks {
		fileName := strings.Replace(filepath.Base(name), ".md", "", -1)

		if len(nameToBlocks[name]) > 0 && nameToBlocks[name][0].GetText() != nil &&
			nameToBlocks[name][0].GetText().Text == fileName {
			nameToBlocks[name] = nameToBlocks[name][1:]
		}

		nameToId[name], err = imp.ctrl.CreateSmartBlock(pb.RpcBlockCreatePageRequest{
			Details: &types.Struct{
				Fields: map[string]*types.Value{
					"name": pbtypes.String(fileName),
				},
			},
		})

		if err != nil {
			return rootLinks, err
		}
	}

	for name := range nameToBlocks {
		for i := range nameToBlocks[name] {
			if link := nameToBlocks[name][i].GetLink(); link != nil && len(nameToId[name]) > 0 {
				link.TargetBlockId = nameToId[strings.Replace(link.TargetBlockId, "%20", " ", -1)]
			}
		}
	}

	for name := range nameToBlocks {
		if len(nameToBlocks[name]) > 0 {
			_, _, _, err = imp.ctrl.Paste(ctx, pb.RpcBlockPasteRequest{
				ContextId: nameToId[name],
				AnySlot:   nameToBlocks[name],
			})
		}

		if err != nil {
			return rootLinks, err
		}

		if err != nil {
			return rootLinks, err
		}
	}

	for name := range nameToBlocks {
		if !isPageLinked[name] {
			rootLinks = append(rootLinks, &model.Block{
				Content: &model.BlockContentOfLink{
					Link: &model.BlockContentLink{
						TargetBlockId: nameToId[name],
						Style:         model.BlockContentLink_Page,
						Fields:        nil,
					},
				},
			})
		}

		for _, b := range nameToBlocks[name] {
			if f := b.GetFile(); f != nil {

				filesCount = filesCount - 1

				err = imp.ctrl.UploadBlockFile(ctx, pb.RpcBlockUploadRequest{
					ContextId: nameToId[name],
					BlockId:   b.Id,
					FilePath:  f.Name,
					Url:       "",
				})

				if err != nil {
					return rootLinks, fmt.Errorf("can not import this file: %s. error: %s", f.Name, err)
				}
			}
		}
	}

	return rootLinks, imp.Apply(s)
}

func (imp *importImpl) DirWithMarkdownToBlocks(directoryPath string) (nameToBlock map[string][]*model.Block, isPageLinked map[string]bool, filesCount int, err error) {
	anymarkConv := anymark.New()

	nameToBlocks := make(map[string][]*model.Block)
	allFileShortPaths := []string{}

	err = filepath.Walk(directoryPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				shortPath := strings.Replace(path, directoryPath+"/", "", -1)
				allFileShortPaths = append(allFileShortPaths, shortPath)
			}

			return nil
		},
	)

	if err != nil {
		return nameToBlocks, isPageLinked, filesCount, err
	}

	isFileExist := make(map[string]bool)

	err = filepath.Walk(directoryPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				extension := filepath.Ext(path)

				dat, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}

				shortPath := strings.Replace(path, directoryPath+"/", "", -1)

				if extension == ".md" {
					datStr := string(dat)
					linkSubmatches := linkRegexp.FindAllStringSubmatch(datStr, -1)

					for _, linkSubmatch := range linkSubmatches {
						l := strings.Replace(linkSubmatch[2], "%20", " ", -1)

						for _, sPath := range allFileShortPaths {
							if strings.Contains(sPath, l) {
								datStr = strings.Replace(datStr, linkSubmatch[2], strings.Replace(sPath, " ", "%20", -1), -1)
							}
						}
					}

					nameToBlocks[shortPath], err = anymarkConv.MarkdownToBlocks([]byte(datStr))
					if err != nil {
						return err
					}
				} else {
					isFileExist[shortPath] = true
				}

			}

			return nil
		})

	isPageLinked = make(map[string]bool)
	for name, _ := range nameToBlocks {
		for i, block := range nameToBlocks[name] {
			nameToBlocks[name][i].Id = uuid.New().String()

			txt := block.GetText()
			if txt != nil && txt.Marks != nil && len(txt.Marks.Marks) == 1 &&
				txt.Marks.Marks[0].Type == model.BlockContentTextMark_Link {

				linkConverted := strings.Replace(txt.Marks.Marks[0].Param, "%20", " ", -1)

				if nameToBlocks[linkConverted] != nil {
					nameToBlocks[name][i], isPageLinked = imp.convertTextToPageLink(block, isPageLinked)
				}

				if isFileExist[linkConverted] {
					nameToBlocks[name][i] = imp.convertTextToFile(block, directoryPath)
				}
			}

			if block.GetFile() != nil {
				block.GetFile().Name = strings.Replace(directoryPath+"/"+block.GetFile().Name, "%20", " ", -1)
				block.GetFile().Type = model.BlockContentFile_Image
			}
		}
	}

	return nameToBlocks, isPageLinked, len(isFileExist), err
}

func (imp *importImpl) convertTextToPageLink(block *model.Block, isPageLinked map[string]bool) (*model.Block, map[string]bool) {
	targetId := strings.Replace(block.GetText().Marks.Marks[0].Param, "%20", " ", -1)
	blockOut := &model.Block{
		Content: &model.BlockContentOfLink{
			Link: &model.BlockContentLink{
				TargetBlockId: targetId,
				Style:         model.BlockContentLink_Page,
			},
		},
	}

	isPageLinked[targetId] = true
	return blockOut, isPageLinked
}

func (imp *importImpl) convertTextToFile(block *model.Block, importPath string) *model.Block {
	fName := strings.Replace(importPath+"/"+block.GetText().Marks.Marks[0].Param, "%20", " ", -1)

	imageFormats := []string{"jpg", "jpeg", "png", "gif", "webp"} // "svg" excluded
	videoFormats := []string{"mp4", "m4v"}

	fileType := model.BlockContentFile_File
	for _, ext := range imageFormats {
		if filepath.Ext(fName)[1:] == ext {
			fileType = model.BlockContentFile_Image
		}
	}

	for _, ext := range videoFormats {
		if filepath.Ext(fName)[1:] == ext {
			fileType = model.BlockContentFile_Video
		}
	}

	blockOut := &model.Block{
		Id: block.Id,
		Content: &model.BlockContentOfFile{
			File: &model.BlockContentFile{
				Name:  fName,
				State: model.BlockContentFile_Empty,
				Type:  fileType,
			},
		},
	}

	return blockOut
}
