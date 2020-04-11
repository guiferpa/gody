package gody

import (
	"reflect"
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

func TestSetTagName(t *testing.T) {
	validator := NewValidator()
	if got, want := validator.tagName, DefaultTagName; got != want {
		t.Errorf("Unexpected default tag value from validator struct type: got: %v, want: %v", got, want)
		return
	}

	newTag := "new-tag"
	validator.SetTagName(newTag)
	if got, want := validator.tagName, newTag; got != want {
		t.Errorf("Unexpected default tag value from validator struct type: got: %v, want: %v", got, want)
		return
	}

	err := validator.SetTagName("")
	if err == nil {
		t.Errorf("Unexpected error as nil")
		return
	}

	if ce, ok := err.(*ErrEmptyTagName); !ok {
		t.Errorf("Unexpected error type: got: %v, want: %v", reflect.TypeOf(ce), reflect.TypeOf(&ErrEmptyTagName{}))
		return
	}
}
