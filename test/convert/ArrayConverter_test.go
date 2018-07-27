package convert

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/convert"
	"github.com/stretchr/testify/assert"
)

func TestToNullableArray(t *testing.T) {
	assert.Nil(t, convert.ToNullableArray(nil))

	a := *convert.ToNullableArray(2)
	assert.Len(t, a, 1)
	assert.Equal(t, int64(2), a[0])

	array := []int{1, 2}
	a = *convert.ToNullableArray(array)
	assert.Len(t, a, 2)
	assert.Equal(t, int64(1), a[0])
	assert.Equal(t, int64(2), a[1])

	stringArray := []string{"ab", "cd"}
	a = *convert.ToNullableArray(stringArray)
	assert.Len(t, a, 2)
	assert.Equal(t, "ab", a[0])
	assert.Equal(t, "cd", a[1])
}

func TestToArray(t *testing.T) {
	a := convert.ToArray(nil)
	assert.Len(t, a, 0)

	a = convert.ToArray(2)
	assert.Len(t, a, 1)
	assert.Equal(t, int64(2), a[0])

	array := []int{1, 2}
	a = convert.ToArray(array)
	assert.Len(t, a, 2)
	assert.Equal(t, int64(1), a[0])
	assert.Equal(t, int64(2), a[1])

	stringArray := []string{"ab", "cd"}
	a = convert.ToArray(stringArray)
	assert.Len(t, a, 2)
	assert.Equal(t, "ab", a[0])
	assert.Equal(t, "cd", a[1])
}
