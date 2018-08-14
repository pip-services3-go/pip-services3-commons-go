package validate

type NotRule struct {
	rule IValidationRule
}

func NewNotRule(rule IValidationRule) *NotRule {
	return &NotRule{
		rule: rule,
	}
}

func (c *NotRule) Validate(path string, schema ISchema, value interface{}) []*ValidationResult {
	if c.rule == nil {
		return nil
	}

	name := path
	if name == "" {
		name = "value"
	}

	results := c.rule.Validate(path, schema, value)

	if results != nil && len(results) > 0 {
		return nil
	}

	return []*ValidationResult{
		NewValidationResult(
			path,
			Error,
			"NOT_FAILED",
			"Negative check for "+name+" failed",
			nil,
			nil,
		),
	}
}
