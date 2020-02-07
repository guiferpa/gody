package gody

import (
	"testing"
)

type TestSerializeStructA struct {
	a string
}

type TestSerializeStructB struct {
	b string
	a TestSerializeStructA
}

type TestSerializeStructC struct {
	c string
	b TestSerializeStructB
	a TestSerializeStructA
}

func TestSerializeBodyStruct(t *testing.T) {
	cases := []struct {
		param interface{}
		ok    bool
	}{
		{map[string]string{"test-key": "test-value"}, false},
		{TestSerializeStructA{a: "a"}, true},
		{TestSerializeStructB{b: "b", a: TestSerializeStructA{a: "a"}}, true},
		{TestSerializeStructC{c: "c", b: TestSerializeStructB{b: "b", a: TestSerializeStructA{a: "a"}}, a: TestSerializeStructA{a: "a"}}, true},
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

func BenchmarkSerializeBodyStruct(b *testing.B) {
	b.ResetTimer()
	body := map[string]string{"test-key": "test-value"}
	for n := 0; n < b.N; n++ {
		Serialize(body)
	}
}
