package gody

import (
	"testing"
)

func TestSerializeBodyStruct(t *testing.T) {
	cases := []struct {
		param interface{}
		ok    bool
	}{
		{map[string]string{"test-key": "test-value"}, false},
		{struct{ Value string }{"test-value"}, true},
		{10, false},
		{struct{}{}, true},
		{"", false},
		{nil, false},
	}

	for _, c := range cases {
		_, err := Serialize(c.param)
		if _, ok := err.(*ErrInvalidBody); ok == c.ok {
			t.Error(err)
		}
	}
}

func TestSerializeBodyTagFormat(t *testing.T) {
	cases := []struct {
		param interface{}
		ok    bool
	}{
		{struct {
			Value string `validate:"required"`
		}{"test-value"}, false},
		{struct {
			Value string `validate:"required=true"`
		}{"test-value"}, true},
		{struct {
			Value string `validate:"required="`
		}{"test-value"}, true},
		{struct {
			Value string
		}{"test-value"}, true},
	}

	for _, c := range cases {
		_, err := Serialize(c.param)
		if _, ok := err.(*ErrInvalidTag); ok == c.ok {
			t.Error(err)
		}
	}
}
