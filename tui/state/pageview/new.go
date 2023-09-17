package pageview

import (
	"github.com/metafates/geminite/browser"
	"github.com/metafates/geminite/page"
)

func New(b *browser.Browser, p *page.Page) *State {
	return &State{
		browser: b,
		page:    p,
		keyMap:  newKeyMap(),
	}
}
