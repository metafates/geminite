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
	"github.com/metafates/geminite/page"
	"github.com/metafates/geminite/stringutil"
	"github.com/metafates/geminite/tui/base"
)

var _ base.State = (*State)(nil)

type State struct {
	page        *page.Page
	viewport    viewport.Model
	keyMap      *keyMap
	initialized bool
}

// Backable implements base.State.
func (*State) Backable() bool {
	return true
}

func (s *State) init(size base.Size) tea.Cmd {
	size.Height--

	if !s.initialized {
		s.viewport = viewport.New(size.Width, size.Height)
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
	page := s.page
	meta := page.Meta
	if meta.Byline == "" {
		return page.URL.Hostname()
	}

	return fmt.Sprintf("%s @ %s", meta.Byline, page.URL.Hostname())
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
		}
	}

	s.viewport, cmd = s.viewport.Update(msg)
	return cmd
}

// View implements base.State.
func (s *State) View(model base.Model) string {
	return s.viewport.View()
}
