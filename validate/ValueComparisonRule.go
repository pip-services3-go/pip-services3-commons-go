package validate

import "github.com/pip-services-go/pip-services-commons-go/convert"

type ValueComparisonRule struct {
	value     interface{}
	operation string
}

func NewValueComparisonRule(operation string, value interface{}) *ValueComparisonRule {
	return &ValueComparisonRule{
		value:     value,
		operation: operation,
	}
}

func (c *ValueComparisonRule) Validate(path string, schema ISchema, value interface{}) []*ValidationResult {
	name := path
	if name == "" {
		name = "value"
	}

	if !ObjectComparator.Compare(value, c.operation, c.value) {
		expectedValue := convert.StringConverter.ToString(c.value)
		actualValue := convert.StringConverter.ToString(value)

		return []*ValidationResult{
			NewValidationResult(
				path,
				Error,
				"BAD_VALUE",
				name+" must "+c.operation+" "+expectedValue+" but found "+actualValue,
				c.operation+" "+expectedValue,
				value,
			),
		}
	}

	return nil
}
