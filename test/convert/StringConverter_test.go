package test_convert

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	assert.Nil(t, convert.ToNullableString(nil))

	assert.Equal(t, "xyz", convert.ToString("xyz"))
	assert.Equal(t, "123", convert.ToString(123))
	assert.Equal(t, "true", convert.ToString(true))

	value := struct{ prop string }{"xyz"}
	assert.Equal(t, "{xyz}", convert.ToString(value))

	array1 := []string{"A", "B", "C"}
	assert.Equal(t, "A,B,C", convert.ToString(array1))

	array2 := []int32{1, 2, 3}
	assert.Equal(t, "1,2,3", convert.ToString(array2))

	assert.Equal(t, "xyz", convert.ToStringWithDefault(nil, "xyz"))
}
