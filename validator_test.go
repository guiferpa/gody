package gody

import (
	"testing"

	"github.com/guiferpa/gody/rule/ruletest"
)

func TestDuplicatedRule(t *testing.T) {
	validator := NewValidator()
	rule := ruletest.NewRule("a", true, nil)
	rules := []Rule{
		rule,
		ruletest.NewRule("b", true, nil),
		ruletest.NewRule("c", true, nil),
		rule,
	}
	err := validator.AddRules(rules)
	if err == nil {
		t.Error("Unexpected nil value for duplicated rule error")
		return
	}

	if _, ok := err.(*ErrDuplicatedRule); !ok {
		t.Errorf("Unexpected error type: got: %v", err)
		return
	}
}
