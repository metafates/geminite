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

func CacheDir() string {
	path := filepath.Join(xdg.CacheHome, "geminite")
	mustMkdir(path)

	return path
}

func BookmarksFile() string {
	dir := filepath.Join(xdg.UserDirs.Documents, "Geminite")
	mustMkdir(dir)

	return filepath.Join(dir, "bookmarks.json")
}
