package rule

import (
	"reflect"
	"testing"
)

func TestIsBooleanName(t *testing.T) {
	r := IsBoolean 
	got := r.Name()
	want := "is_boolean"
	if got != want {
		t.Errorf("IsBoolean.Name(), got: %v, want: %v", got, want)
	}
}

func TestIsBooleanWithSuccessful(t *testing.T) {
	r := IsBoolean 
	cases := [2]string{"true", "false"}
	for _, test := range cases {
		valid, err := r.Validate("text", test, "")
		if got, want := valid, true; got != want {
			t.Errorf(`valid, _ := IsBoolean.Validate("text", "%v", ""), got: %v, want: %v`, got, test, want)
		}
		if got, want := reflect.TypeOf(err), reflect.TypeOf(nil); got != want {
			t.Errorf(`_, err := IsBoolean.Validate("text", "", ""), got: %v, want: %v`, got, want)
		}
	}
}

func TestIsBooleanWithInvalidValues(t *testing.T) {
	r := IsBoolean 
	ok, err := r.Validate("text", "james", "")
	if err == nil {
		t.Error("Unexpected no error")
	}
	if got, want := ok, true; got != want {
		t.Errorf(`_, err := IsBoolean.Validate("text", "", ""), got: %v, want: %v`, got, want)
	}
}

func TestIsBooleanError(t *testing.T) {
	err := &ErrIsBoolean{Field: "text"}
	got := err.Error()
	want := "field text must be boolean"
	if got != want {
		t.Errorf(`&ErrNotEmpty{Field: "text"}.Error(), got: %v, want: %v`, got, want)
	}
}