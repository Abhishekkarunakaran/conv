package main

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/gofrs/uuid"
)

func UuidToB64(uuidString string) (string, error) {
	if uuidString == "" {
		return "", fmt.Errorf("uuid: enter valid content")
	}
	uuidVal, err := uuid.FromString(uuidString)
	if err != nil {
		return "", err
	}
	base64Code := b64.StdEncoding.EncodeToString(uuidVal.Bytes())
	return base64Code, nil
}

func B64ToUuid(b64String string) (string, error) {
	if b64String == "" {
		return "", fmt.Errorf("base64: enter valid content")
	}
	uuidVal, err := b64.StdEncoding.DecodeString(b64String)
	if err != nil {
		return "", fmt.Errorf("base64 : %w", err)
	}
	u, err := uuid.FromBytes(uuidVal)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
