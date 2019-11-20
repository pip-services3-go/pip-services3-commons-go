package validate

/*
Schema to validate ProjectionParams
*/
import "github.com/pip-services3-go/pip-services3-commons-go/convert"

// Creates a new instance of validation schema.
// Returns *ArraySchema
func NewProjectionParamsSchema() *ArraySchema {
	return NewArraySchema(convert.String)
}
