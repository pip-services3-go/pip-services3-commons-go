package test_validate

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/validate"
	"github.com/stretchr/testify/assert"
)

func TestValidateComparisonRule(t *testing.T) {
	schema := validate.NewSchema().
		WithRule(validate.NewValueComparisonRule("EQ", 123))
	results := schema.Validate(123)
	assert.Equal(t, 0, len(results))

	results = schema.Validate(423)
	assert.Equal(t, 1, len(results))
}
