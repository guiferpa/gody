package rule

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type required struct{}

func (r *required) Name() string {
	return "required"
}

// ErrRequired is the representation about any error happened inside of the rule Required
type ErrRequired struct {
	Field string
}

func (e *ErrRequired) Error() string {
	return fmt.Sprintf("%v is required", e.Field)
}

func (r *required) Validate(f, v, p string) (bool, error) {
	b, err := strconv.ParseBool(p)
	if err != nil {
		return false, err
	}
	log.Printf("[guiferpa/gody] :: ATTETION :: required rule is deprecated, please replace to use not_empty.")
	if b && v != "" {
		return true, nil
	}
	return true, &ErrRequired{strings.ToLower(f)}
}
