package rule

import (
	"testing"
)

func TestMinBoundName(t *testing.T) {
	r := MinBound
	if r.Name() != "min_bound" {
		t.Errorf("unexpected result, result: %v, expected: %v", r.Name(), "min_bound")
	}
}

func TestMinBound(t *testing.T) {
	r := MinBound
	cases := []struct {
		value, param string
	}{
		{"", "0"},
		{"a", "1"},
		{"aaa", "2"},
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

func TestMinBoundWithoutLimit(t *testing.T) {
	r := MinBound
	cases := []struct {
		value, param string
	}{
		{"fl", "3"},
		{"123", "4"},
	}
	for _, test := range cases {
		ok, err := r.Validate("", test.value, test.param)
		if _, ok := err.(*ErrMinBound); !ok {
			t.Error(err)
		}
		if !ok {
			t.Errorf("unexpected result, result: %v, expected: %v", ok, true)
		}
	}
}

func TestMinBoundWithInvalidParam(t *testing.T) {
	r := MinBound
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

func TestMinBoundError(t *testing.T) {
	err := &ErrMinBound{Field: "text", Value: "2", Bound: 1}
	got := err.Error()
	want := "the value 2 in field text has character limit less than 1"
	if got != want {
		t.Errorf(`&ErrMinBound{Field: "text", Value: "2", Bound: 1}.Error(), got: %v, want: %v`, got, want)
	}
}
