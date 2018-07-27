package convert

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/convert"
	"github.com/stretchr/testify/assert"
)

func TestToLong(t *testing.T) {
	assert.Nil(t, convert.ToNullableLong(nil))

	assert.Equal(t, int64(123), convert.ToLong(123))
	assert.Equal(t, int64(123), convert.ToLong(123.456))
	assert.Equal(t, int64(123), convert.ToLong("123"))
	assert.Equal(t, int64(123), convert.ToLong("123.456"))

	assert.Equal(t, int64(123), convert.ToLongWithDefault(nil, 123))
	assert.Equal(t, int64(0), convert.ToLongWithDefault(false, 123))
	assert.Equal(t, int64(123), convert.ToLongWithDefault("ABC", 123))
}
