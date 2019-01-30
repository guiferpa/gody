package rule

import (
	"fmt"
	"testing"
)

func TestMaxBoundName(t *testing.T) {
	r := MaxBound
	if r.Name() != "max_bound" {
		t.Errorf("unexpected result, result: %v, expected: %v", r.Name(), "max_bound")
	}
}

func TestMaxBound(t *testing.T) {
	r := MaxBound
	cases := []struct {
		value, param string
	}{
		{"", "0"},
		{"", "1"},
		{"a", "2"},
		{"fla", "3"},
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

func TestMaxBoundWithoutLimit(t *testing.T) {
	r := MaxBound
	cases := []struct {
		value, param string
	}{
		{"fla", "2"},
		{"123", "2"},
		{"1", "0"},
	}
	for _, test := range cases {
		ok, err := r.Validate("", test.value, test.param)
		if _, ok := err.(*ErrMaxBound); !ok {
			t.Error(err)
		}
		if !ok {
			fmt.Errorf("unexpected result, result: %v, expected: %v", ok, true)
		}
	}
}
