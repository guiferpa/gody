package rule

import (
	"fmt"
	"strings"
)

type notEmpty struct{}

func (_ *notEmpty) Name() string {
	return "not_empty"
}

// ErrNotEmpty is the representation about any error happened inside of the rule NotEmpty
type ErrNotEmpty struct {
	Field string
}

func (e *ErrNotEmpty) Error() string {
	return fmt.Sprintf("field %v cannot be empty", e.Field)
}

func (r *notEmpty) Validate(f, v, _ string) (bool, error) {
	if v == "" {
		return true, &ErrNotEmpty{strings.ToLower(f)}
	}
	return true, nil
}
