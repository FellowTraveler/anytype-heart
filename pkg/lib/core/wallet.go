package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/anytypeio/any-sync/util/crypto"
)

var ErrRepoExists = fmt.Errorf("repo not empty, reinitializing would overwrite your account")
var ErrRepoCorrupted = fmt.Errorf("repo is corrupted")

func WalletGenerateMnemonic(wordCount int) (string, error) {
	m, err := crypto.NewMnemonicGenerator().WithWordCount(wordCount)
	if err != nil {
		return "", err
	}
	return string(m), nil
}

func WalletAccountAt(mnemonic string, index int) (crypto.PrivKey, error) {
	return crypto.Mnemonic(mnemonic).DeriveEd25519Key(index)
}

func WalletInitRepo(rootPath string, pk crypto.PrivKey) error {
	devicePriv, _, err := crypto.GenerateRandomEd25519KeyPair()
	if err != nil {
		return err
	}
	repoPath := filepath.Join(rootPath, pk.GetPublic().Account())
	_, err = os.Stat(repoPath)
	if !os.IsNotExist(err) {
		return ErrRepoExists
	}

	os.MkdirAll(repoPath, 0700)
	deviceKeyPath := filepath.Join(repoPath, "device.key")
	proto, err := devicePriv.Marshall()
	if err != nil {
		return err
	}
	encProto, err := pk.GetPublic().Encrypt(proto)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(deviceKeyPath, encProto, 0400)
}
