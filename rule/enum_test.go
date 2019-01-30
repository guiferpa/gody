package rule

import (
	"fmt"
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
			fmt.Errorf("unexpected result, result: %v, expected: %v", ok, true)
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
		if _, ok := err.(*ErrEnum); !ok {
			t.Error(err)
		}
		if !ok {
			fmt.Errorf("unexpected result, result: %v, expected: %v", ok, true)
		}
	}
}
