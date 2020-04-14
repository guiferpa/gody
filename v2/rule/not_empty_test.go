package rule

import (
	"reflect"
	"testing"
)

func TestNotEmptyName(t *testing.T) {
	got := NotEmpty.Name()
	want := "not_empty"
	if got != want {
		t.Errorf("NotEmpty.Name(), got: %v, want: %v", got, want)
	}
}

func TestNotEmptyWithSuccessful(t *testing.T) {
	valid, err := NotEmpty.Validate("text", "test", "")
	if got, want := valid, true; got != want {
		t.Errorf(`valid, _ := NotEmpty.Validate("text", "test", ""), got: %v, want: %v`, got, want)
	}
	if got, want := reflect.TypeOf(err), reflect.TypeOf(nil); got != want {
		t.Errorf(`_, err := NotEmpty.Validate("text", "test", ""), got: %v, want: %v`, got, want)
	}
}

func TestNotEmptyWithEmptyValue(t *testing.T) {
	valid, err := NotEmpty.Validate("text", "", "")
	if got, want := valid, true; got != want {
		t.Errorf(`valid, _ := NotEmpty.Validate("text", "", ""), got: %v, want: %v`, got, want)
	}
	if got, want := reflect.TypeOf(err), reflect.TypeOf(&ErrNotEmpty{}); got != want {
		t.Errorf(`_, err := NotEmpty.Validate("text", "", ""), got: %v, want: %v`, got, want)
	}
}

func TestNotEmptyError(t *testing.T) {
	err := &ErrNotEmpty{Field: "text"}
	got := err.Error()
	want := "field text cannot be empty"
	if got != want {
		t.Errorf(`&ErrNotEmpty{Field: "text"}.Error(), got: %v, want: %v`, got, want)
	}
}
