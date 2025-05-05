package app

import (
	"github.com/Abhishekkarunakaran/ub2/app/types"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.uuidInput.Width = msg.Width - 20
		m.base64Input.Width = msg.Width - 20
		m.msgTab.Width = msg.Width - 9
	case tea.KeyMsg:
		switch types.Key(msg.String()) {
		case types.Esc:
			return m, tea.Quit
		case types.Enter:
			switch {
			case m.uuidInput.Focused():
				m.UuidToB64()
			case m.base64Input.Focused():
				m.B64ToUuid()
			}
		case types.CtrlC:
			m.CopyToClipboard()
		case types.Down:
			m.PressedDownArrowKey()
		case types.Up:
			m.PressedUpArrowKey()
		case types.CtrlL:
			m.ClearField()
		}
	}
	// to update the inputfield
	m.uuidInput, cmd = m.uuidInput.Update(msg)
	cmds = append(cmds, cmd)
	m.base64Input, cmd = m.base64Input.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
