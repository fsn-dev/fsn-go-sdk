package utils

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/fsn-dev/fsn-go-sdk/efsn/accounts"
	"github.com/fsn-dev/fsn-go-sdk/efsn/accounts/keystore"
	"gopkg.in/urfave/cli.v1"
)

var (
	datadirDefaultKeyStore = "keystore"
)

type Config struct {
	DataDir     string
	KeyStoreDir string
}

// SetNodeConfig applies node-related command line flags to the config.
func SetNodeConfig(ctx *cli.Context, cfg *Config) {
	cfg.DataDir = ctx.String(DataDirFlag.Name)
	cfg.KeyStoreDir = ctx.String(KeyStoreDirFlag.Name)
}

// AccountConfig determines the settings for scrypt and keydirectory
func (c *Config) AccountConfig() (int, int, string, error) {
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP

	var (
		keydir string
		err    error
	)
	switch {
	case filepath.IsAbs(c.KeyStoreDir):
		keydir = c.KeyStoreDir
	case c.DataDir != "":
		if c.KeyStoreDir == "" {
			keydir = filepath.Join(c.DataDir, datadirDefaultKeyStore)
		} else {
			keydir, err = filepath.Abs(c.KeyStoreDir)
		}
	case c.KeyStoreDir != "":
		keydir, err = filepath.Abs(c.KeyStoreDir)
	}
	return scryptN, scryptP, keydir, err
}

func (c *Config) AccountManager() (*accounts.Manager, string, error) {
	scryptN, scryptP, keydir, err := c.AccountConfig()
	if err != nil {
		return nil, "", err
	}
	var ephemeral string
	if keydir == "" {
		return nil, "", errors.New("No keystore dir is specified")
	}
	if err := os.MkdirAll(keydir, 0700); err != nil {
		return nil, "", err
	}
	// Assemble the account manager and supported backends
	backends := []accounts.Backend{
		keystore.NewKeyStore(keydir, scryptN, scryptP),
	}
	return accounts.NewManager(backends...), ephemeral, nil
}
