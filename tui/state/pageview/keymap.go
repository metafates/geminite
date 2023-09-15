package pageview

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/metafates/geminite/tui/util"
)

var _ help.KeyMap = (*keyMap)(nil)

type keyMap struct {
	GotoTop, GotoBottom key.Binding
}

func newKeyMap() *keyMap {
	return &keyMap{
		GotoTop:    util.Bind("top", "g"),
		GotoBottom: util.Bind("end", "G"),
	}
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.GotoTop, k.GotoBottom}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}
