package validate

type ISchemaBase interface {
	PerformValidation(path string, value interface{}) []*ValidationResult
}
