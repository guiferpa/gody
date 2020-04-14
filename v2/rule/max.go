package rule

import (
	"fmt"
	"strconv"
	"strings"
)

type max struct{}

func (r *max) Name() string {
	return "max"
}

// ErrMax is the representation about any error happened inside of the rule Max
type ErrMax struct {
	Field string
	Value int
	Max   int
}

func (err *ErrMax) Error() string {
	return fmt.Sprintf("the value %v in field %v is grater than %v", err.Value, err.Field, err.Max)
}

func (r *max) Validate(f, v, p string) (bool, error) {
	n, err := strconv.Atoi(p)
	if err != nil {
		return false, err
	}
	vn, err := strconv.Atoi(v)
	if err != nil {
		return false, err
	}
	if vn > n {
		return true, &ErrMax{strings.ToLower(f), vn, n}
	}
	return true, nil
}
