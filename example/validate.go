package main

import (
	"fmt"
	"log"

	"github.com/guiferpa/gody"
	"github.com/guiferpa/gody/rule"
)

// SimpleValidate is a simple func to example
func SimpleValidate() {
	b := struct {
		Text string `json:"text" validate:"not_empty"`
	}{}

	valid, err := gody.Validate(b, nil)
	if err != nil {
		if !valid {
			log.Println("body do not validated:", err)
		}

		switch err.(type) {
		case *rule.ErrNotEmpty:
			log.Println("not empty error:", err)

		}
	}
}

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

// CustomValidate is a simple func to example about a custom rule
func CustomValidate() {
	b := struct {
		Text       string `json:"text" validate:"min_bound=5"`
		Palindrome string `json:"palindrome" validate:"palindrome"`
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

// Price is a struct about price of ItemProduct
type Price struct {
	Currency string `json:"currency" validate:"enum=BRL,EUR,USD"`
	Value    int    `json:"value" validate:"min=10"`
}

// ItemProduct is a struct to example
type ItemProduct struct {
	Amount int `json:"amount" validate:"min=1"`

	// validate tag's necessary for validation works if not setted it'll be ignored
	Price Price `json:"price" validate:"required"`
}

// DeepValidate is a simple func to example a deep validation
func DeepValidate() {
	ip := ItemProduct{Amount: 10, Price: Price{"BYN", 10000}}

	if valid, err := gody.Validate(ip, nil); err != nil {
		if !valid {
			log.Println("product from cart didn't validate because of", err)
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
