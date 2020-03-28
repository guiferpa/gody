package gody

import "github.com/guiferpa/gody/rule"

func DefaultValidate(b interface{}, customRules []Rule) (bool, error) {
	defaultRules := []Rule{
		rule.NotEmpty,
		rule.Required,
		rule.Enum,
		rule.Max,
		rule.Min,
		rule.MaxBound,
		rule.MinBound,
	}

	return Validate(b, append(defaultRules, customRules...))
}

// Validate contains the entrypoint to validation of struct input
func Validate(b interface{}, rules []Rule) (bool, error) {
	fields, err := Serialize(b)
	if err != nil {
		return false, err
	}

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
