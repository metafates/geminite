package bookmarkview

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/geminite/bookmark"
	"github.com/metafates/geminite/browser"
	"github.com/metafates/geminite/tui/base"
	"github.com/metafates/geminite/tui/state/listwrapper"
	"github.com/metafates/geminite/tui/state/loading"
	"github.com/metafates/geminite/tui/state/pageview"
	"github.com/metafates/geminite/tui/util"
)

var _ base.State = (*State)(nil)

type State struct {
	browser *browser.Browser
	list    *listwrapper.State
	keyMap  *keyMap
}

// Backable implements base.State.
func (s *State) Backable() bool {
	return s.list.Backable()
}

// Init implements base.State.
func (s *State) Init(model base.Model) tea.Cmd {
	bookmarks, err := bookmark.GetAll()
	if err != nil {
		return base.Err(err)
	}

	s.list = listwrapper.New(util.NewList(
		2,
		"bookmark",
		"bookmarks",
		bookmarks,
		func(b bookmark.Bookmark) list.DefaultItem {
			return item{Bookmark: b}
		},
	))

	return s.Resize(model.AvailableSize())
}

// Intermediate implements base.State.
func (*State) Intermediate() bool {
	return false
}

// KeyMap implements base.State.
func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

// Resize implements base.State.
func (s *State) Resize(size base.Size) tea.Cmd {
	return s.list.Resize(size)
}

// Status implements base.State.
func (s *State) Status() string {
	return s.list.Status()
}

// Subtitle implements base.State.
func (s *State) Subtitle() string {
	return s.list.Subtitle()
}

// Title implements base.State.
func (*State) Title() base.Title {
	return base.Title{Text: "Bookmarks"}
}

// Update implements base.State.
func (s *State) Update(model base.Model, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if s.list.FilterState() == list.Filtering {
			goto end
		}

		switch {
		case key.Matches(msg, s.keyMap.Delete):
			selectedItem, ok := s.list.SelectedItem().(item)
			if !ok {
				return nil
			}

			if err := bookmark.Delete(selectedItem.Bookmark.URL); err != nil {
				return base.Err(err)
			}

			return s.Init(model)
		case key.Matches(msg, s.keyMap.Open):
			selectedItem, ok := s.list.SelectedItem().(item)
			if !ok {
				return nil
			}

			return tea.Sequence(
				base.PushState(loading.New("Opening", selectedItem.Title())),
				func() tea.Msg {
					p, err := s.browser.Open(model.Context(), selectedItem.Bookmark.URL)
					if err != nil {
						return err
					}

					return pageview.New(s.browser, p)
				},
			)
		}
	}

end:
	return s.list.Update(model, msg)
}

// View implements base.State.
func (s *State) View(model base.Model) string {
	return s.list.View(model)
}
