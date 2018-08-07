package test_convert

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/convert"
	"github.com/stretchr/testify/assert"
)

func TestObjectToMap(t *testing.T) {
	assert.Nil(t, convert.ToNullableMap(nil))

	v1 := struct{ value1, value2 float64 }{123, 234}
	m := convert.ToMap(v1)
	assert.Len(t, m, 2)
	assert.Equal(t, 123., m["value1"])
	assert.Equal(t, 234., m["value2"])

	v2 := map[string]interface{}{"value1": 123}
	m = convert.ToMap(v2)
	assert.Len(t, m, 1)
	assert.Equal(t, int64(123), m["value1"])
}

func TestToNullableMap(t *testing.T) {
	assert.Nil(t, convert.ToNullableMap(nil))
	assert.Nil(t, convert.ToNullableMap(5))

	array := []int{1, 2}

	m := *convert.ToNullableMap(array)
	assert.Len(t, m, 2)
	assert.Equal(t, int64(1), m["0"])
	assert.Equal(t, int64(2), m["1"])

	values := []string{"ab", "cd"}
	m = *convert.ToNullableMap(values)
	assert.Len(t, m, 2)
	assert.Equal(t, "ab", m["0"])
	assert.Equal(t, "cd", m["1"])

	hash := map[int]string{}
	hash[8] = "title 8"
	hash[11] = "title 11"
	m = *convert.ToNullableMap(hash)
	assert.Len(t, m, 2)
	assert.Equal(t, "title 8", m["8"])
	assert.Equal(t, "title 11", m["11"])
}
