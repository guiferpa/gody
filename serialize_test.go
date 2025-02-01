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
		param any
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
		_, err := Serialize(c.param)
		if _, ok := err.(*ErrInvalidBody); ok == c.ok {
			t.Error(err)
		}
	}
}

func TestSerializeBodyTagFormat(t *testing.T) {
	cases := []struct {
		param any
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
		_, err := Serialize(c.param)
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

	fields, err := Serialize(body)
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

	fields, err := Serialize(body)
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

	fields, err := Serialize(body)
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

func TestNil(t *testing.T) {
	body := struct {
		A string    `validate:"test"`
		B any       `validate:"min=10"`
		C *bool     `validate:"min=100 nullable"`
		D *struct{} `validate:"min=1000"`
	}{"a-value", "b-value", nil, nil}

	fields, err := Serialize(body)
	if err != nil {
		t.Error(err)
		return
	}

	wantedFields := []Field{
		{Name: "a", Value: "a-value", Tags: map[string]string{"test": ""}},
		{Name: "b", Value: "b-value", Tags: map[string]string{"min": "10"}},
		{Name: "d", Value: "<nil>", Tags: map[string]string{"min": "1000"}},
	}
	if got, want := fmt.Sprint(fields), fmt.Sprint(wantedFields); got != want {
		t.Errorf("Serialized fields unexpected: got: %v want: %v", got, want)
		return
	}
}

func TestRawSerializeWithEmptyTagName(t *testing.T) {
	_, err := RawSerialize("", nil)
	if err == nil {
		t.Error("Unexpected nil value for error")
		return
	}

	if _, ok := err.(*ErrEmptyTagName); !ok {
		t.Error("Unexpected error type, not equal *ErrEmptyTagName")
		return
	}
}

func TestRawSerializeWithJSONTagName(t *testing.T) {
	body := struct {
		A string `json:"b" validate:"not_empty"`
	}{}

	fields, err := RawSerialize("validate", body)
	if err != nil {
		t.Error(err)
		return
	}

	field := fields[0]

	if got, want := field.Name, "b"; got != want {
		t.Errorf("Unexpected field name, got: %s, want: %s", got, want)
		return
	}
}

func BenchmarkSerializeBodyStruct(b *testing.B) {
	b.ResetTimer()
	body := map[string]string{"test-key": "test-value"}
	for n := 0; n < b.N; n++ {
		_, _ = Serialize(body)
	}
}
