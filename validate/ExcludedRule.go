package validate

import (
	"strings"

	"github.com/pip-services-go/pip-services-commons-go/convert"
)

type ExcludedRule struct {
	values []interface{}
}

func NewExcludedRule(values ...interface{}) *ExcludedRule {
	return &ExcludedRule{
		values: values,
	}
}

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
