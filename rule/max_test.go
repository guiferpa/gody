package rule

import (
	"fmt"
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
			fmt.Errorf("unexpected result, result: %v, expected: %v", ok, true)
		}
	}
}
