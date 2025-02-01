package rule

import (
	"fmt"
	"strconv"
)

type isBool struct{}

func (r *isBool) Name() string {
	return "is_bool"
}

// ErrIsBool is the representation about any error happened inside of the rule IsBool
type ErrIsBool struct {
	Field string
}

func (err *ErrIsBool) Error() string {
	return fmt.Sprintf("field %v must be 'true' or 'false'", err.Field)
}

func (r *isBool) Validate(f, v, p string) (bool, error) {
	if _, err := strconv.ParseBool(v); err != nil {
		return true, &ErrIsBool{f}
	}
	return true, nil
}
