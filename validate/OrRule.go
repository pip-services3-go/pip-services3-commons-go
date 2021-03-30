package validate

/*
alidation rule to combine rules with OR logical operation. When one of rules returns no errors, than this rule also returns no errors. When all rules return errors, than the rule returns all errors.

see
IValidationRule

Example
 var schema = NewSchema()
     .WithRule(NewOrRule(
         NewValueComparisonRle("LT", 1),
         NewValueComparisonule("GT", 10)
     ));
 
 schema.Validate();          // Result: no error
 schema.Validate5);          // Result: 5 must be less than 1 or 5 must be more than 10
 schema.Validate(20);        // Result: no error
*/
type OrRule struct {
	rules []IValidationRule
}

// Creates a new validation rule and ses its values
// Parameters:
// 			- rule IValidationRule
// 			a rule to be negaed.
// Returns *OrRule
func NewOrRule(rules ...IValidationRule) *OrRule {
	return &OrRule{
		rules: rules,
	}
}

// Validates a given value against this rule.
// Parameters:
// 			- path string
// 			a dot notation path to th value.
// 			- schema  ISchema
// 			a schema this rule is called from
// 			value interface{}
// 			a valueto be validated.
// Retruns []*ValidationResult
// a list with validation results to add new results.
func (c *OrRule) Validate(path string, schema ISchema, value interface{}) []*ValidationResult {
	if c.rules == nil || len(c.rules) == 0 {
		return nil
	}

	results := []*ValidationResult{}

	for _, rule := range c.rules {
		ruleResults := rule.Validate(path, schema, value)

		if ruleResults == nil || len(ruleResults) == 0 {
			return nil
		}

		results = append(results, ruleResults...)
	}

	return results
}
