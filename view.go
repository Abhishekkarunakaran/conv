package main

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {

	m.uuidInput.Width = m.width - m.width/10
	m.base64Input.Width = m.width - m.width/10
	m.msgTab.Width = m.width - m.width/10

	uuidStyle := m.styles.InputField
	if m.uuidInput.Focused() {
		uuidStyle = uuidStyle.BorderForeground(m.styles.FocusedColor)
	} else {
		uuidStyle = uuidStyle.BorderForeground(m.styles.BlurredColor)
	}

	base64Style := m.styles.InputField
	if m.base64Input.Focused() {
		base64Style = base64Style.BorderForeground(m.styles.FocusedColor)
	} else {
		base64Style = base64Style.BorderForeground(m.styles.BlurredColor)
	}

	msgTabStyle := m.styles.InputField
	switch m.messageType {
	case "error":
		msgTabStyle = msgTabStyle.Foreground(lipgloss.Color("#ff5555")) // red
	case "success":
		msgTabStyle = msgTabStyle.Foreground(lipgloss.Color("#33cc33")) // green
	case "warning":
		msgTabStyle = msgTabStyle.Foreground(lipgloss.Color("#ffcc00")) // yellow
	default:
		msgTabStyle = msgTabStyle.Foreground(lipgloss.Color("#ffffff")) // white
	}
	instructions := lipgloss.NewStyle().Foreground(lipgloss.Color("#555555")).Padding(1, 0, 0, 0)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Center,
			uuidStyle.Render(
				"   uuid", m.uuidInput.View()),
			base64Style.Render(
				" base64",
				m.base64Input.View()),
			msgTabStyle.Render(
				"message",
				m.msgTab.View(),
			),
			instructions.Render(
				"↑/↓: select • enter: convert • ctrl+p: copy to clipboard • ctrl+c: exit",
			),
		))
}
