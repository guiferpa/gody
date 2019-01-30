package rule

import (
	"fmt"
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
			fmt.Errorf("unexpected result, result: %v, expected: %v", ok, true)
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
			fmt.Errorf("unexpected result, result: %v, expected: %v", ok, true)
		}
	}
}
