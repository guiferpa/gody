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
			break
		}
	}
}

type ErrInvalidPalindrome struct {
	Value string
}

func (e *ErrInvalidPalindrome) Error() string {
	return fmt.Sprintln("invalid palindrome:", e.Value)
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
			log.Println("required error", err)
			break

		case *ErrInvalidPalindrome:
			log.Println("palindrome error:", err)
			break
		}
	}
}

func main() {
	SimpleValidate()
	CustomValidate()
}
