package pageview

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/metafates/geminite/bookmark"
	"github.com/metafates/geminite/browser"
	"github.com/metafates/geminite/page"
	"github.com/metafates/geminite/stringutil"
	"github.com/metafates/geminite/tui/base"
	"github.com/metafates/geminite/tui/state/anchorsselect"
	"github.com/skratchdot/open-golang/open"
)

var _ base.State = (*State)(nil)

type State struct {
	page        *page.Page
	browser     *browser.Browser
	viewport    viewport.Model
	keyMap      *keyMap
	initialized bool
}

// Backable implements base.State.
func (*State) Backable() bool {
	return true
}

func (s *State) init(size base.Size) tea.Cmd {
	size.Height -= 1

	if !s.initialized {
		s.viewport = viewport.New(size.Width, size.Height)
		s.initialized = true
	} else {
		s.viewport.Width = size.Width
		s.viewport.Height = size.Height
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(size.Width),
	)
	if err != nil {
		return base.Err(err)
	}

	out, err := s.page.Render(renderer)
	if err != nil {
		return base.Err(err)
	}

	offset := s.viewport.YOffset
	s.viewport.SetContent(out)

	s.viewport.SetYOffset(offset)

	return nil
}

// Init implements base.State.
func (s *State) Init(model base.Model) tea.Cmd {
	return s.init(model.AvailableSize())
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
	return s.init(size)
}

// Status implements base.State.
func (s *State) Status() string {
	return fmt.Sprintf("%3.f%% - %s", s.viewport.ScrollPercent()*100, s.estimateReadingDuration())
}

// Subtitle implements base.State.
func (s *State) Subtitle() string {
	p := s.page
	meta := p.Meta
	if meta.Byline == "" {
		return p.URL.Hostname()
	}

	return fmt.Sprintf("%s @ %s", meta.Byline, p.URL.Hostname())
}

func (s *State) estimateReadingDuration() string {
	duration := s.page.Meta.Duration

	durationLeft := time.Duration(math.Round(float64(duration) * (1 - s.viewport.ScrollPercent())))
	minutes := int(durationLeft.Round(time.Minute).Minutes())

	return stringutil.Quantify(minutes, "minute", "minutes") + " left"
}

// Title implements base.State.
func (s *State) Title() base.Title {
	return base.Title{Text: s.page.Meta.Title}
}

// Update implements base.State.
func (s *State) Update(model base.Model, msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.GotoTop):
			s.viewport.GotoTop()
		case key.Matches(msg, s.keyMap.GotoBottom):
			s.viewport.GotoBottom()
		case key.Matches(msg, s.keyMap.Open):
			err := open.Start(s.page.URL.String())
			if err != nil {
				return base.Err(err)
			}

			return nil
		case key.Matches(msg, s.keyMap.Bookmark):
			err := bookmark.Save(s.page.AsBookmark())
			if err != nil {
				return base.Err(err)
			}

			return nil
		case key.Matches(msg, s.keyMap.Anchors):
			onSelect := func(anchor *page.Anchor) tea.Cmd {
				return func() tea.Msg {
					p, err := s.browser.Open(model.Context(), anchor.URL)
					if err != nil {
						return err
					}

					return New(s.browser, p)
				}
			}

			return base.PushState(anchorsselect.New(s.page, onSelect))
		}
	}

	s.viewport, cmd = s.viewport.Update(msg)
	return cmd
}

// View implements base.State.
func (s *State) View(model base.Model) string {
	return s.viewport.View()
}
