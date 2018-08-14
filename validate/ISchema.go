package validate

type ISchema interface {
	PerformValidation(path string, value interface{}) []*ValidationResult
}
