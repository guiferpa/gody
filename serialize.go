package gody

import (
	"fmt"
	"reflect"
	"strings"
)

// ErrInvalidBody represents all invalid body report
type ErrInvalidBody struct {
	Kind reflect.Kind
}

func (e *ErrInvalidBody) Error() string {
	return fmt.Sprintln("invalid body:", e.Kind)
}

// ErrInvalidTag represents all invalid tag report
type ErrInvalidTag struct {
	Format string
}

func (e *ErrInvalidTag) Error() string {
	return fmt.Sprintln("invalid tag:", e.Format)
}

type ErrEmptyTagName struct{}

func (e *ErrEmptyTagName) Error() string {
	return "tag name is empty"
}

// Field is a struct to represents the domain about a field inner gody lib
type Field struct {
	Name  string
	Value string
	Tags  map[string]string
}

func Serialize(b any) ([]Field, error) {
	return RawSerialize(DefaultTagName, b)
}

// RawSerialize is a func to serialize/parse all content about the struct input
func RawSerialize(tn string, b any) ([]Field, error) {
	if tn == "" {
		return nil, &ErrEmptyTagName{}
	}

	if b == nil {
		return nil, &ErrInvalidBody{}
	}

	valueOf := reflect.ValueOf(b)
	typeOf := reflect.TypeOf(b)

	if kindOfBody := typeOf.Kind(); kindOfBody != reflect.Struct {
		return nil, &ErrInvalidBody{Kind: kindOfBody}
	}

	fields := make([]Field, 0)
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		tagString := field.Tag.Get(tn)
		if tagString == "" && field.Type.Kind() != reflect.Slice && field.Type.Kind() != reflect.Struct {
			continue
		}

		tagFormats := strings.Fields(tagString)
		tags := make(map[string]string)
		for _, tagFormat := range tagFormats {
			tagFormatSplitted := strings.Split(tagFormat, "=")

			if len(tagFormatSplitted) == 2 {
				tagFormatRule := tagFormatSplitted[0]
				tagFormatValue := tagFormatSplitted[1]
				if tagFormatValue == "" {
					return nil, &ErrInvalidTag{Format: tagFormat}
				}

				tags[tagFormatRule] = tagFormatValue
				continue
			}

			if len(tagFormatSplitted) == 1 {
				tagFormatRule := tagFormatSplitted[0]

				tags[tagFormatRule] = ""
				continue
			}

			return nil, &ErrInvalidTag{Format: tagFormat}
		}

		fieldValue := valueOf.FieldByName(field.Name)
		fieldNameToLower := strings.ToLower(field.Name)
		if fieldNameFromJSONTag := field.Tag.Get("json"); fieldNameFromJSONTag != "" {
			fieldNameToLower = fieldNameFromJSONTag
		}
		if kindOfField := field.Type.Kind(); kindOfField == reflect.Struct {
			if fieldConverted := fieldValue.Convert(fieldValue.Type()); fieldConverted.CanInterface() {
				payload := fieldConverted.Interface()
				serialized, err := RawSerialize(tn, payload)
				if err != nil {
					return nil, err
				}
				for _, item := range serialized {
					fields = append(fields, Field{
						Name:  fmt.Sprintf("%s.%s", fieldNameToLower, item.Name),
						Value: item.Value,
						Tags:  item.Tags,
					})
				}
			}
		} else if kindOfField := field.Type.Kind(); kindOfField == reflect.Slice {
			j := fieldValue.Len()
			for i := 0; i < j; i++ {
				sliceFieldValue := fieldValue.Index(i)
				if sliceFieldConverted := sliceFieldValue.Convert(sliceFieldValue.Type()); sliceFieldConverted.CanInterface() {
					payload := sliceFieldValue.Convert(sliceFieldValue.Type()).Interface()
					serialized, err := RawSerialize(tn, payload)
					if err != nil {
						return nil, err
					}
					for _, item := range serialized {
						fields = append(fields, Field{
							Name:  fmt.Sprintf("%s[%v].%s", fieldNameToLower, i, item.Name),
							Value: item.Value,
							Tags:  item.Tags,
						})
					}
				}
			}
		} else {
			if fieldValue.Kind() == reflect.Pointer {
				_, isNullable := tags[NullableTagName]
				if fieldValue.IsNil() && !isNullable {
					continue
				}
			}
			fieldValueString := fmt.Sprintf("%v", fieldValue)
			fields = append(fields, Field{
				Name:  fieldNameToLower,
				Value: fieldValueString,
				Tags:  tags,
			})
		}
	}

	return fields, nil
}
