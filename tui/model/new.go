package model

import (
	"context"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/metafates/geminite/tui/base"
	"github.com/zyedidia/generic/stack"
	"golang.org/x/term"
)

func New(state base.State) *Model {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width, height = 80, 40
	}

	model := &Model{
		state:   state,
		history: stack.New[base.State](),
		size: base.Size{
			Width:  width,
			Height: height,
		},
		keyMap: newKeyMap(),
		help:   help.New(),
		styles: base.DefaultStyles(),
	}

	defer model.resize(model.AvailableSize())

	model.context, model.contextCancelFunc = context.WithCancel(context.Background())

	return model
}
