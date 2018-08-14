package validate

import (
	"strings"

	"github.com/pip-services-go/pip-services-commons-go/reflect"
)

type OnlyOneExistsRule struct {
	properties []string
}

func NewOnlyOneExistsRule(properties ...string) *OnlyOneExistsRule {
	return &OnlyOneExistsRule{
		properties: properties,
	}
}

func (c *OnlyOneExistsRule) Validate(path string, schema ISchema, value interface{}) []*ValidationResult {
	name := path
	if name == "" {
		name = "value"
	}

	found := 0

	for _, property := range c.properties {
		propertyValue := reflect.ObjectReader.GetProperty(value, property)
		if propertyValue != nil {
			found++
		}
	}

	if found == 0 {
		return []*ValidationResult{
			NewValidationResult(
				path,
				Error,
				"VALUE_NULL",
				name+" must have at least one property from "+strings.Join(c.properties, ","),
				c.properties,
				nil,
			),
		}
	} else if found > 1 {
		return []*ValidationResult{
			NewValidationResult(
				path,
				Error,
				"VALUE_ONLY_ONE",
				name+" must have only one property from "+strings.Join(c.properties, ","),
				c.properties,
				nil,
			),
		}
	}

	return nil
}
