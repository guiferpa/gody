package gody

import (
	"github.com/guiferpa/gody/rule"
)

func Validate(b interface{}, customRules []rule.Rule) (bool, error) {
	fields, err := Serialize(b)
	if err != nil {
		return false, err
	}

	defaultRules := []rule.Rule{
		rule.Required,
		rule.Enum,
		rule.Max,
		rule.Min,
		rule.MaxBound,
		rule.MinBound,
	}

	rules := append(defaultRules, customRules...)

	for _, field := range fields {
		for _, r := range rules {
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
