package where

import (
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/metafates/geminite/afs"
)

func mustMkdir(path string) {
	if err := afs.AFS.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func ConfigDir() string {
	path := filepath.Join(xdg.ConfigHome, "geminite")
	mustMkdir(path)

	return path
}

func ConfigFile() string {
	return filepath.Join(ConfigDir(), "config.toml")
}
