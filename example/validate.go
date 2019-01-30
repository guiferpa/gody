package main

import (
	"fmt"

	"github.com/guiferpa/gody"
)

type Body struct {
	Text string `json:"text" validate:"required=true"`
}

func validate(b interface{}) {
	valid, err := gody.Validate(b)
	if !valid {
		fmt.Println("body do not validated")
	}

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// validation OK without error
	ba := Body{Text: "test-text"}
	validate(ba)

	// without validation because invalid body: ptr
	bb := &Body{Text: "test-text"}
	validate(bb)

	// validation OK with required error
	bc := Body{Text: ""}
	validate(bc)

	// without validation because invalid body: string
	bd := "test-text"
	validate(bd)
}
