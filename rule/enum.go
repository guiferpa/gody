package rule

import (
	"fmt"
	"strings"
)

type enum struct{}

func (r *enum) Name() string {
	return "enum"
}

// ErrEnum is the representation about any error happened inside of the rule Enum
type ErrEnum struct {
	Field string
	Value string
	Enum  []string
}

func (err *ErrEnum) Error() string {
	return fmt.Sprintf("the value %v in field %v not contains in %v", err.Value, err.Field, err.Enum)
}

func (r *enum) Validate(f, v, p string) (bool, error) {
	if v == "" {
		return true, nil
	}
	es := strings.Split(p, ",")
	for _, e := range es {
		if v == e {
			return true, nil
		}
	}
	return false, &ErrEnum{f, v, es}
}
