package validate

/*
Schema to validate TokenizedPagingParams.
*/
import "github.com/pip-services3-go/pip-services3-commons-go/convert"

// Creates a new instance of validation schema.
// Returns *TokenizedPagingParamsSchema
func NewTokenizedPagingParamsSchema() *ObjectSchema {
	return NewObjectSchema().
		WithOptionalProperty("token", convert.String).
		WithOptionalProperty("take", convert.Long).
		WithOptionalProperty("total", convert.Boolean)
}
