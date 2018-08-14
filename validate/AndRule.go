package validate

type AndRule struct {
	rules []IValidationRule
}

func NewAndRule(rules ...IValidationRule) *AndRule {
	return &AndRule{
		rules: rules,
	}
}

func (c *AndRule) Validate(path string, schema ISchema, value interface{}) []*ValidationResult {
	if c.rules == nil || len(c.rules) == 0 {
		return nil
	}

	results := []*ValidationResult{}

	for _, rule := range c.rules {
		ruleResults := rule.Validate(path, schema, value)
		if ruleResults != nil {
			results = append(results, ruleResults...)
		}
	}

	return results
}
