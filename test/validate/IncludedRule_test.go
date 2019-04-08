package test_validate

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/validate"
	"github.com/stretchr/testify/assert"
)

func TestIncludedRule(t *testing.T) {
	schema := validate.NewSchema().
		WithRule(validate.NewIncludedRule("AAA", "BBB", "CCC", nil))

	results := schema.Validate("AAA")
	assert.Equal(t, 0, len(results))

	results = schema.Validate("ABC")
	assert.Equal(t, 1, len(results))
}
