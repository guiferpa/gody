package gody

import "fmt"

type ErrDuplicatedRule struct {
	RuleDuplicated Rule
}

func (err *ErrDuplicatedRule) Error() string {
	return fmt.Sprintf("rule %s is duplicated", err.RuleDuplicated.Name())
}

type Validator struct {
	tagName    string
	rulesMap   map[string]Rule
	addedRules []Rule
}

func (v *Validator) AddRules(rs []Rule) error {
	for _, r := range rs {
		if dr, exists := v.rulesMap[r.Name()]; exists {
			return &ErrDuplicatedRule{RuleDuplicated: dr}
		}
		v.rulesMap[r.Name()] = r
	}
	v.addedRules = append(v.addedRules, rs...)
	return nil
}

func (v *Validator) SetTagName(tn string) {
	v.tagName = tn
}

func (v *Validator) Validate(b interface{}) (bool, error) {
	return DefaultValidate(b, v.addedRules, v.tagName)
}

func NewValidator() *Validator {
	tagName := DefaultTagName
	rulesMap := make(map[string]Rule)
	addedRules := make([]Rule, 0)
	return &Validator{tagName, rulesMap, addedRules}
}
