package validate

type PropertySchema struct {
	Schema
	name string
	typ  interface{}
}

func NewPropertySchema() *PropertySchema {
	c := &PropertySchema{}
	c.Schema = *InheritSchema(c)
	return c
}

func NewPropertySchemaWithRules(name string, typ interface{}, required bool, rules []IValidationRule) *PropertySchema {
	c := &PropertySchema{
		name: name,
		typ:  typ,
	}
	c.Schema = *InheritSchemaWithRules(c, required, rules)
	return c
}

func (c *PropertySchema) Name() string {
	return c.name
}

func (c *PropertySchema) SetName(value string) {
	c.name = value
}

func (c *PropertySchema) Type() interface{} {
	return c.typ
}

func (c *PropertySchema) SetType(value interface{}) {
	c.typ = value
}

func (c *PropertySchema) PerformValidation(path string, value interface{}) []*ValidationResult {
	if path != "" {
		path = path + "." + c.name
	} else {
		path = c.name
	}

	results := []*ValidationResult{}

	innerResults := c.Schema.PerformValidation(path, value)
	if innerResults != nil {
		results = append(results, innerResults...)
	}

	innerResults = c.Schema.PerformTypeValidation(path, c.Type(), value)
	if innerResults != nil {
		results = append(results, innerResults...)
	}

	return results
}
