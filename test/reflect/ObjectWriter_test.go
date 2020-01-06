package test_reflect

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/reflect"
	"github.com/stretchr/testify/assert"
)

func TestWriterSetProperty(t *testing.T) {
	obj := NewTestClass()
	assert.Equal(t, true, reflect.ObjectReader.GetProperty(obj, "RootPublicProperty"))
	reflect.ObjectWriter.SetProperty(obj, "RootPublicProperty", false)
	assert.Equal(t, false, reflect.ObjectReader.GetProperty(obj, "RootPublicProperty"))
	assert.Equal(t, "BBB", reflect.ObjectReader.GetProperty(obj, "PublicField"))
	reflect.ObjectWriter.SetProperty(obj, "PublicField", "XYZ")
	assert.Equal(t, "XYZ", reflect.ObjectReader.GetProperty(obj, "PublicField"))

	dict := map[string]interface{}{
		"1": "AAA",
		"A": 111,
	}
	assert.Equal(t, "AAA", reflect.ObjectReader.GetProperty(dict, "1"))
	reflect.ObjectWriter.SetProperty(dict, "1", "XYZ")
	assert.Equal(t, "XYZ", reflect.ObjectReader.GetProperty(dict, "1"))
	assert.Equal(t, 111, reflect.ObjectReader.GetProperty(dict, "A"))
	reflect.ObjectWriter.SetProperty(dict, "A", 777)
	assert.Equal(t, 777, reflect.ObjectReader.GetProperty(dict, "A"))

	list := []interface{}{"BBB", 222}
	assert.Equal(t, "BBB", reflect.ObjectReader.GetProperty(list, "0"))
	reflect.ObjectWriter.SetProperty(list, "0", "XYZ")
	assert.Equal(t, "XYZ", reflect.ObjectReader.GetProperty(list, "0"))
	assert.Equal(t, 222, reflect.ObjectReader.GetProperty(list, "1"))
	reflect.ObjectWriter.SetProperty(list, "1", 777)
	assert.Equal(t, 777, reflect.ObjectReader.GetProperty(list, "1"))
}

func TestWriterSetProperties(t *testing.T) {
	obj := NewTestClass()
	assert.Equal(t, true, reflect.ObjectReader.GetProperty(obj, "RootPublicProperty"))
	assert.Equal(t, "BBB", reflect.ObjectReader.GetProperty(obj, "PublicField"))

	values := map[string]interface{}{
		"RootPublicProperty": false,
		"PublicField":        "XYZ",
	}
	reflect.ObjectWriter.SetProperties(obj, values)

	assert.Equal(t, false, reflect.ObjectReader.GetProperty(obj, "RootPublicProperty"))
	assert.Equal(t, "XYZ", reflect.ObjectReader.GetProperty(obj, "PublicField"))

	dict := map[string]interface{}{
		"1": "AAA",
		"A": 111,
	}
	assert.Equal(t, "AAA", reflect.ObjectReader.GetProperty(dict, "1"))
	assert.Equal(t, 111, reflect.ObjectReader.GetProperty(dict, "A"))

	values = map[string]interface{}{
		"1": "XYZ",
		"A": 777,
	}
	reflect.ObjectWriter.SetProperties(dict, values)

	assert.Equal(t, "XYZ", reflect.ObjectReader.GetProperty(dict, "1"))
	assert.Equal(t, 777, reflect.ObjectReader.GetProperty(dict, "A"))

	list := []interface{}{"BBB", 222}
	assert.Equal(t, "BBB", reflect.ObjectReader.GetProperty(list, "0"))
	assert.Equal(t, 222, reflect.ObjectReader.GetProperty(list, "1"))

	values = map[string]interface{}{
		"0": "XYZ",
		"1": 777,
	}
	reflect.ObjectWriter.SetProperties(list, values)

	assert.Equal(t, "XYZ", reflect.ObjectReader.GetProperty(list, "0"))
	assert.Equal(t, 777, reflect.ObjectReader.GetProperty(list, "1"))
}
