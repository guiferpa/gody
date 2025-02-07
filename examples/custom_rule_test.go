package main

import (
    "testing"
    "log"
)


// Test generated using Keploy
func TestErrInvalidPalindromeError(t *testing.T) {
    err := &ErrInvalidPalindrome{Value: "not-a-palindrome"}
    expected := "invalid palindrome: not-a-palindrome"
    if err.Error() != expected {
        t.Errorf("Expected %v, got %v", expected, err.Error())
    }
}

// Test generated using Keploy
func TestPalindromeRuleName(t *testing.T) {
    rule := &PalindromeRule{}
    expected := "palindrome"
    if rule.Name() != expected {
        t.Errorf("Expected %v, got %v", expected, rule.Name())
    }
}


// Test generated using Keploy
func TestCustomRuleValidation(t *testing.T) {
    // Mocking log output
    var logOutput string
    log.SetOutput(&mockWriter{func(p []byte) (n int, err error) {
        logOutput += string(p)
        return len(p), nil
    }})

    CustomRuleValidation()

    if logOutput == "" {
        t.Error("Expected log output, got none")
    }
}

type mockWriter struct {
    writeFunc func(p []byte) (n int, err error)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
    return m.writeFunc(p)
}

