package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
)

func main() {

	uuidString := os.Args[1]
	uuid, err := uuid.FromString(uuidString)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	base64Code := b64.StdEncoding.EncodeToString(uuid.Bytes())
	fmt.Println(base64Code)

}
