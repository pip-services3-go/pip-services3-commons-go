package validate

type IValidationRule interface {
	Validate(path string, schema ISchema, value interface{}) []*ValidationResult
}
