package model

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/geminite/tui/base"
)

var (
	_ base.State  = (*errState)(nil)
	_ help.KeyMap = (*errKeyMap)(nil)
)

type errState struct {
	err    error
	keyMap *errKeyMap
}

type errKeyMap struct{}

// FullHelp implements help.KeyMap.
func (*errKeyMap) FullHelp() [][]key.Binding {
	return nil
}

// ShortHelp implements help.KeyMap.
func (*errKeyMap) ShortHelp() []key.Binding {
	return nil
}

func newErrKeyMap() *errKeyMap {
	return &errKeyMap{}
}

func newErrState(err error) *errState {
	return &errState{
		err:    err,
		keyMap: newErrKeyMap(),
	}
}

// Backable implements base.State.
func (*errState) Backable() bool {
	return true
}

// Init implements base.State.
func (*errState) Init(model base.Model) tea.Cmd {
	return nil
}

// Intermediate implements base.State.
func (*errState) Intermediate() bool {
	return true
}

// KeyMap implements base.State.
func (e *errState) KeyMap() help.KeyMap {
	return e.keyMap
}

// Resize implements base.State.
func (*errState) Resize(size base.Size) tea.Cmd {
	return nil
}

// Status implements base.State.
func (*errState) Status() string {
	return ""
}

// Subtitle implements base.State.
func (*errState) Subtitle() string {
	return ""
}

// Title implements base.State.
func (*errState) Title() base.Title {
	return base.Title{
		Text: "Error",
	}
}

// Update implements base.State.
func (*errState) Update(model base.Model, msg tea.Msg) tea.Cmd {
	return nil
}

// View implements base.State.
func (e *errState) View(model base.Model) string {
	return e.err.Error()
}
