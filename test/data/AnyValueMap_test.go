package test_data

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/data"
	"github.com/stretchr/testify/assert"
)

func TestAnyValueMapNew(t *testing.T) {
	value := data.NewEmptyAnyValueMap()
	assert.Empty(t, value.GetAsObject("key1"))

	value = data.NewAnyValueMapFromValue(map[string]interface{}{
		"key1": 1,
		"key2": "A",
	})
	assert.Equal(t, int64(1), value.Get("key1"))
	assert.Equal(t, "A", value.Get("key2"))

	value = data.NewAnyValueMapFromMaps(map[string]interface{}{
		"key1": 1,
		"key2": "A",
	})
	assert.Equal(t, 1, value.Get("key1"))
	assert.Equal(t, "A", value.Get("key2"))

	value = data.NewAnyValueMapFromTuples(
		"key1", 1,
		"key2", "A",
	)
	assert.Equal(t, 1, value.Get("key1"))
	assert.Equal(t, "A", value.Get("key2"))
}

func TestAnyValueMapGetAndSet(t *testing.T) {
	value := data.NewEmptyAnyValueMap()
	assert.Empty(t, value.GetAsObject("key1"))

	value.SetAsObject("key1", 1)
	assert.Equal(t, 1, value.GetAsInteger("key1"))
	assert.True(t, 1.0-value.GetAsFloat("key1") < 0.001)
	assert.Equal(t, "1", value.GetAsString("key1"))

	value.Put("key2", "1")
	assert.Equal(t, 1, value.GetAsInteger("key2"))
	assert.True(t, 1.0-value.GetAsFloat("key2") < 0.001)
	assert.Equal(t, "1", value.GetAsString("key2"))

	value.Remove("key2")
	assert.Empty(t, value.GetAsObject("key2"))
}
