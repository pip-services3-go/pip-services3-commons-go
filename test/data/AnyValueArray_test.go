package test_data

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/data"
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
}
