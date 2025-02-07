package main

import (
    "log"

    gody "github.com/guiferpa/gody/v2"
    "github.com/guiferpa/gody/v2/rule"
)

func SimpleDefaultValidation(text string) error {
    b := struct {
        Text string `json:"text" validate:"not_empty"`
    }{
        Text: text,
    }

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
    return err
}

func SimplePureValidation(text string) error {
    b := struct {
        Text string `json:"text" validate:"not_empty"`
    }{
        Text: text,
    }

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
    return err
}

func SimpleValidationFromValidator(text string) error {
    b := struct {
        Text string `json:"text" validate:"not_empty"`
    }{
        Text: text,
    }

    validator := gody.NewValidator()

    if err := validator.AddRules(rule.NotEmpty); err != nil {
        log.Println(err)
        return err
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
    return err
}
