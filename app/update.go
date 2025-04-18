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
	case tea.KeyMsg:
		switch types.Key(msg.String()) {
		case types.CtrlC:
			return m, tea.Quit
		case types.Enter:
			switch {
			case m.uuidInput.Focused():
				m.UuidToB64()
			case m.base64Input.Focused():
				m.B64ToUuid()
			}
		case types.CtrlP:
			m.CopyToClipboard()
		case types.Down:
			m.PressedArrowKeys()
		case types.Up:
			m.PressedArrowKeys()
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
