package validate

import "github.com/pip-services3-go/pip-services3-commons-go/convert"

func NewPagingParamsSchema() *ObjectSchema {
	return NewObjectSchema().
		WithOptionalProperty("skip", convert.Long).
		WithOptionalProperty("take", convert.Long).
		WithOptionalProperty("total", convert.Boolean)
}
