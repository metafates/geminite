package bookmarkview

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/metafates/geminite/bookmark"
)

var _ list.DefaultItem = (*item)(nil)

type item struct {
	Bookmark bookmark.Bookmark
}

// Description implements list.DefaultItem.
func (i item) Description() string {
	return i.Bookmark.URL.Hostname()
}

// FilterValue implements list.DefaultItem.
func (i item) FilterValue() string {
	return i.Bookmark.Title
}

// Title implements list.DefaultItem.
func (i item) Title() string {
	return i.FilterValue()
}
