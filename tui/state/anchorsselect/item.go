package anchorsselect

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/metafates/geminite/page"
	"github.com/metafates/geminite/stringutil"
)

var _ list.DefaultItem = (*item)(nil)

type item struct {
	state  *State
	anchor *page.Anchor
}

func (i item) FilterValue() string {
	return i.anchor.Text
}

func (i item) Title() string {
	return stringutil.Trim(i.FilterValue(), i.state.size.Width)
}

func (i item) Description() string {
	return i.anchor.URL.String()
}
