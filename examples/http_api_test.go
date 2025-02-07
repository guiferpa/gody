package main

import (
    "testing"
)


// Test generated using Keploy
func TestErrIsAdultErrorMessage(t *testing.T) {
    err := &ErrIsAdult{}
    expected := "The client isn't a adult then it isn't allowed buy"
    if err.Error() != expected {
        t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
    }
}

// Test generated using Keploy
func TestIsAdultRuleName(t *testing.T) {
    rule := &IsAdultRule{}
    expected := "is_adult"
    if rule.Name() != expected {
        t.Errorf("Expected rule name '%s', got '%s'", expected, rule.Name())
    }
}


// Test generated using Keploy
func TestIsAdultRuleValidateEmptyValue(t *testing.T) {
    rule := &IsAdultRule{adultAge: 21}
    valid, err := rule.Validate("", "", "")
    if !valid || err == nil {
        t.Errorf("Expected validation to fail with ErrIsAdult, got valid=%v, err=%v", valid, err)
    }
}


// Test generated using Keploy
func TestIsAdultRuleValidateNonIntegerValue(t *testing.T) {
    rule := &IsAdultRule{adultAge: 21}
    valid, err := rule.Validate("", "notanumber", "")
    if valid || err == nil {
        t.Errorf("Expected validation to fail with an error, got valid=%v, err=%v", valid, err)
    }
}


// Test generated using Keploy
func TestIsAdultRuleValidateUnderage(t *testing.T) {
    rule := &IsAdultRule{adultAge: 21}
    valid, err := rule.Validate("", "18", "")
    if !valid || err == nil {
        t.Errorf("Expected validation to fail with ErrIsAdult, got valid=%v, err=%v", valid, err)
    }
}


// Test generated using Keploy
func TestIsAdultRuleValidateAdult(t *testing.T) {
    rule := &IsAdultRule{adultAge: 21}
    valid, err := rule.Validate("", "21", "")
    if !valid || err != nil {
        t.Errorf("Expected validation to pass, got valid=%v, err=%v", valid, err)
    }
}

