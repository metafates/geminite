package model

import (
	"context"
	"errors"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/geminite/tui/base"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cmd := m.resize(base.Size{
			Width:  msg.Width,
			Height: msg.Height,
		})

		return m, cmd
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keyMap.Back) && m.state.Backable():
			return m, m.back()
		case key.Matches(msg, m.keyMap.Help):
			m.help.ShowAll = !m.help.ShowAll
			m.resize(m.size)
			return m, nil
		}
	case base.MsgBack:
		// this msg can override Backable() output
		return m, m.back()
	case base.State:
		return m, m.pushState(msg)
	case error:
		if errors.Is(msg, context.Canceled) || strings.Contains(msg.Error(), context.Canceled.Error()) {
			return m, nil
		}

		return m, m.pushState(newErrState(msg))
	}

	cmd := m.state.Update(m, msg)
	return m, cmd
}
