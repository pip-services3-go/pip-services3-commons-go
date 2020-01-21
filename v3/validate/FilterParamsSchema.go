package validate

/*
Schema to validate FilterParams.
*/
import "github.com/pip-services3-go/pip-services3-commons-go/v3/convert"

// Creates a new instance of validation schema.
// Returns *MapSchema
func NewFilterParamsSchema() *MapSchema {
	return NewMapSchema(convert.String, nil)
}
