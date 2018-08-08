package validate

type ValidationResult struct {
	path     string
	typ      ValidationResultType
	code     string
	message  string
	expected interface{}
	actual   interface{}
}

func NewValidationResult(path string, typ ValidationResultType, code string, message string,
	expected interface{}, actual interface{}) *ValidationResult {
	return &ValidationResult{
		path:     path,
		typ:      typ,
		code:     code,
		message:  message,
		expected: expected,
		actual:   actual,
	}
}

func (c *ValidationResult) Path() string {
	return c.path
}

func (c *ValidationResult) Type() ValidationResultType {
	return c.typ
}

func (c *ValidationResult) Code() string {
	return c.code
}

func (c *ValidationResult) Message() string {
	return c.message
}

func (c *ValidationResult) Expected() interface{} {
	return c.expected
}

func (c *ValidationResult) Actual() interface{} {
	return c.actual
}
