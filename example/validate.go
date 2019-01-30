package main

import (
	"fmt"

	"github.com/guiferpa/gody"
)

type Body struct {
	Text   string `json:"text" validate:"required=true"`
	Number int    `json:"number" validate:"min=10"`
}

func validate(b interface{}) {
	valid, err := gody.Validate(b, nil)
	if !valid {
		fmt.Println("body do not validated")
	}

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// validation OK without error
	ba := Body{Text: "test-text", Number: 10}
	validate(ba)

	// without validation because invalid body: ptr
	bb := &Body{Text: "test-text", Number: 11}
	validate(bb)

	// validation OK with required error
	bc := Body{Text: "", Number: 11}
	validate(bc)

	// validation OK with min error
	bd := Body{Text: "test-text", Number: 9}
	validate(bd)

	// without validation because invalid body: string
	be := "test-text"
	validate(be)
}
