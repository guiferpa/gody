package rule

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

	// IsBool is a rule implemented
	IsBool = &isBool{}
)
