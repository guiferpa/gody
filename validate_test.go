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
