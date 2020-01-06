package validate

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/convet"
)

/*
Validation rule to check that value is excluded from the list of constants.

see
IValidationRule

Example:
schema := NewSchema()
    .WithRule(NewExcludedRule(1, 2, 3));

schema.Validate(2);      // Result: 2 must not be one of 1, 2, 3
schema.Validate(10);     // Result: no errors
*/
type ExcludedRule struct {
	values []interface{}
}

// Creates a new validation rule and sets its values.
// Parameters:
// values ...interface{}
// a list of constants that value must be excluded from
// Returns *ExcludedRule
func NewExcludedRule(values ...interface{}) *ExcludedRule {
	return &ExcludedRule{
		values: values,
	}
}

// Validates the given value. None of the values set in this ExcludedRule object must exist in the value that is given for validation to pass.
// Parameters:
// 			-path string
// 			the dot notation path to the value that is to be validated.
// 			- schema ISchema
// 			(not used in this implementation).
// 			- value interface{}
// 			the value that is to be validated.
// Return  [*]ValidationResult
// the results of the validation.
func (c *ExcludedRule) Validate(path string, schema ISchema, value interface{}) []*ValidationResult {
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

	if found {
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
				"VALUE_INCLUDED",
				name+" must not be one of "+expectedValues.String(),
				c.values,
				nil,
			),
		}
	}

	return nil
}
