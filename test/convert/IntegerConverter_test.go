package test_convert

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/convert"
	"github.com/stretchr/testify/assert"
)

func TestToInteger(t *testing.T) {
	assert.Nil(t, convert.ToNullableInteger(nil))

	assert.Equal(t, int(123), convert.ToInteger(123))
	assert.Equal(t, int(123), convert.ToInteger(123.456))
	assert.Equal(t, int(123), convert.ToInteger("123"))
	assert.Equal(t, int(123), convert.ToInteger("123.456"))

	assert.Equal(t, int(123), convert.ToIntegerWithDefault(nil, 123))
	assert.Equal(t, int(0), convert.ToIntegerWithDefault(false, 123))
	assert.Equal(t, int(123), convert.ToIntegerWithDefault("ABC", 123))
}
