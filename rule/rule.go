package rule

// Rule is a interface with the contract to implement a any rule
type Rule interface {
	Name() string
	Validate(name, value, param string) (bool, error)
}

var (
	// NotEmpty is a rule implemented
	NotEmpty = &notEmpty{}

	// Required is a rule implemented
	Required = &required{}

	// Max is a rule implemented
	Max = &max{}

	// Min is a rule implemented
	Min = &min{}

	// Enum is a rule implemented
	Enum = &enum{}

	// MaxBound is a rule implemented
	MaxBound = &maxBound{}

	// MinBound is a rule implemented
	MinBound = &minBound{}
)
