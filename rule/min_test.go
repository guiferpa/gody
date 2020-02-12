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

func TestMinWithInvalidParam(t *testing.T) {
	r := Min
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

func TestMinWithInvalidValue(t *testing.T) {
	r := Min
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

func TestMinFailure(t *testing.T) {
	r := Min
	cases := []struct {
		value, param string
	}{
		{"12", "30"},
		{"5000", "20000"},
		{"-1", "0"},
		{"-22", "-20"},
	}
	for _, test := range cases {
		_, err := r.Validate("", test.value, test.param)
		if _, ok := err.(*ErrMin); !ok {
			t.Errorf("unexpected error: %v", err)
		}
	}
}
