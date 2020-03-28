package gody

// Rule is a interface with the contract to implement a any rule
type Rule interface {
	Name() string
	Validate(name, value, param string) (bool, error)
}
