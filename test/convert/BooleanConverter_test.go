package convert

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/convert"
	"github.com/stretchr/testify/assert"
)

func TestToBoolean(t *testing.T) {
	assert.Nil(t, convert.ToNullableBoolean(nil))

	assert.True(t, convert.ToBoolean(true))
	assert.True(t, convert.ToBoolean(1))
	assert.True(t, convert.ToBoolean("True"))
	assert.True(t, convert.ToBoolean("yes"))
	assert.True(t, convert.ToBoolean("1"))
	assert.True(t, convert.ToBoolean("Y"))

	assert.False(t, convert.ToBoolean(false))
	assert.False(t, convert.ToBoolean(0))
	assert.False(t, convert.ToBoolean("False"))
	assert.False(t, convert.ToBoolean("no"))
	assert.False(t, convert.ToBoolean("0"))
	assert.False(t, convert.ToBoolean("N"))

	assert.False(t, convert.ToBoolean(123))
	assert.False(t, convert.ToBoolean(nil))
	assert.True(t, convert.ToBooleanWithDefault("XYZ", true))
}
