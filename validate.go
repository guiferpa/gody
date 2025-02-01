package gody

import "github.com/guiferpa/gody/v2/rule"

func DefaultValidate(b any, customRules []Rule) (bool, error) {
	return RawDefaultValidate(b, DefaultTagName, customRules)
}

// Validate contains the entrypoint to validation of struct input
func Validate(b any, rules []Rule) (bool, error) {
	return RawValidate(b, DefaultTagName, rules)
}

func RawDefaultValidate(b any, tn string, customRules []Rule) (bool, error) {
	defaultRules := []Rule{
		rule.NotEmpty,
		rule.Required,
		rule.Enum,
		rule.Max,
		rule.Min,
		rule.MaxBound,
		rule.MinBound,
		rule.IsBool,
	}

	return RawValidate(b, tn, append(defaultRules, customRules...))
}

func RawValidate(b any, tn string, rules []Rule) (bool, error) {
	fields, err := RawSerialize(tn, b)
	if err != nil {
		return false, err
	}

	return ValidateFields(fields, rules)
}

func ValidateFields(fields []Field, rules []Rule) (bool, error) {
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
