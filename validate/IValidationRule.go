package validate

type IValidationRule interface {
	Validate(path string, schema Schema, value interface{}, results []*ValidationResult)
}
