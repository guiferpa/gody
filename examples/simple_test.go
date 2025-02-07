package main

import (
    "testing"
    "github.com/guiferpa/gody/v2/rule"
)


// Test generated using Keploy
func TestSimpleDefaultValidation_EmptyText(t *testing.T) {
    err := SimpleDefaultValidation("")
    if err == nil {
        t.Errorf("Expected error for empty Text, but got nil")
    }
    if _, ok := err.(*rule.ErrNotEmpty); !ok {
        t.Errorf("Expected ErrNotEmpty, but got %v", err)
    }
}

// Test generated using Keploy
func TestSimplePureValidation_EmptyText(t *testing.T) {
    err := SimplePureValidation("")
    if err == nil {
        t.Errorf("Expected error for empty Text, but got nil")
    }
    if _, ok := err.(*rule.ErrNotEmpty); !ok {
        t.Errorf("Expected ErrNotEmpty, but got %v", err)
    }
}


// Test generated using Keploy
func TestSimpleValidationFromValidator_EmptyText(t *testing.T) {
    err := SimpleValidationFromValidator("")
    if err == nil {
        t.Errorf("Expected error for empty Text, but got nil")
    }
    if _, ok := err.(*rule.ErrNotEmpty); !ok {
        t.Errorf("Expected ErrNotEmpty, but got %v", err)
    }
}

