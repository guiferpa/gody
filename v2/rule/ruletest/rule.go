package ruletest

type Rule struct {
	name           string
	validated      bool
	err            error
	ValidateCalled bool
}

func (r *Rule) Name() string {
	return r.name
}

func (r *Rule) Validate(name, value, param string) (bool, error) {
	r.ValidateCalled = true
	return r.validated, r.err
}

func NewRule(name string, validated bool, err error) *Rule {
	return &Rule{name, validated, err, false}
}
