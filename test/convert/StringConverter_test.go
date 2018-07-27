package convert

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/convert"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	assert.Nil(t, convert.ToNullableString(nil))

	assert.Equal(t, "xyz", convert.ToString("xyz"))
	assert.Equal(t, "123", convert.ToString(123))
	assert.Equal(t, "true", convert.ToString(true))

	value := struct{ prop string }{"xyz"}
	assert.Equal(t, "{xyz}", convert.ToString(value))

	assert.Equal(t, "xyz", convert.ToStringWithDefault(nil, "xyz"))
}
