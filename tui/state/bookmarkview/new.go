package bookmarkview

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/metafates/geminite/browser"
	"github.com/metafates/geminite/tui/state/listwrapper"
)

func New(b *browser.Browser) *State {
	return &State{
		browser: b,
		keyMap:  newKeyMap(),
		list:    listwrapper.New(list.New(nil, list.DefaultDelegate{}, 0, 0)),
	}
}
