package gody

import (
	"github.com/guiferpa/gody/rule"
)

func Validate(b interface{}) (bool, error) {
	fields, err := Serialize(b)
	if err != nil {
		return false, err
	}

	defaultRules := []rule.Rule{
		rule.Required,
	}

	for _, field := range fields {
		for _, r := range defaultRules {
			val, ok := field.Tags[r.Name()]
			if !ok {
				continue
			}
			if ok, err := r.Validate(field.Name, field.Value, val); err != nil {
				return ok, err
			}
		}
	}

	return true, nil
}
