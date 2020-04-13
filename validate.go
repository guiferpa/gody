package gody

import "github.com/guiferpa/gody/rule"

// DefaultValidate validates the validation subject against pre-defined rules
func DefaultValidate(validationSubject interface{}, customRules []Rule, tn ...string) (bool, error) {
	defaultRules := []Rule{
		rule.NotEmpty,
		rule.Required,
		rule.Enum,
		rule.Max,
		rule.Min,
		rule.MaxBound,
		rule.MinBound,
	}

	return Validate(validationSubject, append(defaultRules, customRules...), tn...)
}

// Validate contains the entrypoint to validation of struct input
func Validate(validationSubject interface{}, rules []Rule, tn ...string) (bool, error) {
	fields, err := Serialize(validationSubject, tn...)
	if err != nil {
		return false, err
	}

	return validateFields(fields, rules)
}

func validateFields(fields []Field, rules []Rule) (bool, error) {
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
