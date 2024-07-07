package password

type PasswordConfig struct {
	Length             int
	IncludeSpecialChar bool
	IncludeUpperCase   bool
	IncludeNumbers     bool
}

type Option func(*PasswordConfig)

func WithLength(length int) Option {
	return func(pc *PasswordConfig) {
		pc.Length = length
	}
}

func WithIncludeSpecialChar(flag bool) Option {
	return func(pc *PasswordConfig) {
		pc.IncludeSpecialChar = flag
	}
}

func WithIncludeUpperCase(flag bool) Option {
	return func(pc *PasswordConfig) {
		pc.IncludeUpperCase = flag
	}
}

func WithIncludeNumbers(flag bool) Option {
	return func(pc *PasswordConfig) {
		pc.IncludeNumbers = flag
	}
}

func NewPasswordConfig(opts ...Option) (pc PasswordConfig) {
	pc = PasswordConfig{
		Length:             20,
		IncludeSpecialChar: false,
		IncludeUpperCase:   false,
		IncludeNumbers:     false,
	}

	for _, opt := range opts {
		opt(&pc)
	}
	return
}
