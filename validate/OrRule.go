package validate

type OrRule struct {
	rules []IValidationRule
}

func NewOrRule(rules ...IValidationRule) *OrRule {
	return &OrRule{
		rules: rules,
	}
}

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
