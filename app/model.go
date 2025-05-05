package app

import (
	"github.com/Abhishekkarunakaran/ub2/app/constants"
	"github.com/Abhishekkarunakaran/ub2/app/types"
	"github.com/Abhishekkarunakaran/ub2/app/util"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
)

type model struct {
	width        int
	height       int
	uuidInput    textinput.Model
	base64Input  textinput.Model
	msgTab       viewport.Model
	messageLevel types.Level
	styles       *types.Styles
}

func NewModel() model {
	styles := util.DefaultStyles()
	uuidInput := util.NewTextInput(36, constants.PhUUID, true)
	base64Input := util.NewTextInput(64, constants.PhBase64, false)
	msgTab := viewport.New(0,1)
	msgTab.SetContent(constants.PhMessage)
	return model{
		uuidInput:   uuidInput,
		base64Input: base64Input,
		msgTab:      msgTab,
		styles:      styles,
	}
}
