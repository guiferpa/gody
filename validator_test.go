package gody

import (
	"testing"

	"github.com/guiferpa/gody/v2/rule/ruletest"
)

func TestValidator(t *testing.T) {
	payload := struct {
		A int `validate:"test"`
	}{10}

	validator := NewValidator()

	rule := ruletest.NewRule("test", true, nil)
	if err := validator.AddRules(rule); err != nil {
		t.Error("Unexpected error")
		return
	}
	validated, err := validator.Validate(payload)
	if !validated {
		t.Error("Validated result is not expected")
		return
	}
	if err != nil {
		t.Error("Error result from validate is not expected")
		return
	}
	if !rule.ValidateCalled {
		t.Error("The rule validate wasn't call")
		return
	}
}

func TestDuplicatedRule(t *testing.T) {
	validator := NewValidator()
	rule := ruletest.NewRule("a", true, nil)
	rules := []Rule{
		rule,
		ruletest.NewRule("b", true, nil),
		ruletest.NewRule("c", true, nil),
		rule,
	}
	err := validator.AddRules(rules...)
	if err == nil {
		t.Error("Unexpected nil value for duplicated rule error")
		return
	}

	if _, ok := err.(*ErrDuplicatedRule); !ok {
		t.Errorf("Unexpected error type: got: %v", err)
		return
	}
}
