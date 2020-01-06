package validate

import "github.com/pip-services3-go/pip-services3-commons-go/v3/convert"

/*

Validation rule that compares value to a constant.

see
IValidationRule

Example
var schema = NewSchema()
    .WithRule(NewValueComparisonRule("EQ", 1));

schema.Validate(1);          // Result: no errors
schema.Validate(2);          // Result: 2 is not equal to 1
*/
type ValueComparisonRule struct {
	value     interface{}
	operation string
}

// Creates a new validation rule and sets its values.
// Parameters:
// 			- operation string
// 			a comparison operation: "==" ("=", "EQ"), "!= " ("<>", "NE"); "<"/">" ("LT"/"GT"), "<="/">=" ("LE"/"GE"); "LIKE".
// 			- value interface{}
// 			a constant value to compare to
// Returns *ValueComparisonRule
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
