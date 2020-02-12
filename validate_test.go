package gody

import (
	"testing"

	"github.com/guiferpa/gody/rule"
)

type TestStruct struct {
	A string `validate:"required=true"`
}

func TestValidate(t *testing.T) {
	b := TestStruct{"test"}
	if _, err := Validate(b, nil); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

type TestSimpleStruct struct {
	A string `validate:"required=true"`
	B int    `json:"b" validate:"required=true"`
	C string `json:"c" validate:"required=true"`
	D string `validate:"required=true"`
}

func TestValidateRequiredWithASimpleStruct(t *testing.T) {
	scenes := []struct {
		TestSimpleStruct TestSimpleStruct
		Field            string
	}{
		{TestSimpleStruct{A: "a"}, "c"},
		{TestSimpleStruct{A: "a", C: "c"}, "d"},
	}

	for _, scene := range scenes {
		ok, err := Validate(scene.TestSimpleStruct, nil)
		if !ok {
			t.Error("Struct from current scene isn't valid, something wrong happens")
			return
		}

		ce, ok := err.(*rule.ErrRequired)
		if !ok {
			t.Error("Unexpected error")
			return
		}

		want := scene.Field
		got := ce.Field
		if want != ce.Field {
			t.Errorf("Validate(): want: %s, got: %s", want, got)
		}

	}
}

type TestComplexStructG struct {
	I string `validate:"required=true"`
}

type TestComplexStructE struct {
	F string             `validate:"required=true"`
	G TestComplexStructG `validate:"required=true"`
	H bool
}

type TestComplexStruct struct {
	A string             `validate:"required=true"`
	B int                `json:"b" validate:"required=true"`
	C string             `json:"c" validate:"required=true"`
	D string             `validate:"required=true"`
	E TestComplexStructE `validate:"required=true"`
}

func TestValidateRequiredWithAComplexStruct(t *testing.T) {
	scenes := []struct {
		TestComplexStruct TestComplexStruct
		Field             string
	}{
		{TestComplexStruct{A: "a"}, "c"},
		{TestComplexStruct{A: "a", C: "c"}, "d"},
		{TestComplexStruct{A: "a", B: 1, C: "c", D: "d", E: TestComplexStructE{F: ""}}, "e.f"},
		{TestComplexStruct{A: "a", B: 1, C: "c", D: "d", E: TestComplexStructE{F: "f", G: TestComplexStructG{I: ""}}}, "e.g.i"},
	}

	for _, scene := range scenes {
		ok, err := Validate(scene.TestComplexStruct, nil)
		if !ok {
			t.Error("Struct from current scene isn't valid, something wrong happens")
			return
		}

		ce, ok := err.(*rule.ErrRequired)
		if !ok {
			t.Error("Unexpected error")
			return
		}

		want := scene.Field
		got := ce.Field
		if want != ce.Field {
			t.Errorf("Validate(): want: %s, got: %s", want, got)
		}

	}
}

type TestComplexSliceStructC struct {
	D string `json:"d"`
	E string `json:"e" validate:"enum=e required=true"`
}

type TestComplexSliceStructB struct {
	Cs []TestComplexSliceStructC `json:"cs" validate:"required=true"`
}

type TestComplexiSliceStruct struct {
	A string                  `json:"a" validate:"enum=1,2,3,a"`
	B TestComplexSliceStructB `json:"b" validate:"required=true"`
}

func TestValidateRequiredWithAComplexSliceStruct(t *testing.T) {
	scenes := []struct {
		TestComplexSliceStruct TestComplexiSliceStruct
		Field                  string
	}{
		{TestComplexiSliceStruct{
			A: "a",
			B: TestComplexSliceStructB{
				Cs: []TestComplexSliceStructC{
					{E: "e"},
					{D: "d", E: "e"},
					{D: "test", E: ""},
					{D: "test", E: "e"},
				},
			},
		}, "b.cs[2].e"},
	}

	for _, scene := range scenes {
		ok, err := Validate(scene.TestComplexSliceStruct, nil)
		if !ok {
			t.Error("Struct from current scene isn't valid, something wrong happens")
			return
		}

		ce, ok := err.(*rule.ErrRequired)
		if !ok {
			t.Error("Unexpected error")
			return
		}

		want := scene.Field
		got := ce.Field
		if want != ce.Field {
			t.Errorf("Validate(): want: %s, got: %s", want, got)
		}

	}
}

func TestValidateEnumWithAComplexSliceStruct(t *testing.T) {
	scenes := []struct {
		TestComplexSliceStruct TestComplexiSliceStruct
		Field                  string
	}{
		{TestComplexiSliceStruct{
			A: "a",
			B: TestComplexSliceStructB{
				Cs: []TestComplexSliceStructC{
					{E: "e"},
					{D: "d", E: "e"},
					{D: "test", E: "r"},
					{D: "test", E: "e"},
				},
			},
		}, "b.cs[2].e"},
	}

	for _, scene := range scenes {
		ok, err := Validate(scene.TestComplexSliceStruct, nil)
		if !ok {
			t.Error("Struct from current scene isn't valid, something wrong happens")
			return
		}

		ce, ok := err.(*rule.ErrEnum)
		if !ok {
			t.Error("Unexpected error")
			return
		}

		want := scene.Field
		got := ce.Field
		if want != ce.Field {
			t.Errorf("Validate(): want: %s, got: %s", want, got)
		}

	}
}

func BenchmarkValidateASimpleStruct(b *testing.B) {
	b.ResetTimer()
	body := TestComplexStruct{A: "a", C: "c"}
	for n := 0; n < b.N; n++ {
		Validate(body, nil)
	}
}

func BenchmarkValidateAComplexStruct(b *testing.B) {
	b.ResetTimer()
	body := TestComplexStruct{A: "a", B: 1, C: "c", D: "d", E: TestComplexStructE{F: "f", G: TestComplexStructG{I: ""}}}
	for n := 0; n < b.N; n++ {
		Validate(body, nil)
	}
}

func BenchmarkValidateAComplexSliceStruct(b *testing.B) {
	b.ResetTimer()
	body := TestComplexiSliceStruct{
		A: "a",
		B: TestComplexSliceStructB{
			Cs: []TestComplexSliceStructC{
				{E: "e"},
				{D: "d", E: "e"},
				{D: "test", E: "r"},
				{D: "test", E: "e"},
			},
		},
	}
	for n := 0; n < b.N; n++ {
		Validate(body, nil)
	}
}
