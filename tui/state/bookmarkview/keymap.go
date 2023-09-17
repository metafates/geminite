package bookmarkview

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/metafates/geminite/tui/util"
)

var _ help.KeyMap = (*keyMap)(nil)

type keyMap struct {
	Open   key.Binding
	Delete key.Binding
}

func newKeyMap() *keyMap {
	return &keyMap{
		Open:   util.Bind("open", "enter"),
		Delete: util.Bind("delete", "d", "backspace"),
	}
}

// FullHelp implements help.KeyMap.
func (s *keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		s.ShortHelp(),
	}
}

// ShortHelp implements help.KeyMap.
func (s *keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		s.Open,
		s.Delete,
	}
}
