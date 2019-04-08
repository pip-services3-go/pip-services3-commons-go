package validate

import "github.com/pip-services3-go/pip-services3-commons-go/convert"

func NewProjectionParamsSchema() *ArraySchema {
	return NewArraySchema(convert.String)
}
