package application

import (
	"path/filepath"
	"os"
	"github.com/anyproto/anytype-heart/pb"
	oserror "github.com/anyproto/anytype-heart/util/os"
	"github.com/anyproto/anytype-heart/core/anytype/config"
	"github.com/anyproto/anytype-heart/core/filestorage"
	"strings"
	cp "github.com/otiai10/copy"
	"errors"
)

var (
	ErrGetConfig          = errors.New("get config")
	ErrIdentifyAccountDir = errors.New("can't identify account dir")
)

func (s *Service) AccountMove(req *pb.RpcAccountMoveRequest) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	dirs := []string{filestorage.FlatfsDirName}
	conf := s.app.MustComponent(config.CName).(*config.Config)

	configPath := conf.GetConfigPath()
	srcPath := conf.RepoPath
	fileConf := config.ConfigRequired{}
	if err := config.GetFileConfig(configPath, &fileConf); err != nil {
		return errors.Join(ErrGetConfig, err)
	}
	if fileConf.CustomFileStorePath != "" {
		srcPath = fileConf.CustomFileStorePath
	}

	parts := strings.Split(srcPath, string(filepath.Separator))
	accountDir := parts[len(parts)-1]
	if accountDir == "" {
		return ErrIdentifyAccountDir
	}

	destination := filepath.Join(req.NewPath, accountDir)
	if srcPath == destination {
		return errors.Join(ErrFailedToCreateLocalRepo, errors.New("source path should not be equal destination path"))
	}

	if _, err := os.Stat(destination); !os.IsNotExist(err) { // if already exist (in case of the previous fail moving)
		if err := removeDirsRelativeToPath(destination, dirs); err != nil {
			return errors.Join(ErrRemoveAccountData, oserror.TransformError(err))
		}
	}

	err := os.MkdirAll(destination, 0700)
	if err != nil {
		return errors.Join(ErrFailedToCreateLocalRepo, oserror.TransformError(err))
	}

	err = s.stop()
	if err != nil {
		return errors.Join(ErrFailedToStopApplication, err)
	}

	for _, dir := range dirs {
		if _, err := os.Stat(filepath.Join(srcPath, dir)); !os.IsNotExist(err) { // copy only if exist such dir
			if err := cp.Copy(filepath.Join(srcPath, dir), filepath.Join(destination, dir), cp.Options{PreserveOwner: true}); err != nil {
				return errors.Join(ErrFailedToCreateLocalRepo, err)
			}
		}
	}

	err = config.WriteJsonConfig(configPath, config.ConfigRequired{CustomFileStorePath: destination})
	if err != nil {
		return errors.Join(ErrFailedToWriteConfig, err)
	}

	if err := removeDirsRelativeToPath(srcPath, dirs); err != nil {
		return errors.Join(ErrRemoveAccountData, oserror.TransformError(err))
	}

	if srcPath != conf.RepoPath { // remove root account dir, if move not from anytype source dir
		if err := os.RemoveAll(srcPath); err != nil {
			return errors.Join(ErrRemoveAccountData, oserror.TransformError(err))
		}
	}
	return nil
}

func removeDirsRelativeToPath(rootPath string, dirs []string) error {
	for _, dir := range dirs {
		if err := os.RemoveAll(filepath.Join(rootPath, dir)); err != nil {
			return err
		}
	}
	return nil
}
