package rule

import (
	"fmt"
	"strconv"
	"strings"
)

type maxBound struct{}

func (r *maxBound) Name() string {
	return "max_bound"
}

// ErrMaxBound is the representation about any error happened inside of the rule MaxBound
type ErrMaxBound struct {
	Field string
	Value string
	Bound int
}

func (err *ErrMaxBound) Error() string {
	return fmt.Sprintf("the value %v in field %v has character limit greater than %v", err.Value, err.Field, err.Bound)
}

func (r *maxBound) Validate(f, v, p string) (bool, error) {
	n, err := strconv.Atoi(p)
	if err != nil {
		return false, err
	}
	if len(v) > n {
		return true, &ErrMaxBound{strings.ToLower(f), v, n}
	}
	return true, nil
}
