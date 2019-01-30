package gody

func Validate(b interface{}) (bool, error) {
	fields, err := Serialize(b)
	if err != nil {
		return false, err
	}
	for _, field := range fields {
		for _, tag := range field.Tags {
			switch tag.Key {
			case rule.MaxBound.Name():
				if ok, err := rule.MaxBound.Validate(field.Name, field.Value, tag.Value); err != nil {
					return ok, err
				}
			case rule.MinBound.Name():
				if ok, err := rule.MinBound.Validate(field.Name, field.Value, tag.Value); err != nil {
					return ok, err
				}
			case rule.Required.Name():
				if ok, err := rule.Required.Validate(field.Name, field.Value, tag.Value); err != nil {
					return ok, err
				}
			case rule.Min.Name():
				if ok, err := rule.Min.Validate(field.Name, field.Value, tag.Value); err != nil {
					return ok, err
				}
			case rule.Enum.Name():
				if ok, err := rule.Enum.Validate(field.Name, field.Value, tag.Value); err != nil {
					return ok, err
				}
			}

		}
	}

	return true, nil
}
