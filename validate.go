package gody

import (
	"github.com/guiferpa/gody/rule"
)

func Validate(b interface{}) (bool, error) {
	fields, err := Serialize(b)
	if err != nil {
		return false, err
	}

	for _, field := range fields {
		for _, tag := range field.Tags {
			switch tag.Key {
			case rule.Required.Name():
				if ok, err := rule.Required.Validate(field.Name, field.Value, tag.Value); err != nil {
					return ok, err
				}
			}
		}
	}

	return true, nil
}
