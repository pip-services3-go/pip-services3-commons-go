package test_data

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/stretchr/testify/assert"
)

func TestAnyValueArrayCreate(t *testing.T) {
	array := data.NewEmptyAnyValueArray()
	assert.Equal(t, 0, array.Len())

	array = data.NewAnyValueArray([]interface{}{1, 2, 3})
	assert.Equal(t, 3, array.Len())
	assert.Equal(t, "1,2,3", array.String())

	array = data.NewAnyValueArrayFromString("Fatal,Error,Info,", ",", true)
	assert.Equal(t, 3, array.Len())

	array = data.NewAnyValueArray([]interface{}{1, 2, 3})
	assert.Equal(t, 3, array.Len())
	assert.True(t, array.Contains(1))

	array = data.NewAnyValueArrayFromValue([]interface{}{1, 2, 3})
	assert.Equal(t, 3, array.Len())
	assert.Equal(t, int64(1), array.Get(0))

	array2 := data.NewAnyValueArray([]interface{}{1, map[string]interface{}{"number": "123.456"}, "2018-01-01"})
	value2 := array2.GetAsMapWithDefault(1, data.NewAnyValueMap(map[string]interface{}{"key1": 1}))
	assert.Equal(t, value2.GetAsString("number"), "123.456")
}
