package listwrapper

import (
	"github.com/charmbracelet/bubbles/list"
)

func New(list list.Model) *State {
	return &State{
		list:   list,
		keyMap: newKeyMap(&list.KeyMap),
	}
}
