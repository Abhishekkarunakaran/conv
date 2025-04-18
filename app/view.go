package app

import (
	"github.com/Abhishekkarunakaran/ub2/app/constants"
	"github.com/Abhishekkarunakaran/ub2/app/types"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {

	m.uuidInput.Width = m.width - m.width/10
	m.base64Input.Width = m.width - m.width/10
	m.msgTab.Width = m.width - m.width/10

	helpStyle := lipgloss.NewStyle().Foreground(types.HelpGrey).Padding(1, 0, 0, 0)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.renderField(constants.LabelUUID, m.uuidInput, true),
			m.renderField(constants.LabelBase64, m.base64Input, true),
			m.renderField(constants.LabelMesssage, m.msgTab, false),
			helpStyle.Render(constants.HelpString),
		))
}

func (m *model) renderField(label string, field textinput.Model, isInputField bool) string {

	style := m.styles.InputField
	if isInputField {
		if field.Focused() {
			style = style.BorderForeground(m.styles.FocusedColor)
		} else {
			style = style.BorderForeground(m.styles.BlurredColor)
		}
	} else {
		switch m.messageLevel {
		case types.Error:
			style = style.Foreground(types.Red)
		case types.Success:
			style = style.Foreground(types.Green)
		case types.Warn:
			style = style.Foreground(types.Yellow)
		default:
			style = style.Foreground(types.White)
		}
	}
	return style.Render(label, field.View())
}
