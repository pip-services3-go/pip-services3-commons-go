package test_validate

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/validate"
	"github.com/stretchr/testify/assert"
)

func TestEmptySchema(t *testing.T) {
	schema := validate.NewSchema()
	results := schema.Validate(nil)
	assert.Equal(t, 0, len(results))
}

func TestSchemaRequired(t *testing.T) {
	schema := validate.NewSchema().MakeRequired()
	results := schema.Validate(nil)
	assert.Equal(t, 1, len(results))
}
