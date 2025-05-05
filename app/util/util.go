package util

import (
	"github.com/Abhishekkarunakaran/ub2/app/types"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

func NewTextInput(charLimit int, placeholder string, focus bool) textinput.Model {
	tm := textinput.New()
	if charLimit != 0 {
		tm.CharLimit = charLimit
	}
	tm.Placeholder = placeholder
	if focus {
		tm.Focus()
	}

	return tm
}

func NewViewport(height int) viewport.Model {
	vp := viewport.New(1, height)
	return vp
}

func DefaultStyles() *types.Styles {
	s := new(types.Styles)
	s.FocusedColor = types.FocusBlue
	s.BlurredColor = types.BlurGrey
	s.InputFieldStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).Padding(1)
	s.ViewportStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).Padding(1)
	return s
}
