package anchorsselect

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/geminite/page"
	"github.com/metafates/geminite/tui/state/listwrapper"
	"github.com/metafates/geminite/tui/util"
)

func New(p *page.Page, onSelect func(*page.Anchor) tea.Cmd) *State {
	lst := util.NewList(
		2,
		"anchor",
		"anchors",
		p.Anchors,
		func(anchor page.Anchor) list.DefaultItem {
			return item{&anchor}
		},
	)

	return &State{
		page:     p,
		list:     listwrapper.New(lst),
		keyMap:   newKeyMap(),
		onSelect: onSelect,
	}
}
