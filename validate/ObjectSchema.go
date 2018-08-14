package validate

import (
	"strings"

	"github.com/pip-services-go/pip-services-commons-go/reflect"
)

type ObjectSchema struct {
	Schema
	properties     []*PropertySchema
	allowUndefined bool
}

func NewObjectSchema() *ObjectSchema {
	c := &ObjectSchema{
		allowUndefined: false,
	}
	c.Schema = *InheritSchema(c)
	return c
}

func NewObjectSchemaWithRules(allowUndefined bool, required bool, rules []IValidationRule) *ObjectSchema {
	c := &ObjectSchema{
		allowUndefined: allowUndefined,
	}
	c.Schema = *InheritSchemaWithRules(c, required, rules)
	return c
}

func (c *ObjectSchema) Properties() []*PropertySchema {
	return c.properties
}

func (c *ObjectSchema) SetProperties(value []*PropertySchema) {
	c.properties = value
}

func (c *ObjectSchema) UndefinedAllowed() bool {
	return c.allowUndefined
}

func (c *ObjectSchema) SetUndefinedAllowed(value bool) {
	c.allowUndefined = value
}

func (c *ObjectSchema) AllowUndefined(value bool) *ObjectSchema {
	c.allowUndefined = value
	return c
}

func (c *ObjectSchema) WithProperty(schema *PropertySchema) *ObjectSchema {
	if c.properties == nil {
		c.properties = []*PropertySchema{}
	}
	c.properties = append(c.properties, schema)
	return c
}

func (c *ObjectSchema) WithRequiredProperty(name string, typ interface{}, rules ...IValidationRule) *ObjectSchema {
	schema := NewPropertySchemaWithRules(name, typ, true, rules)
	return c.WithProperty(schema)
}

func (c *ObjectSchema) WithOptionalProperty(name string, typ interface{}, rules ...IValidationRule) *ObjectSchema {
	schema := NewPropertySchemaWithRules(name, typ, false, rules)
	return c.WithProperty(schema)
}

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
