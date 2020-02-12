package rule

import (
	"fmt"
	"strconv"
	"strings"
)

type min struct{}

func (r *min) Name() string {
	return "min"
}

// ErrMin is the representation about any error happened inside of the rule Min
type ErrMin struct {
	Field string
	Value int
	Min   int
}

func (err *ErrMin) Error() string {
	return fmt.Sprintf("the value %v in field %v is less than %v", err.Value, err.Field, err.Min)
}

func (r *min) Validate(f, v, p string) (bool, error) {
	n, err := strconv.Atoi(p)
	if err != nil {
		return false, err
	}
	vn, err := strconv.Atoi(v)
	if err != nil {
		return false, err
	}
	if vn < n {
		return true, &ErrMin{strings.ToLower(f), vn, n}
	}
	return true, nil
}
