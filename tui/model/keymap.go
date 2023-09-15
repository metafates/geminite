package model

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/metafates/geminite/tui/util"
)

type keyMap struct {
	Back, Quit, Help key.Binding
}

func newKeyMap() *keyMap {
	return &keyMap{
		Back: util.Bind("back", "esc"),
		Quit: util.Bind("quit", "ctrl+c", "ctrl+d"),
		Help: util.Bind("help", "?"),
	}
}
