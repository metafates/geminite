package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/geminite/tui/base"
	"github.com/metafates/geminite/tui/model"
)

func Run(state base.State) error {
	program := tea.NewProgram(model.New(state), tea.WithAltScreen())

	_, err := program.Run()
	return err
}
