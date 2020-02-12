package main

import (
	"fmt"
	"log"

	"github.com/guiferpa/gody"
	"github.com/guiferpa/gody/rule"
)

func SimpleValidate() {
	b := struct {
		Text string `json:"text" validate:"required=true"`
	}{}

	valid, err := gody.Validate(b, nil)
	if err != nil {
		if !valid {
			log.Println("body do not validated", err)
		}

		switch err.(type) {
		case *rule.ErrRequired:
			log.Println("required error", err)

		}
	}
}

type ErrInvalidPalindrome struct {
	Value string
}

func (e *ErrInvalidPalindrome) Error() string {
	return fmt.Sprintf("invalid palindrome: %s", e.Value)
}

type PalindromeRule struct{}

func (r *PalindromeRule) Name() string {
	return "palindrome"
}

func (r *PalindromeRule) Validate(f, v, p string) (bool, error) {
	// TODO: The algorithm for palindrome validation
	return true, &ErrInvalidPalindrome{Value: v}
}

func CustomValidate() {
	b := struct {
		Text       string `json:"text" validate:"required=true"`
		Palindrome string `json:"palindrome" validate:"palindrome=true"`
	}{
		Text:       "test-text",
		Palindrome: "test-palindrome",
	}

	customRules := []rule.Rule{
		&PalindromeRule{},
	}

	valid, err := gody.Validate(b, customRules)
	if err != nil {
		if !valid {
			log.Println("body do not validated", err)
		}

		switch err.(type) {
		case *rule.ErrRequired:
			log.Println("required error:", err)

		case *ErrInvalidPalindrome:
			log.Println("palindrome error:", err)
		}
	}
}

type Price struct {
	Currency string `json:"currency" validate:"enum=BRL,EUR,USD"`
	Value    int    `json:"value" validate:"min=10"`
}

type ItemProduct struct {
	Amount int `json:"amount" validate:"min=1"`

	// validate tag's necessary for validation works if not setted it'll be ignored
	Price Price `json:"price" validate:"required=true"`
}

func DeepValidate() {
	ip := ItemProduct{Amount: 10, Price: Price{"BYN", 10000}}

	if valid, err := gody.Validate(ip, nil); err != nil {
		if !valid {
			log.Println("product from cart didn't validate because of", err)
			return
		}

		switch err.(type) {
		case *rule.ErrRequired:
			log.Println("required error:", err)

		case *rule.ErrEnum:
			log.Println("enum error:", err)
		}
	}
}

func main() {
	SimpleValidate()
	CustomValidate()
	DeepValidate()
}
