package main

import (
	"log"

	"github.com/Abhishekkarunakaran/ub2/app"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := app.NewModel()

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
