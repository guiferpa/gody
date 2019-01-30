package rule

type Rule interface {
	Name() string
	Validate(name, value, param string) (bool, error)
}

var (
	Required = &required{}
	Max      = &max{}
	Min      = &min{}
	Enum     = &enum{}
	MaxBound = &maxBound{}
	MinBound = &minBound{}
)
