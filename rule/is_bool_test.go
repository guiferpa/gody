package rule

import (
	"reflect"
	"testing"
)

func TestIsBoolName(t *testing.T) {
	got := IsBool.Name()
	want := "is_bool"
	if got != want {
		t.Errorf("IsBool.Name(), got: %v, want: %v", got, want)
	}
}

func TestIsBoolWithSuccessful(t *testing.T) {
	valid, err := IsBool.Validate("text", "false", "")
	if got, want := valid, true; got != want {
		t.Errorf(`valid, _ := IsBool.Validate("text", "false", ""), got: %v, want: %v`, got, want)
	}
	if got, want := reflect.TypeOf(err), reflect.TypeOf(nil); got != want {
		t.Errorf(`_, err := IsBool.Validate("text", "false", ""), got: %v, want: %v`, got, want)
	}

	valid, err = IsBool.Validate("text", "true", "")
	if got, want := valid, true; got != want {
		t.Errorf(`valid, _ := IsBool.Validate("text", "true", ""), got: %v, want: %v`, got, want)
	}
	if got, want := reflect.TypeOf(err), reflect.TypeOf(nil); got != want {
		t.Errorf(`_, err := IsBool.Validate("text", "true", ""), got: %v, want: %v`, got, want)
	}
}

func TestIsBoolWithEmptyValue(t *testing.T) {
	valid, err := IsBool.Validate("text", "", "")
	if got, want := valid, true; got != want {
		t.Errorf(`valid, _ := IsBool.Validate("text", "", ""), got: %v, want: %v`, got, want)
	}
	if got, want := reflect.TypeOf(err), reflect.TypeOf(&ErrIsBool{}); got != want {
		t.Errorf(`_, err := IsBool.Validate("text", "", ""), got: %v, want: %v`, got, want)
	}
}

func TestIsBoolError(t *testing.T) {
	err := &ErrIsBool{Field: "text"}
	got := err.Error()
	want := "field text must be 'true' or 'false'"
	if got != want {
		t.Errorf(`&ErrIsBool{Field: "text"}.Error(), got: %v, want: %v`, got, want)
	}
}
