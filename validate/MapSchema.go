package validate

import (
	refl "reflect"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/reflect"
)

type MapSchema struct {
	Schema
	keyType   interface{}
	valueType interface{}
}

func NewMapSchema(keyType interface{}, valueType interface{}) *MapSchema {
	c := &MapSchema{
		keyType:   keyType,
		valueType: valueType,
	}
	c.Schema = *InheritSchema(c)
	return c
}

func NewMapSchemaWithRules(keyType interface{}, valueType interface{}, required bool, rules []IValidationRule) *MapSchema {
	c := &MapSchema{
		keyType:   keyType,
		valueType: valueType,
	}
	c.Schema = *InheritSchemaWithRules(c, required, rules)
	return c
}

func (c *MapSchema) KeyType() interface{} {
	return c.keyType
}

func (c *MapSchema) SetKeyType(value interface{}) {
	c.keyType = value
}

func (c *MapSchema) ValueType() interface{} {
	return c.valueType
}

func (c *MapSchema) SetValueType(value interface{}) {
	c.valueType = value
}

func (c *MapSchema) PerformValidation(path string, value interface{}) []*ValidationResult {
	value = reflect.ObjectReader.GetValue(value)

	results := c.Schema.PerformValidation(path, value)
	if results == nil {
		results = []*ValidationResult{}
	}

	if value == nil {
		return results
	}

	name := path
	if name == "" {
		name = "value"
	}

	val := refl.ValueOf(value)

	if val.Kind() == refl.Map {
		for _, keyVal := range val.MapKeys() {
			elementPath := convert.StringConverter.ToString(keyVal.Interface())
			if path != "" {
				elementPath = path + "." + elementPath
			}

			keyResults := c.PerformTypeValidation(elementPath, c.keyType, keyVal.Interface())
			if keyResults != nil {
				results = append(results, keyResults...)
			}

			elemResults := c.PerformTypeValidation(elementPath, c.valueType, val.MapIndex(keyVal).Interface())
			if elemResults != nil {
				results = append(results, elemResults...)
			}
		}
	} else {
		if c.Required() {
			results = append(results,
				NewValidationResult(
					path,
					Error,
					"VALUE_ISNOT_MAP",
					name+" type must be Map",
					convert.Map,
					convert.TypeConverter.ToTypeCode(value),
				),
			)
		}
	}

	return results
}
