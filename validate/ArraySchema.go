package validate

import (
	refl "reflect"
	"strconv"

	"github.com/pip-services-go/pip-services-commons-go/convert"
	"github.com/pip-services-go/pip-services-commons-go/reflect"
)

type ArraySchema struct {
	Schema
	valueType interface{}
}

func NewArraySchema(valueType interface{}) *ArraySchema {
	c := &ArraySchema{
		valueType: valueType,
	}
	c.Schema = *InheritSchema(c)
	return c
}

func (c *ArraySchema) ValueType() interface{} {
	return c.valueType
}

func (c *ArraySchema) SetValueType(value interface{}) {
	c.valueType = value
}

func (c *ArraySchema) PerformValidation(path string, value interface{}) []*ValidationResult {
	name := path
	if name == "" {
		name = "value"
	}
	value = reflect.ObjectReader.GetValue(value)

	results := c.Schema.PerformValidation(path, value)
	if results == nil {
		results = []*ValidationResult{}
	}

	if value == nil {
		return results
	}

	val := refl.ValueOf(value)
	if val.Kind() == refl.Ptr {
		val = val.Elem()
	}

	if val.Kind() == refl.Slice || val.Kind() == refl.Array {
		for index := 0; index < val.Len(); index++ {
			elementPath := strconv.Itoa(index)
			if path != "" {
				elementPath = path + "." + elementPath
			}
			elemResults := c.PerformTypeValidation(elementPath, c.valueType, val.Index(index).Interface())
			if elemResults != nil {
				results = append(results, elemResults...)
			}
		}
	} else {
		results = append(results,
			NewValidationResult(
				path,
				Error,
				"VALUE_ISNOT_ARRAY",
				name+" type must to be List or Array",
				convert.Array,
				convert.TypeConverter.ToTypeCode(value),
			),
		)
	}

	return results
}
