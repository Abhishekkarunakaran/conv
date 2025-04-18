package main

import (
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
		key := msg.String()
		switch key {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			switch {
			case m.uuidInput.Focused():
				uuidString := m.uuidInput.Value()
				b64String,messageType,err  := UuidToB64(uuidString)
				if err != nil {
					m.messageType = messageType
					m.msgTab.SetValue(err.Error())
				} else {
					m.msgTab.SetValue("")
					m.uuidInput.Blur()
					m.base64Input.Focus()
				}
				m.base64Input.SetValue(b64String)
			case m.base64Input.Focused():
				b64String := m.base64Input.Value()

				uuidString,messageType,err := B64ToUuid(b64String)
				if err != nil {
					m.msgTab.SetValue(err.Error())
					m.messageType= messageType
				} else {
					m.msgTab.SetValue("")
					m.base64Input.Blur()
					m.uuidInput.Focus()
				}
				m.uuidInput.SetValue(uuidString)
			}
		case "ctrl+p":
			if err := copyToClipboard(&m); err != nil {
				m.msgTab.SetValue(err.Error())
			}
			m.msgTab.SetValue("copied!!")
			m.messageType= "success"
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

