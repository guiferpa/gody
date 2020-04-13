package gody

import (
	"fmt"
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
	B TestSerializeStructB
	A TestSerializeStructA
}

func TestSerializeBodyStruct(t *testing.T) {
	cases := []struct {
		param interface{}
		ok    bool
	}{
		{map[string]string{"test-key": "test-value"}, false},
		{TestSerializeStructA{a: "a"}, true},
		{TestSerializeStructB{b: "b", a: TestSerializeStructA{a: "a"}}, true},
		{TestSerializeStructC{c: "c", B: TestSerializeStructB{b: "b", a: TestSerializeStructA{a: "a"}}, A: TestSerializeStructA{a: "a"}}, true},
		{10, false},
		{struct{}{}, true},
		{"", false},
		{nil, false},
	}

	for _, c := range cases {
		_, err := serialize(c.param)
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
		}{"test-value"}, true},
		{struct {
			Value string `validate:"required=true"`
		}{"test-value"}, true},
		{struct {
			Value string `validate:"required="`
		}{"test-value"}, false},
		{struct {
			Value string `validate:"=required="`
		}{"test-value"}, false},
		{struct {
			Value string `validate:"="`
		}{"test-value"}, false},
		{struct {
			Value string
		}{"test-value"}, true},
	}

	for _, c := range cases {
		_, err := serialize(c.param)
		if _, ok := err.(*ErrInvalidTag); ok == c.ok {
			t.Error(err)
		}
	}
}

func TestSerialize(t *testing.T) {
	body := struct {
		A string `validate:"test"`
		B int    `json:"b"`
		C bool   `validate:"test test_number=true"`
	}{"a-value", 10, true}

	fields, err := serialize(body)
	if err != nil {
		t.Error(err)
		return
	}

	if got, want := len(fields), 2; got != want {
		t.Errorf("Length of serialized fields isn't equals: got: %v want: %v", got, want)
		return
	}

	wantedFields := []Field{
		{Name: "a", Value: "a-value", Tags: map[string]string{"test": ""}},
		{Name: "c", Value: "true", Tags: map[string]string{"test": "", "test_number": "true"}},
	}
	if got, want := fmt.Sprint(fields), fmt.Sprint(wantedFields); got != want {
		t.Errorf("Serialized fields unexpected: got: %v want: %v", got, want)
		return
	}
}

type TestSerializeSliceA struct {
	E int `validate:"test-slice"`
}

func TestSliceSerialize(t *testing.T) {
	body := struct {
		A string `validate:"test"`
		B []TestSerializeSliceA
	}{"a-value", []TestSerializeSliceA{{10}, {}}}

	fields, err := serialize(body)
	if err != nil {
		t.Error(err)
		return
	}

	if got, want := len(fields), 3; got != want {
		t.Errorf("Length of serialized fields isn't equals: got: %v want: %v", got, want)
		return
	}

	wantedFields := []Field{
		{Name: "a", Value: "a-value", Tags: map[string]string{"test": ""}},
		{Name: "b[0].e", Value: "10", Tags: map[string]string{"test-slice": ""}},
		{Name: "b[1].e", Value: "0", Tags: map[string]string{"test-slice": ""}},
	}
	if got, want := fmt.Sprint(fields), fmt.Sprint(wantedFields); got != want {
		t.Errorf("Serialized fields unexpected: got: %v want: %v", got, want)
		return
	}
}

type TestSerializeStructE struct {
	a string `validate:"test-private-struct-field=300"`
}

type TestSerializeStructD struct {
	J string `validate:"test-struct"`
	I TestSerializeStructE
}

func TestStructSlice(t *testing.T) {
	body := struct {
		A string `validate:"test"`
		B TestSerializeStructD
	}{"a-value", TestSerializeStructD{J: "j-test-struct", I: TestSerializeStructE{a: "a-test-private-struct-field"}}}

	fields, err := serialize(body)
	if err != nil {
		t.Error(err)
		return
	}

	wantedFields := []Field{
		{Name: "a", Value: "a-value", Tags: map[string]string{"test": ""}},
		{Name: "b.j", Value: "j-test-struct", Tags: map[string]string{"test-struct": ""}},
		{Name: "b.i.a", Value: "a-test-private-struct-field", Tags: map[string]string{"test-private-struct-field": "300"}},
	}
	if got, want := fmt.Sprint(fields), fmt.Sprint(wantedFields); got != want {
		t.Errorf("Serialized fields unexpected: got: %v want: %v", got, want)
		return
	}
}

func BenchmarkSerializeBodyStruct(b *testing.B) {
	b.ResetTimer()
	body := map[string]string{"test-key": "test-value"}
	for n := 0; n < b.N; n++ {
		serialize(body)
	}
}
