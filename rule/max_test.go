package rule

import (
	"testing"
)

func TestMaxName(t *testing.T) {
	r := Max
	if r.Name() != "max" {
		t.Errorf("unexpected result, result: %v, expected: %v", r.Name(), "max")
	}
}

func TestMax(t *testing.T) {
	r := Max
	cases := []struct {
		value, param string
	}{
		{"2", "3"},
		{"500", "2000"},
		{"-1", "0"},
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

func TestMaxWithInvalidParam(t *testing.T) {
	r := Max
	cases := []struct {
		value, param string
	}{
		{"2", "test"},
		{"500", "true"},
	}
	for _, test := range cases {
		ok, err := r.Validate("", test.value, test.param)
		if err == nil {
			t.Error("unexpected no error")
		}
		if ok {
			t.Errorf("unexpected validation result as okay")
		}
	}
}

func TestMaxWithInvalidValue(t *testing.T) {
	r := Max
	cases := []struct {
		value, param string
	}{
		{"test", "2"},
		{"true", "500"},
	}
	for _, test := range cases {
		ok, err := r.Validate("", test.value, test.param)
		if err == nil {
			t.Error("unexpected no error")
		}
		if ok {
			t.Errorf("unexpected validation result as okay")
		}
	}
}

func TestMaxFailure(t *testing.T) {
	r := Max
	cases := []struct {
		value, param string
	}{
		{"12", "3"},
		{"5000", "2000"},
		{"0", "-1"},
		{"-20", "-22"},
	}
	for _, test := range cases {
		_, err := r.Validate("", test.value, test.param)
		if _, ok := err.(*ErrMax); !ok {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

func TestMaxError(t *testing.T) {
	err := &ErrMax{Field: "text", Value: 3, Max: 1}
	got := err.Error()
	want := "the value 3 in field text is grater than 1"
	if got != want {
		t.Errorf(`&ErrMax{Field: "text", Value: 3, Max: 1}.Error(), got: %v, want: %v`, got, want)
	}
}
