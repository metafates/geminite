package anchorsselect

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/metafates/geminite/tui/util"
)

var _ help.KeyMap = (*keyMap)(nil)

type keyMap struct {
	Select key.Binding
}

func newKeyMap() *keyMap {
	return &keyMap{
		Select: util.Bind("select", "enter"),
	}
}

func (k *keyMap) ShortHelp() []key.Binding {
	return nil
}

func (k *keyMap) FullHelp() [][]key.Binding {
	return nil
}
