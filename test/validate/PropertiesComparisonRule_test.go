package test_validate

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/validate"
	"github.com/stretchr/testify/assert"
)

func TestPropertiesComparisonRule(t *testing.T) {
	obj := &TestClass{}
	schema := validate.NewSchema().WithRule(validate.NewPropertiesComparisonRule("StringField1", "EQ", "StringField2"))

	obj.StringField1 = "ABC"
	obj.StringField2 = "ABC"
	results := schema.Validate(obj)
	assert.Equal(t, 0, len(results))

	obj.StringField2 = "XYZ"
	results = schema.Validate(obj)
	assert.Equal(t, 1, len(results))
}
