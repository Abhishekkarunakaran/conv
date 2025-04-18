package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BlurredColor lipgloss.Color
	FocusedColor lipgloss.Color
	InputField   lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.FocusedColor = lipgloss.Color("#0000ff")
	s.BlurredColor = lipgloss.Color("#888888")
	s.InputField = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).Padding(1)

	return s
}

type model struct {
	width       int
	height      int
	uuidInput   textinput.Model
	base64Input textinput.Model
	errTab      textinput.Model
	styles      *Styles
}

func NewModel() model {
	styles := DefaultStyles()
	uuidInput := textinput.New()
	uuidInput.CharLimit = 36
	uuidInput.Placeholder = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	uuidInput.Focus()
	base64Input := textinput.New()
	base64Input.CharLimit = 64
	base64Input.Placeholder = "xxxxxxxxx"
	errTab := textinput.New()
	errTab.Placeholder = "error message"
	return model{
		uuidInput:   uuidInput,
		base64Input: base64Input,
		errTab:      errTab,
		styles:      styles,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
