package test_data

import (
	"encoding/json"
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/stretchr/testify/assert"
)

func TestFilterParamsCreate(t *testing.T) {
	filter := data.NewFilterParamsFromTuples(
		"value1", 123,
		"value2", "ABC",
	)

	assert.Equal(t, 2, filter.Len())
}

func TestFilterParamsJsonSerialization(t *testing.T) {
	json1 := []byte("{\"key1\":\"1\",\"key2\":\"A\"}")

	var value *data.StringValueMap
	err := json.Unmarshal(json1, &value)
	assert.Empty(t, err)
	assert.Equal(t, "1", value.Get("key1"))
	assert.Equal(t, "A", value.Get("key2"))

	json2, err2 := json.Marshal(value)
	assert.Empty(t, err2)
	assert.Equal(t, json1, json2)
}
