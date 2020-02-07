package rule

import (
	"testing"
)

func TestMinName(t *testing.T) {
	r := Min
	if r.Name() != "min" {
		t.Errorf("unexpected result, result: %v, expected: %v", r.Name(), "min")
	}
}

func TestMin(t *testing.T) {
	r := Min
	cases := []struct {
		value, param string
	}{
		{"3", "2"},
		{"2000", "500"},
		{"0", "-1"},
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
