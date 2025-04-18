package app

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/Abhishekkarunakaran/ub2/app/types"
	"github.com/gofrs/uuid"
	clipboard "github.com/tiagomelo/go-clipboard/clipboard"
)

func (m *model) UuidToB64() {
	uuidString := m.uuidInput.Value()
	if uuidString == "" {
		m.messageLevel = types.Warn
		m.msgTab.SetValue("uuid: enter valid content")
		return
	}
	uuidVal, err := uuid.FromString(uuidString)
	if err != nil {
		m.messageLevel = types.Error
		m.msgTab.SetValue(err.Error())
		return
	}
	base64Code := b64.StdEncoding.EncodeToString(uuidVal.Bytes())
	m.messageLevel = types.Success
	m.msgTab.SetValue("successfully converted uuid to base64")
	m.base64Input.Focus()
	m.uuidInput.Blur()
	m.base64Input.SetValue(base64Code)
}

func (m *model) B64ToUuid() {
	b64String := m.base64Input.Value()
	if b64String == "" {
		m.messageLevel = types.Warn
		m.msgTab.SetValue("base64: enter valid content")
		return
	}
	uuidVal, err := b64.StdEncoding.DecodeString(b64String)
	if err != nil {
		m.messageLevel = types.Error
		m.msgTab.SetValue(fmt.Sprintf("base64 : %s", err.Error()))
		return
	}
	u, err := uuid.FromBytes(uuidVal)
	if err != nil {
		m.messageLevel = types.Error
		m.msgTab.SetValue(fmt.Sprintf("base64 : %s", err.Error()))
		return
	}
	m.messageLevel = types.Success
	m.msgTab.SetValue("successfully converted base64 to uuid")
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
	c := clipboard.New()
	if stringToCopy == "" || len(stringToCopy) == 0 {
		m.messageLevel = types.Warn
		m.msgTab.SetValue("trying to copy empty string")
		return
	}
	if err := c.CopyText(stringToCopy); err != nil {
		m.messageLevel = types.Error
		m.msgTab.SetValue("failed to copy to clipboard")
	}
	m.messageLevel = types.Success
	m.msgTab.SetValue("copied!")
}

func (m *model) PressedArrowKeys() {
	switch {
	case m.uuidInput.Focused():
		m.uuidInput.Blur()
		m.base64Input.Focus()

	case m.base64Input.Focused():
		m.base64Input.Blur()
		m.uuidInput.Focus()
	}
}
