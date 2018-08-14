package validate

import (
	refl "reflect"

	"github.com/pip-services-go/pip-services-commons-go/convert"
	"github.com/pip-services-go/pip-services-commons-go/errors"
	"github.com/pip-services-go/pip-services-commons-go/reflect"
)

type Schema struct {
	base     ISchema
	required bool
	rules    []IValidationRule
}

func NewSchema() *Schema {
	c := &Schema{
		required: false,
		rules:    []IValidationRule{},
	}
	c.base = c
	return c
}

func NewSchemaWithRules(required bool, rules []IValidationRule) *Schema {
	c := &Schema{
		required: required,
		rules:    rules,
	}
	c.base = c
	return c
}

func InheritSchema(base ISchema) *Schema {
	c := &Schema{
		required: false,
		rules:    []IValidationRule{},
	}
	c.base = base
	return c
}

func InheritSchemaWithRules(base ISchema, required bool, rules []IValidationRule) *Schema {
	c := &Schema{
		required: required,
		rules:    rules,
	}
	c.base = base
	return c
}

func (c *Schema) Required() bool {
	return c.required
}

func (c *Schema) SetRequired(value bool) {
	c.required = value
}

func (c *Schema) Rules() []IValidationRule {
	return c.rules
}

func (c *Schema) SetRules(value []IValidationRule) {
	c.rules = value
}

func (c *Schema) MakeRequired() *Schema {
	c.required = true
	return c
}

func (c *Schema) MakeOptional() *Schema {
	c.required = false
	return c
}

func (c *Schema) WithRule(rule IValidationRule) *Schema {
	if c.rules == nil {
		c.rules = []IValidationRule{}
	}
	c.rules = append(c.rules, rule)
	return c
}

func (c *Schema) PerformValidation(path string, value interface{}) []*ValidationResult {
	results := []*ValidationResult{}

	name := path
	if name == "" {
		name = "value"
	}

	if value == nil {
		if c.Required() {
			results = append(results, NewValidationResult(
				path,
				Error,
				"VALUE_IS_NULL",
				name+" must not be null",
				"NOT NULL",
				nil,
			))
		}
	} else {
		value = reflect.ObjectReader.GetValue(value)

		// Check validation rules
		if c.rules != nil {
			for _, rule := range c.rules {
				ruleResults := rule.Validate(path, c, value)
				if ruleResults != nil {
					results = append(results, ruleResults...)
				}
			}
		}
	}

	return results
}

func (c *Schema) typeToString(typ interface{}) string {
	if typ == nil {
		return "unknown"
	}
	typeCode := convert.IntegerConverter.ToNullableInteger(typ)
	if typeCode != nil {
		return convert.TypeConverter.ToString(convert.TypeCode(*typeCode))
	}
	return convert.StringConverter.ToString(typ)
}

func (c *Schema) PerformTypeValidation(path string, typ interface{}, value interface{}) []*ValidationResult {
	results := []*ValidationResult{}

	// If type it not defined then skip
	if typ == nil {
		return results
	}

	// Perform validation against the schema
	schema, ok := typ.(ISchema)
	if ok {
		schemaResults := schema.PerformValidation(path, value)
		if schemaResults != nil {
			results = append(results, schemaResults...)
		}
		return results
	}

	// If value is null then skip
	value = reflect.ObjectReader.GetValue(value)
	if value == nil {
		return results
	}

	name := path
	if name == "" {
		name = "value"
	}
	valueType := refl.TypeOf(value)
	valueTypeCode := convert.TypeConverter.ToTypeCode(value)

	// Match types
	if reflect.TypeMatcher.MatchType(typ, valueType) {
		return results
	}

	results = append(results,
		NewValidationResult(
			path,
			Error,
			"TYPE_MISMATCH",
			name+" type must be "+c.typeToString(typ)+" but found "+c.typeToString(valueType),
			typ,
			convert.TypeConverter.ToString(valueTypeCode),
		),
	)
	return results
}

func (c *Schema) Validate(value interface{}) []*ValidationResult {
	return c.base.PerformValidation("", value)
}

func (c *Schema) ValidateAndReturnError(correlationId string, value interface{}, strict bool) *errors.ApplicationError {
	results := c.Validate(value)
	return NewValidationErrorFromResults(correlationId, results, strict)
}

func (c *Schema) ValidateAndThrowError(correlationId string, value interface{}, strict bool) {
	results := c.Validate(value)
	ThrowValidationErrorIfNeeded(correlationId, results, strict)
}
