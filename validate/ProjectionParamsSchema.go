package validate

import "github.com/pip-services-go/pip-services-commons-go/convert"

func NewProjectionParamsSchema() *ArraySchema {
	return NewArraySchema(convert.String)
}
