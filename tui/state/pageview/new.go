package pageview

import (
	"github.com/metafates/geminite/page"
)

func New(p *page.Page) *State {
	return &State{
		page:   p,
		keyMap: newKeyMap(),
	}
}
