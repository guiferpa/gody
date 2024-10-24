package rule

import (
	"reflect"
	"testing"
)

func TestIsBooleanName(t *testing.T) {
	got := NotEmpty.Name()
	want := "is_boolean"
	if got != want {
		t.Error("IsBoolean.Name(), got: %v, want: %v", got, want)
	}
}

func TestIsBooleanWithSuccessful(t *testing.T) {
	valid, err := NotEmpty.Validate("text", "true", "")
	if got, want := valid, true; got != want {
		t.Errorf(`valid, _ := IsBoolean.Validate("text", "true", ""), got: %v, want: %v`, got, want)
	}
	if got, want := reflect.TypeOf(err), reflect.TypeOf(nil); got != want {
		t.Errorf(`_, err := IsBoolean.Validate("text", "test", ""), got: %v, want: %v`, got, want)
	}
}

func TestIsBooleanError(t *testing.T) {
	err := &ErrIsBoolean{field: "text"}
	got := err.Error()
	want := "field text cannot be empty"
	if got != want {
		t.Errorf(`&ErrNotEmpty{Field: "text"}.Error(), got: %v, want: %v`, got, want)
	}
}