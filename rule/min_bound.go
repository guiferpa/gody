package rule

import (
	"fmt"
	"strconv"
	"strings"
)

type minBound struct{}

func (r *minBound) Name() string {
	return "min_bound"
}

type ErrMinBound struct {
	Field string
	Value string
	Bound int
}

func (err *ErrMinBound) Error() string {
	return fmt.Sprintf("the value %v in field %v has character limit less than %v", err.Value, err.Field, err.Bound)
}

func (r *minBound) Validate(f, v, p string) (bool, error) {
	n, err := strconv.Atoi(p)
	if err != nil {
		return false, err
	}
	if len(v) < n {
		return true, &ErrMinBound{strings.ToLower(f), v, n}
	}
	return true, nil
}
