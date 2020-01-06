package test_convert

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/convert"
	"github.com/stretchr/testify/assert"
)

func TestToFloat(t *testing.T) {
	assert.Nil(t, convert.ToNullableFloat(nil))

	assert.Equal(t, float32(123.), convert.ToFloat(123))
	assert.Equal(t, float32(123.456), convert.ToFloat(123.456))
	assert.Equal(t, float32(123.), convert.ToFloat("123"))
	assert.Equal(t, float32(123.456), convert.ToFloat("123.456"))

	assert.Equal(t, float32(123.), convert.ToFloatWithDefault(nil, 123))
	assert.Equal(t, float32(0.), convert.ToFloatWithDefault(false, 123))
	assert.Equal(t, float32(123.), convert.ToFloatWithDefault("ABC", 123))
}
