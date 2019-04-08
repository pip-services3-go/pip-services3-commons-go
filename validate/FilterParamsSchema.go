package validate

import "github.com/pip-services3-go/pip-services3-commons-go/convert"

func NewFilterParamsSchema() *MapSchema {
	return NewMapSchema(convert.String, nil)
}
