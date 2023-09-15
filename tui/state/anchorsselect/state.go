package anchorsselect

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/geminite/page"
	"github.com/metafates/geminite/tui/base"
	"github.com/metafates/geminite/tui/state/listwrapper"
	"github.com/metafates/geminite/tui/state/loading"
)

var _ base.State = (*State)(nil)

type State struct {
	page *page.Page
	list *listwrapper.State

	onSelect func(*page.Anchor) tea.Cmd

	keyMap *keyMap
}

func (s *State) Intermediate() bool {
	return true
}

func (s *State) Backable() bool {
	return s.list.Backable()
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Title() base.Title {
	return base.Title{Text: s.page.Meta.Title}
}

func (s *State) Subtitle() string {
	return s.list.Subtitle()
}

func (s *State) Status() string {
	return s.list.Status()
}

func (s *State) Resize(size base.Size) tea.Cmd {
	return s.list.Resize(size)
}

func (s *State) Update(model base.Model, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.Select):
			item, ok := s.list.SelectedItem().(item)
			if !ok {
				return nil
			}

			return tea.Sequence(
				base.PushState(loading.New("Loading", item.Text)),
				s.onSelect(item.Anchor),
			)
		}
	}

	return s.list.Update(model, msg)
}

func (s *State) View(model base.Model) string {
	return s.list.View(model)
}

func (s *State) Init(model base.Model) tea.Cmd {
	return s.list.Init(model)
}
