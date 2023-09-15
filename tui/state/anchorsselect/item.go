package anchorsselect

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/metafates/geminite/page"
	"github.com/metafates/geminite/stringutil"
)

var _ list.DefaultItem = (*item)(nil)

type item struct {
	*page.Anchor
}

func (i item) FilterValue() string {
	return i.Text
}

func (i item) Title() string {
	return stringutil.Trim(i.FilterValue(), 20)
}

func (i item) Description() string {
	return i.URL.String()
}
