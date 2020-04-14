package main

import (
	"log"

	"github.com/guiferpa/gody"
	"github.com/guiferpa/gody/rule"
)

func SimpleDefaultValidation() {
	b := struct {
		Text string `json:"text" validate:"not_empty"`
	}{}

	valid, err := gody.DefaultValidate(b, nil)
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

func SimplePureValidation() {
	b := struct {
		Text string `json:"text" validate:"not_empty"`
	}{}

	rules := []gody.Rule{
		rule.NotEmpty,
	}
	valid, err := gody.Validate(b, rules)
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

func SimpleValidationFromValidator() {
	b := struct {
		Text string `json:"text" validate:"not_empty"`
	}{}

	validator := gody.NewValidator()

	if err := validator.AddRules(rule.NotEmpty); err != nil {
		log.Println(err)
		return
	}

	valid, err := validator.Validate(b)
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
