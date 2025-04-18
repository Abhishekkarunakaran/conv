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
	errTab	 	textinput.Model
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
		errTab: errTab,
		styles:      styles,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			switch {
			case m.uuidInput.Focused():
				uuidString := m.uuidInput.Value()
				b64String, err := UuidToB64(uuidString)
				if err != nil {
					m.errTab.SetValue(err.Error())
				} else {
					m.errTab.SetValue("")
					m.uuidInput.Blur()
				m.base64Input.Focus()
				}
				m.base64Input.SetValue(b64String)
			case m.base64Input.Focused():
				b64String := m.base64Input.Value()
				
				uuidString, err := B64ToUuid(b64String)
				if err != nil {
					m.errTab.SetValue(err.Error())
				} else {
					m.errTab.SetValue("")
					m.base64Input.Blur()
					m.uuidInput.Focus()
					}	
				m.uuidInput.SetValue(uuidString)
			}

		case "down":
			if m.uuidInput.Focused() {
				m.uuidInput.Blur()
				m.base64Input.Focus()
			}
		case "up":
			if m.base64Input.Focused() {
				m.base64Input.Blur()
				m.uuidInput.Focus()
			}
		}
	}
	m.uuidInput, cmd = m.uuidInput.Update(msg)
	cmds = append(cmds, cmd)
	m.base64Input, cmd = m.base64Input.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
