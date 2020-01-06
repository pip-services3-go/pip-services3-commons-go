package validate

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/convert"
)

/*
Validation rule to check that value is included into the list of constants.

see
IValidationRule

Example
var schema = NewSchema()
    .WithRule(NewIncludedRule(1, 2, 3));

schema.Validate(2);      // Result: no errors
schema.Validate(10);     // Result: 10 must be one of 1, 2, 3
*/
type IncludedRule struct {
	values []interface{}
}

// Creates a new validation rule and sets its values.
// Parameters:
// 			 - values ...interface{}
// 			a list of constants that value must be included to
// Returns *IncludedRule
func NewIncludedRule(values ...interface{}) *IncludedRule {
	return &IncludedRule{
		values: values,
	}
}

// Validates a given value against this rule.
// Parameters:
// 			- path string
// 			a dot notation path to the value.
// 			- schema ISchema
// 			a schema this rule is called from
// 			- value interface{}y
// 			a value to be validated.
// 	Returns []*ValidationResult
// a list with validation results to add new results.
func (c *IncludedRule) Validate(path string, schema ISchema, value interface{}) []*ValidationResult {
	if c.values == nil || len(c.values) == 0 {
		return nil
	}

	name := path
	if name == "" {
		name = "value"
	}

	found := false

	for _, thisValue := range c.values {
		if ObjectComparator.AreEqual(value, thisValue) {
			found = true
			break
		}
	}

	if !found {
		expectedValues := strings.Builder{}
		for _, thisValue := range c.values {
			if expectedValues.Len() > 0 {
				expectedValues.WriteString(",")
			}
			expectedValues.WriteString(convert.StringConverter.ToString(thisValue))
		}

		return []*ValidationResult{
			NewValidationResult(
				path,
				Error,
				"VALUE_NOT_INCLUDED",
				name+" must be one of "+expectedValues.String(),
				c.values,
				nil,
			),
		}
	}

	return nil
}
