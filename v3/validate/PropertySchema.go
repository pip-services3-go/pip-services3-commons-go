package validate

/*
Schema to validate object properties

see
ObjectSchema

Example
var schema = NewObjectSchema()
    .WithProperty(NewPropertySchema("id", TypeCode.String));

schema.Validate({ id: "1", name: "ABC" });       // Result: no errors
schema.Validate({ name: "ABC" });                // Result: no errors
schema.Validate({ id: 1, name: "ABC" });         // Result: id type mismatch
*/
type PropertySchema struct {
	Schema
	name string
	typ  interface{}
}

// Creates a new validation schema and sets its values.
// Returns *PropertySchema
func NewPropertySchema() *PropertySchema {
	c := &PropertySchema{}
	c.Schema = *InheritSchema(c)
	return c
}

// Creates a new validation schema and sets its values.
// see
// IValidationRule
// see
// TypeCode
// Parameters:
// 			- name string
// 			a property name
// 			- type interface{}
// 			a property type
// 			- required bool
// 			true to always require non-null values.
// 			- rules []IValidationRule
// 			a list with validation rules.
// Returns *PropertySchema
func NewPropertySchemaWithRules(name string, typ interface{}, required bool, rules []IValidationRule) *PropertySchema {
	c := &PropertySchema{
		name: name,
		typ:  typ,
	}
	c.Schema = *InheritSchemaWithRules(c, required, rules)
	return c
}

// Gets the property name.
// Returns string
// the property name.
func (c *PropertySchema) Name() string {
	return c.name
}

// Sets the property name.
// Parameters:
// 			- value string
// 			a new property name.
func (c *PropertySchema) SetName(value string) {
	c.name = value
}

// Gets the property type.
// Returns any
// the property type.
func (c *PropertySchema) Type() interface{} {
	return c.typ
}

// Sets a new property type. The type can be defined as type, type name or TypeCode
// Parameters:
// 			- value interface{}
// 			a new property type.
func (c *PropertySchema) SetType(value interface{}) {
	c.typ = value
}

// Validates a given value against the schema and configured validation rules.
// Parameters:
// 			- path string
// 			a dot notation path to the value.
// 			- value interface{}
// 			a value to be validated.
// Return  []*ValidationResult
// a list with validation results to add new results.
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
