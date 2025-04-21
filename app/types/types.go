package types

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	BlurredColor lipgloss.Color
	FocusedColor lipgloss.Color
	InputField   lipgloss.Style
}

type Level string

var (
	Error   Level = "error"
	Warn    Level = "warn"
	Success Level = "success"
	Nil     Level = ""
)

type Key string

var (
	CtrlC Key = "ctrl+c"
	Enter Key = "enter"
	Down  Key = "down"
	Up    Key = "up"
	CtrlL Key = "ctrl+l"
	Esc Key = "esc"
)
