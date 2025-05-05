package app

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/Abhishekkarunakaran/ub2/app/constants"
	"github.com/Abhishekkarunakaran/ub2/app/types"
	clipboard "github.com/atotto/clipboard"
	"github.com/gofrs/uuid"
)

func (m *model) setMessage(message string, level types.Level) {
	m.messageLevel = level
	m.msgTab.SetContent(fmt.Sprintf(" %s",message))
}
func (m *model) UuidToB64() {
	uuidString := m.uuidInput.Value()
	if uuidString == "" {
		m.setMessage("uuid: enter valid content", types.Warn)
		return
	}
	uuidVal, err := uuid.FromString(uuidString)
	if err != nil {
		m.setMessage(err.Error(), types.Error)
		return
	}
	base64Code := b64.StdEncoding.EncodeToString(uuidVal.Bytes())
	m.setMessage("successfully converted uuid to base64", types.Success)
	m.base64Input.Focus()
	m.uuidInput.Blur()
	m.base64Input.SetValue(base64Code)
}

func (m *model) B64ToUuid() {
	b64String := m.base64Input.Value()
	if b64String == "" {
		m.setMessage("base64: enter valid content", types.Warn)
		return
	}
	uuidVal, err := b64.StdEncoding.DecodeString(b64String)
	if err != nil {
		m.setMessage(fmt.Sprintf("base64 : %s", err.Error()), types.Error)
		return
	}
	u, err := uuid.FromBytes(uuidVal)
	if err != nil {
		m.setMessage(err.Error(), types.Error)
		return
	}
	m.setMessage("successfully converted base64 to uuid", types.Success)
	m.base64Input.Blur()
	m.uuidInput.Focus()
	m.uuidInput.SetValue(u.String())
}

func (m *model) CopyToClipboard() {
	switch {
	case m.uuidInput.Focused():
		m.copy(m.uuidInput.Value())
	case m.base64Input.Focused():
		m.copy(m.base64Input.Value())
	}
}

func (m *model) copy(stringToCopy string) {

	if stringToCopy == "" || len(stringToCopy) == 0 {
		m.setMessage("trying to copy empty string", types.Warn)
		return
	}
	if err := clipboard.WriteAll(stringToCopy); err != nil {
		m.setMessage("failed to copy to clipboard", types.Error)
	}
	m.setMessage("copied!", types.Success)
}

func (m *model) PressedUpArrowKey() {
	if m.base64Input.Focused() {
		m.base64Input.Blur()
		m.uuidInput.Focus()
		m.setMessage(constants.PhMessage, types.Nil)
	}
}

func (m *model) PressedDownArrowKey() {
	if m.uuidInput.Focused() {
		m.uuidInput.Blur()
		m.base64Input.Focus()
		m.setMessage(constants.PhMessage, types.Nil)
	}
}

func (m *model) ClearField() {
	switch {
	case m.uuidInput.Focused() && m.uuidInput.Value() != "":
		m.uuidInput.SetValue("")
		m.setMessage("successfully cleared!", types.Success)
	case m.base64Input.Focused() && m.base64Input.Value() != "":
		m.base64Input.SetValue("")
		m.setMessage("successfully cleared!", types.Success)
	default:
		m.setMessage("nothing to clear!", types.Warn)
	}
}
