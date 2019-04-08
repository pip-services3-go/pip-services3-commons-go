package validate

import "github.com/pip-services3-go/pip-services3-commons-go/reflect"

type PropertiesComparisonRule struct {
	property1 string
	property2 string
	operation string
}

func NewPropertiesComparisonRule(property1 string, operation string, property2 string) *PropertiesComparisonRule {
	return &PropertiesComparisonRule{
		property1: property1,
		property2: property2,
		operation: operation,
	}
}

func (c *PropertiesComparisonRule) Validate(path string, schema ISchema, value interface{}) []*ValidationResult {
	name := path
	if name == "" {
		name = "value"
	}
	value1 := reflect.ObjectReader.GetProperty(value, c.property1)
	value2 := reflect.ObjectReader.GetProperty(value, c.property2)

	if !ObjectComparator.Compare(value1, c.operation, value2) {
		return []*ValidationResult{
			NewValidationResult(
				path,
				Error,
				"PROPERTIES_NOT_MATCH",
				name+" must have "+c.property1+" "+c.operation+" "+c.property2,
				value2,
				value1,
			),
		}
	}
	return nil
}
