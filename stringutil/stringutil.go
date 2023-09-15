package stringutil

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Trim trims a string to a maximum length, appending an ellipsis if necessary.
// The trimmed string with ellipsis will never be longer than max.
// Works with ANSI escape codes.
func Trim(s string, max int) string {
	if max <= 0 {
		panic("max must be greater than 0")
	}

	stringLength := lipgloss.Width(s)

	if s == "" {
		return s
	}

	const ellipsis = '…'

	if max == 1 {
		return string(ellipsis)
	}

	if stringLength < max {
		return s
	}

	trimmed := lipgloss.NewStyle().MaxWidth(max - 1).Render(s)

	// get index of \x1b
	idx := strings.LastIndex(trimmed, "\x1b")
	if idx == -1 {
		return trimmed + string(ellipsis)
	}

	// insert ellipsis before \x1b
	return trimmed[:idx] + string(ellipsis) + trimmed[idx:]
}

// Quantify returns a string with the quantity and the correct form of the
// word, depending on the quantity.
func Quantify(n int, singular, plural string) string {
	var form string
	if n == 1 {
		form = singular
	} else {
		form = plural
	}

	return fmt.Sprint(n, " ", form)
}
