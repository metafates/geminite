package model

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/geminite/stringutil"
	"github.com/muesli/reflow/wordwrap"
)

func (m *Model) View() string {
	const newline = "\n"

	title := m.state.Title()
	titleStyle := m.styles.Title.Copy()

	if title.Background != "" {
		titleStyle.Background(title.Background)
	}

	if title.Foreground != "" {
		titleStyle.Foreground(title.Foreground)
	}

	titleText := stringutil.Trim(title.Text, m.size.Width/2)

	header := m.styles.TitleBar.Render(titleStyle.Render(titleText) + " " + m.state.Status())

	subtitle := m.state.Subtitle()
	if subtitle != "" {
		header = lipgloss.JoinVertical(lipgloss.Left, header, m.styles.TitleBar.Render(m.styles.Subtitle.Render(m.state.Subtitle())))
		// header += m.styles.TitleBar.Render(m.styles.Subtitle.Render(m.state.Subtitle()))
	}

	view := wordwrap.String(m.state.View(m), m.size.Width)

	keyMapHelp := m.styles.HelpBar.Render(m.help.View(m))

	headerHeight := lipgloss.Height(header)
	viewHeight := lipgloss.Height(view)
	helpHeight := lipgloss.Height(keyMapHelp)

	diff := m.size.Height - headerHeight - viewHeight - helpHeight

	var filler string
	if diff > 0 {
		filler = strings.Repeat(newline, diff)
	}

	// return lipgloss.JoinVertical(lipgloss.Left, header, view, filler, keyMapHelp)

	return header + newline + view + filler + newline + keyMapHelp
}
