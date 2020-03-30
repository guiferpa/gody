package gody

import (
	"testing"

	"github.com/guiferpa/gody/rule/ruletest"
)

type StructAForTest struct {
	A string `validate:"fake"`
	B string
}

func TestValidateMatchRule(t *testing.T) {
	payload := StructAForTest{A: "", B: "test-b"}
	rule := ruletest.NewRule("fake", true, nil)

	validated, err := Validate(payload, []Rule{rule})
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

func TestValidateNoMatchRule(t *testing.T) {
	payload := StructAForTest{A: "", B: "test-b"}
	rule := ruletest.NewRule("mock", true, nil)

	validated, err := Validate(payload, []Rule{rule})
	if !validated {
		t.Error("Validated result is not expected")
		return
	}
	if err != nil {
		t.Error("Error result from validate is not expected")
		return
	}
	if rule.ValidateCalled {
		t.Error("The rule validate was call")
		return
	}
}

type StructBForTest struct {
	C int `validate:"test"`
	D bool
}

type errStructBForValidation struct{}

func (_ *errStructBForValidation) Error() string {
	return ""
}

func TestValidateWithRuleError(t *testing.T) {
	payload := StructBForTest{C: 10}
	rule := ruletest.NewRule("test", true, &errStructBForValidation{})

	validated, err := Validate(payload, []Rule{rule})
	if !validated {
		t.Error("Validated result is not expected")
		return
	}
	if !rule.ValidateCalled {
		t.Error("The rule validate was call")
		return
	}
	if _, ok := err.(*errStructBForValidation); !ok {
		t.Errorf("Unexpected error type: got: %v", err)
		return
	}
}
