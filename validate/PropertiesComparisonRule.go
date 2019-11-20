package validate

import "github.com/pip-services3-go/pip-services3-commons-go/reflect"

/*
Validation rule that compares two object properties.

see
IValidationRule

Example
var schema = NewObjectSchema()
    .WithRule(NewPropertyComparisonRule("field1", "NE", "field2"));

schema.Validate({ field1: 1, field2: 2 });       // Result: no errors
schema.Validate({ field1: 1, field2: 1 });       // Result: field1 shall not be equal to field2
schema.Validate({});                             // Result: no errors
*/
type PropertiesComparisonRule struct {
	property1 string
	property2 string
	operation string
}

// Creates a new validation rule and sets its arguments.
// see
// ObjectComparator.compare
// Parameters:
// 			- property1 string
// 			a name of the first property to compare.
// 			- operation string
// 			a comparison operation: "==" ("=", "EQ"), "!= " ("<>", "NE"); "<"/">" ("LT"/"GT"), "<="/">=" ("LE"/"GE"); "LIKE".
// 			property2 string
// 			a name of the second property to compare.
// Returns *PropertiesComparisonRule
func NewPropertiesComparisonRule(property1 string, operation string, property2 string) *PropertiesComparisonRule {
	return &PropertiesComparisonRule{
		property1: property1,
		property2: property2,
		operation: operation,
	}
}

// Validates a given value against this rule.
// Parameters:
// 			- path string
// 			a dot notation path to the value.
// 			- schema Schema
// 			a schema this rule is called from
// 			- value interface{}
// 			a value to be validated.
// Return []*ValidationResult
// a list with validation results to add new results.
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
