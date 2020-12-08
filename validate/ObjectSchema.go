package validate

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/reflect"
)

/*
// Schema to validate user defined objects.

// Example
//  var schema = NewObjectSchema(false)
//      .WithOptionalProperty("id", TypeCode.String)
//      .WithRequiredProperty("name", TypeCode.String);
 
//  schema.validate({ id: "1", name: "ABC" });       // Result: no errors
//  schema.validate({ name: "ABC" });                // Result: no errors
//  schema.validate({ id: 1, name: "ABC" });         // Result: id type mismatch
//  schema.validate({ id: 1, _name: "ABC" });        // Result: name is missing, unexpected _name
//  schema.validate("ABC");                          // Result: type mismatch
*/
type ObjectSchema struct {
	Schema
	properties     []*PropertySchema
	allowUndefined bool
}

// Creates a new validation schema and sets its values.
// Returns *ObjectSchema
func NewObjectSchema() *ObjectSchema {
	c := &ObjectSchema{
		allowUndefined: false,
	}
	c.Schema = *InheritSchema(c)
	return c
}

// Creates a new validation schema and sets its values.
// see
// IValidationRule
// Parameters:
//  - allowUndefined bool
//  true to allow properties undefines in the schema
//  - required bool
//  true to always require non-null values.
//  - rules []IValidationRule
//  a list with validation rules.
// Returns *ObjectSchema
func NewObjectSchemaWithRules(allowUndefined bool, required bool, rules []IValidationRule) *ObjectSchema {
	c := &ObjectSchema{
		allowUndefined: allowUndefined,
	}
	c.Schema = *InheritSchemaWithRules(c, required, rules)
	return c
}

// Gets validation schemas for object properties.
// see
// PropertySchema
// Returns []*PropertySchema
// the list of property validation schemas.
func (c *ObjectSchema) Properties() []*PropertySchema {
	return c.properties
}

// Sets validation schemas for object properties.
// see
// PropertySchema
// Parameters:
//  - value []*PropertySchema
//  a list of property validation schemas.
func (c *ObjectSchema) SetProperties(value []*PropertySchema) {
	c.properties = value
}

// Gets flag to allow undefined properties
// Returns bool
// true to allow undefined properties and false to disallow.
func (c *ObjectSchema) UndefinedAllowed() bool {
	return c.allowUndefined
}

// Gets flag to allow undefined properties
// Parameters:
//  - value bool
//  true to allow undefined properties and false to disallow.
func (c *ObjectSchema) SetUndefinedAllowed(value bool) {
	c.allowUndefined = value
}

// Sets flag to allow undefined properties
// This method returns reference to this exception to implement Builder pattern to chain additional calls.
// Parameters:
//  - value bool
//  true to allow undefined properties and false to disallow.
// Returns *ObjectSchema
// this validation schema.
func (c *ObjectSchema) AllowUndefined(value bool) *ObjectSchema {
	c.allowUndefined = value
	return c
}

// Adds a validation schema for an object property.
// This method returns reference to this exception to implement Builder pattern to chain additional calls.
// see
// PropertySchema
// Parameters:
//  - schema *PropertySchema
//  a property validation schema to be added.
// Returns *ObjectSchema
// this validation schema.
func (c *ObjectSchema) WithProperty(schema *PropertySchema) *ObjectSchema {
	if c.properties == nil {
		c.properties = []*PropertySchema{}
	}
	c.properties = append(c.properties, schema)
	return c
}

// Adds a validation schema for a required object property.
// Parameters:
//  - name string
//  a property name.
//  - type interface{}
//  a property schema or type.
//  - rules ...IValidationRule
//  a list of property validation rules.
// Returns *ObjectSchema
func (c *ObjectSchema) WithRequiredProperty(name string, typ interface{}, rules ...IValidationRule) *ObjectSchema {
	schema := NewPropertySchemaWithRules(name, typ, true, rules)
	return c.WithProperty(schema)
}

// Adds a validation schema for an optional object property.
// Parameters:
//  - name string
//  a property name.
//  - type interface{}
//  a property schema or type.
//  - rules ...IValidationRule
//   a list of property validation rules.
// Returns *ObjectSchema
func (c *ObjectSchema) WithOptionalProperty(name string, typ interface{}, rules ...IValidationRule) *ObjectSchema {
	schema := NewPropertySchemaWithRules(name, typ, false, rules)
	return c.WithProperty(schema)
}

// Validates a given value against the schema and configured validation rules.
// Parameters:
//  - path string
//  a dot notation path to the value.
//  - value interface{}
//  a value to be validated.
// Return []*ValidationResult
// a list with validation results to add new results.
func (c *ObjectSchema) PerformValidation(path string, value interface{}) []*ValidationResult {
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
	properties := reflect.ObjectReader.GetProperties(value)

	if properties != nil {
		for _, propertySchema := range c.properties {
			processedName := ""

			for propertyName, propertyValue := range properties {
				if strings.EqualFold(propertySchema.Name(), propertyName) {
					propResults := propertySchema.PerformValidation(path, propertyValue)
					if propResults != nil {
						results = append(results, propResults...)
					}
					processedName = propertyName
					break
				}
			}

			if processedName != "" {
				delete(properties, processedName)
			} else {
				propResults := propertySchema.PerformValidation(path, nil)
				if propResults != nil {
					results = append(results, propResults...)
				}
			}
		}
	}

	if !c.allowUndefined {
		for propertyName := range properties {
			propertyPath := propertyName
			if path != "" {
				propertyPath = path + "." + propertyName
			}

			results = append(results, NewValidationResult(
				propertyPath,
				Warning,
				"UNEXPECTED_PROPERTY",
				name+" contains unexpected property "+propertyName,
				nil,
				propertyName,
			))
		}
	}

	return results
}
