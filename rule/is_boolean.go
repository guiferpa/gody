package rule

import (
	"fmt"
	"strings"
	"strconv"
)

type isBoolean struct{}

func (r *isBoolean) Name() string { return "is_boolean" }

// ErrIsBoolean is the representation about any error happened inside of the rule IsBoolean
type ErrIsBoolean struct {
	Field string
}

func (e *ErrIsBoolean) Error() string { return fmt.Sprintf("field %v must be boolean", e.Field) }

func (r *isBoolean) Validate(f, v, _ string) (bool, error) {
	_, err := strconv.ParseBool(v)
	if err != nil { return true, &ErrIsBoolean{strings.ToLower(f)} }
	return true, nil
}
