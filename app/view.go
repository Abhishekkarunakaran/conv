package app

import (
	"github.com/Abhishekkarunakaran/ub2/app/constants"
	"github.com/Abhishekkarunakaran/ub2/app/types"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {

	helpStyle := lipgloss.NewStyle().Foreground(types.HelpGrey).Padding(1, 0, 0, 0)

	return lipgloss.Place(
		m.width, m.height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.renderField(constants.LabelUUID, m.uuidInput),
			m.renderField(constants.LabelBase64, m.base64Input),
			m.renderViewport(m.msgTab),
			helpStyle.Render(constants.HelpString),
		),
	)
}

func (m *model) renderField(label string, field textinput.Model) string {

	style := m.styles.InputFieldStyle

	if field.Focused() {
		style = style.BorderForeground(m.styles.FocusedColor)
	} else {
		style = style.BorderForeground(m.styles.BlurredColor)
	}
	return style.Render(label, field.View())
}

func (m *model) renderViewport(vp viewport.Model) string {

	style := m.styles.ViewportStyle
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

	return style.Render(vp.View())
}
