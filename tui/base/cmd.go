package base

import tea "github.com/charmbracelet/bubbletea"

func Back() tea.Msg {
	return MsgBack{}
}

func Err(err error) tea.Cmd {
	return func() tea.Msg {
		return err
	}
}
