package validate

/*
Schema to validate PagingParams.
*/
import "github.com/pip-services3-go/pip-services3-commons-go/convert"

// Creates a new instance of validation schema.
// Returns *PagingParamsSchema
func NewPagingParamsSchema() *ObjectSchema {
	return NewObjectSchema().
		WithOptionalProperty("skip", convert.Long).
		WithOptionalProperty("take", convert.Long).
		WithOptionalProperty("total", convert.Boolean)
}
