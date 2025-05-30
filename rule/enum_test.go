package rule

import (
	"testing"
)

func TestEnumName(t *testing.T) {
	r := Enum
	if r.Name() != "enum" {
		t.Errorf("unexpected result, result: %v, expected: %v", r.Name(), "enum")
	}
}

func TestEnum(t *testing.T) {
	r := Enum
	cases := []struct {
		value, param string
	}{
		{"a", "a,b,c,d"},
		{"b", "a,b,c,d"},
		{"c", "a,b,c,d"},
		{"d", "a,b,c,d"},
		{"", "a,b,c,d"},
	}
	for _, test := range cases {
		ok, err := r.Validate("", test.value, test.param)
		if err != nil {
			t.Error(err)
		}
		if !ok {
			t.Errorf("unexpected result, result: %v, expected: %v", ok, true)
		}
	}
}

func TestEnumWithInvalidParams(t *testing.T) {
	r := Enum
	cases := []struct {
		value, param string
	}{
		{"d", "a,b,c"},
		{"1", "a,b,c"},
	}
	for _, test := range cases {
		ok, err := r.Validate("", test.value, test.param)
		if _, isErrEnum := err.(*ErrEnum); !isErrEnum {
			t.Error(err)
		}
		if !ok {
			t.Errorf("unexpected result, result: %v, expected: %v", ok, true)
		}
	}
}

func TestEnumError(t *testing.T) {
	err := &ErrEnum{Field: "text", Value: "2", Enum: []string{"A", "B"}}
	got := err.Error()
	want := "the value 2 in field text not contains in [A B]"
	if got != want {
		t.Errorf(`&ErrEnum{Field: "text", Value: "2", Enum: []string{"A", "B"}}.Error(), got: %v, want: %v`, got, want)
	}
}
