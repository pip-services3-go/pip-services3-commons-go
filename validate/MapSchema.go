package validate

import (
	refl "reflect"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/reflect"
)

/*
Schema to validate maps.

Example
 var schema = NewMapSchema(TypeCode.String, TypeCode.Integer);
 
 schema.Validate({ "key1": "A", "key2": "B" });       // Result: no errors
 schema.Validate({ "key1": 1, "key2": 2 });           // Result: element type mismatch
 schema.Validate([ 1, 2, 3 ]);                        // Result: type mismatch
*/
type MapSchema struct {
	Schema
	keyType   interface{}
	valueType interface{}
}

// Creates a new instance of validation schema and sets its values.
// see
// IValidationRule
// see
// TypeCode
// Parameters:
// 			- keyType interface{}
// 			a type of map keys. Null means that keys may have any type.
// 			- valueType interface{}
// 			a type of map values. Null means that values may have any type.
// Returns *MapSchema
func NewMapSchema(keyType interface{}, valueType interface{}) *MapSchema {
	c := &MapSchema{
		keyType:   keyType,
		valueType: valueType,
	}
	c.Schema = *InheritSchema(c)
	return c
}

// Creates a new instance of validation schema and sets its values.
// see
// IValidationRule
// see
// TypeCode
// Parameters:
// 			 - keyType interface{}
// 			 a type of map keys. Null means that keys may have any type.
// 			 - valueType interface{}
// 			 a type of map values. Null means that values may have any type.
// 			 - required: boolean
// 			 true to always require non-null values.
// 			 - rules: []IValidationRule
// 			 a list with validation rules.
// Returns *MapSchema
func NewMapSchemaWithRules(keyType interface{}, valueType interface{}, required bool, rules []IValidationRule) *MapSchema {
	c := &MapSchema{
		keyType:   keyType,
		valueType: valueType,
	}
	c.Schema = *InheritSchemaWithRules(c, required, rules)
	return c
}

// Gets the type of map keys. Null means that keys may have any type.
// Returns interface{}
// the type of map keys.
func (c *MapSchema) KeyType() interface{} {
	return c.keyType
}

// Sets the type of map keys. Null means that keys may have any type.
// Parameters:
// 			- value interface{}
// 			a type of map keys.
func (c *MapSchema) SetKeyType(value interface{}) {
	c.keyType = value
}

// Gets the type of map values. Null means that values may have any type.
// Returns interface{}
// the type of map values.
func (c *MapSchema) ValueType() interface{} {
	return c.valueType
}

// Sets the type of map values. Null means that values may have any type.
// Parameters:
// 			- value interface{}
// 			a type of map values.
func (c *MapSchema) SetValueType(value interface{}) {
	c.valueType = value
}

// Validates a given value against the schema and configured validation rules.
// Parameters:
// 			 - path string
// 			 a dot notation path to the value.
// 			 - value interface{}
// 			 a value to be validated.
// REturns  []*ValidationResult[]
// a list with validation results to add new results.
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
