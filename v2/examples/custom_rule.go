package main

import (
	"fmt"
	"log"

	gody "github.com/guiferpa/gody/v2"
	"github.com/guiferpa/gody/v2/rule"
)

// ErrInvalidPalindrome is a custom error to a specific rule implementation
type ErrInvalidPalindrome struct {
	Value string
}

func (e *ErrInvalidPalindrome) Error() string {
	return fmt.Sprintf("invalid palindrome: %s", e.Value)
}

// PalindromeRule is a struct that implements the Rule interface
type PalindromeRule struct{}

// Name is a func from the Rule contract
func (r *PalindromeRule) Name() string {
	return "palindrome"
}

// Validate is a func from the Rule contract
func (r *PalindromeRule) Validate(f, v, p string) (bool, error) {
	// TODO: The algorithm for palindrome validation
	return true, &ErrInvalidPalindrome{Value: v}
}

func CustomRuleValidation() {
	b := struct {
		Text       string `json:"text" validate:"min_bound=5"`
		Palindrome string `json:"palindrome" validate:"palindrome"`
	}{
		Text:       "test-text",
		Palindrome: "test-palindrome",
	}

	customRules := []gody.Rule{
		&PalindromeRule{},
	}

	valid, err := gody.DefaultValidate(b, customRules)
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
