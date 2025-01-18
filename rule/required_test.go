package rule

import (
	"testing"
)

func TestRequiredName(t *testing.T) {
	r := Required
	if r.Name() != "required" {
		t.Errorf("unexpected result, result: %v, expected: %v", r.Name(), "required")
	}
}

func TestRequired(t *testing.T) {
	r := Required
	ok, err := r.Validate("", "test", "true")
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Errorf("unexpected result, result: %v, expected: %v", ok, true)
	}
}

func TestRequiredWithEmptyValue(t *testing.T) {
	r := Required
	ok, err := r.Validate("", "", "true")
	if _, ok := err.(*ErrRequired); !ok {
		t.Error(err)
	}
	if !ok {
		t.Errorf("unexpected result, result: %v, expected: %v", ok, true)
	}
}

func TestRequiredWithInvalidParam(t *testing.T) {
	r := Required
	ok, err := r.Validate("", "", "axl-rose&slash")
	if err == nil {
		t.Errorf("unexpected error as result")
	}
	if ok {
		t.Errorf("unexpected result, result: %v, expected: %v", ok, false)
	}
}

func TestRequiredError(t *testing.T) {
	err := &ErrRequired{Field: "text"}
	got := err.Error()
	want := "text is required"
	if got != want {
		t.Errorf(`&ErrRequired{Field: "text"}.Error(), got: %v, want: %v`, got, want)
	}
}
