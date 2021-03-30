package validate

/*
Result generated by schema validation
*/
type ValidationResult struct {
	path     string
	typ      ValidationResultType
	code     string
	message  string
	expected interface{}
	actual   interface{}
}

// Creates a new instance of validation ressult and sets its values.
// see
// ValidationResultType
// Parameters:
//   - path string
//   a dot notation path of the validated element.
//   - type: ValidationResultType
//   a type of the validation result: Information, Warning, or Error.
//   - code string
//   an error code.
//   - message string
//   a human readable message.
//   - expected interface{}
//   an value expected by schema validation.
//   - actual interface{}
//   an actual value found by schema validation.
// Returns *ValidationResult
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

// Gets dot notation path of the validated element.
// Returns string
// the path of the validated element.
func (c *ValidationResult) Path() string {
	return c.path
}

// Gets the type of the validation result: Information, Warning, or Error.
// see
// ValidationResultType
// Returns ValidationResultType
// the type of the validation result.
func (c *ValidationResult) Type() ValidationResultType {
	return c.typ
}

// Gets the error code.
// Returns string
// the error code
func (c *ValidationResult) Code() string {
	return c.code
}

// Gets the human readable message.
// Returns string
// the result message.
func (c *ValidationResult) Message() string {
	return c.message
}

// Gets the value expected by schema validation.

// Returns any
// the expected value.
func (c *ValidationResult) Expected() interface{} {
	return c.expected
}

// Gets the actual value found by schema validation.
// Returns any
// the actual value.
func (c *ValidationResult) Actual() interface{} {
	return c.actual
}
