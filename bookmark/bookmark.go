package bookmark

import (
	"encoding/json"
	"io"
	"net/url"
	"os"
	"sort"
	"sync"

	"github.com/metafates/geminite/afs"
	"github.com/metafates/geminite/where"
	"github.com/spf13/afero"
)

var mu sync.Mutex

type Bookmark struct {
	URL   *url.URL
	Title string
}

func Save(bookmark Bookmark) error {
	raw, err := bookmarksRaw()
	if err != nil {
		return err
	}

	raw[bookmark.URL.String()] = bookmark
	return writeBookmarks(raw)
}

func Delete(URL *url.URL) error {
	raw, err := bookmarksRaw()
	if err != nil {
		return err
	}

	delete(raw, URL.String())

	return writeBookmarks(raw)
}

func GetAll() ([]Bookmark, error) {
	raw, err := bookmarksRaw()
	if err != nil {
		return nil, err
	}

	bookmarks := make([]Bookmark, 0, len(raw))

	for _, value := range raw {
		bookmarks = append(bookmarks, value)
	}

	sort.Slice(bookmarks, func(i, j int) bool {
		return bookmarks[i].Title < bookmarks[j].Title
	})

	return bookmarks, nil
}

func writeBookmarks(bookmarks map[string]Bookmark) error {
	mu.Lock()
	defer mu.Unlock()

	file, err := afs.AFS.OpenFile(where.BookmarksFile(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(bookmarks)
}

func bookmarksRaw() (map[string]Bookmark, error) {
	mu.Lock()
	defer mu.Unlock()

	exists, err := afs.AFS.Exists(where.BookmarksFile())
	if err != nil {
		return nil, err
	}

	bookmarks := make(map[string]Bookmark)

	var bookmarksFile afero.File
	if !exists {
		bookmarksFile, err = afs.AFS.Create(where.BookmarksFile())
		if err != nil {
			return nil, err
		}

		if err = json.NewEncoder(bookmarksFile).Encode(bookmarks); err != nil {
			return nil, err
		}

		// rewind back
		_, err = bookmarksFile.Seek(0, io.SeekStart)
		if err != nil {
			return nil, err
		}
	} else {
		bookmarksFile, err = afs.AFS.Open(where.BookmarksFile())
		if err != nil {
			return nil, err
		}
	}
	defer bookmarksFile.Close()

	if err = json.NewDecoder(bookmarksFile).Decode(&bookmarks); err != nil {
		return nil, err
	}

	return bookmarks, nil
}
