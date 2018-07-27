package data

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/data"
	"github.com/stretchr/testify/assert"
)

func TestProjectionParamsFromNull(t *testing.T) {
	parameters := data.NewProjectionParamsFromValue(nil)

	assert.Equal(t, 0, parameters.Length())
}

func TestProjectionParamsFromValue(t *testing.T) {
	parameters := data.NewProjectionParamsFromValue([]interface{}{"field1", "field2", "field3"})

	assert.Equal(t, 3, parameters.Length())
	assert.Equal(t, "field1", parameters.Get(0))
	assert.Equal(t, "field2", parameters.Get(1))
	assert.Equal(t, "field3", parameters.Get(2))
}

func TestParseProjectionParams(t *testing.T) {
	parameters := data.ParseProjectionParams("field1", "field2", "field3")

	assert.Equal(t, 3, parameters.Length())
	assert.Equal(t, "field1", parameters.Get(0))
	assert.Equal(t, "field2", parameters.Get(1))
	assert.Equal(t, "field3", parameters.Get(2))

	parameters = data.ParseProjectionParams("field1,field2, field3")

	assert.Equal(t, 3, parameters.Length())
	assert.Equal(t, "field1", parameters.Get(0))
	assert.Equal(t, "field2", parameters.Get(1))
	assert.Equal(t, "field3", parameters.Get(2))

	parameters = data.ParseProjectionParams("object1(field1)", "object2(field1, field2)", "field3")

	assert.Equal(t, 4, parameters.Length())
	assert.Equal(t, "object1.field1", parameters.Get(0))
	assert.Equal(t, "object2.field1", parameters.Get(1))
	assert.Equal(t, "object2.field2", parameters.Get(2))
	assert.Equal(t, "field3", parameters.Get(3))

	parameters = data.ParseProjectionParams("object1(object2(field1,field2,object3(field1)))")

	assert.Equal(t, 3, parameters.Length())
	assert.Equal(t, "object1.object2.field1", parameters.Get(0))
	assert.Equal(t, "object1.object2.field2", parameters.Get(1))
	assert.Equal(t, "object1.object2.object3.field1", parameters.Get(2))

	parameters = data.ParseProjectionParams("object1(field1, object2(field1, field2, field3, field4), field3)", "field2")

	assert.Equal(t, 7, parameters.Length())
	assert.Equal(t, "object1.field1", parameters.Get(0))
	assert.Equal(t, "object1.object2.field1", parameters.Get(1))
	assert.Equal(t, "object1.object2.field2", parameters.Get(2))
	assert.Equal(t, "object1.object2.field3", parameters.Get(3))
	assert.Equal(t, "object1.object2.field4", parameters.Get(4))
	assert.Equal(t, "object1.field3", parameters.Get(5))
	assert.Equal(t, "field2", parameters.Get(6))

	parameters = data.ParseProjectionParams("object1(field1, object2(field1), field3)", "field2")

	assert.Equal(t, 4, parameters.Length())
	assert.Equal(t, "object1.field1", parameters.Get(0))
	assert.Equal(t, "object1.object2.field1", parameters.Get(1))
	assert.Equal(t, "object1.field3", parameters.Get(2))
	assert.Equal(t, "field2", parameters.Get(3))

	parameters = data.ParseProjectionParams("object1(field1, object2(field1, field2, object3(field1), field4), field3)", "field2")

	assert.Equal(t, 7, parameters.Length())
	assert.Equal(t, "object1.field1", parameters.Get(0))
	assert.Equal(t, "object1.object2.field1", parameters.Get(1))
	assert.Equal(t, "object1.object2.field2", parameters.Get(2))
	assert.Equal(t, "object1.object2.object3.field1", parameters.Get(3))
	assert.Equal(t, "object1.object2.field4", parameters.Get(4))
	assert.Equal(t, "object1.field3", parameters.Get(5))
	assert.Equal(t, "field2", parameters.Get(6))

	parameters = data.ParseProjectionParams("object1(object2(object3(field1)), field2)", "field2")

	assert.Equal(t, 3, parameters.Length())
	assert.Equal(t, "object1.object2.object3.field1", parameters.Get(0))
	assert.Equal(t, "object1.field2", parameters.Get(1))
	assert.Equal(t, "field2", parameters.Get(2))

	parameters = data.ParseProjectionParams("field1,object1(field1),object2(field1,field2),object3(field1),field2,field3")

	assert.Equal(t, 7, parameters.Length())
	assert.Equal(t, "field1", parameters.Get(0))
	assert.Equal(t, "object1.field1", parameters.Get(1))
	assert.Equal(t, "object2.field1", parameters.Get(2))
	assert.Equal(t, "object2.field2", parameters.Get(3))
	assert.Equal(t, "object3.field1", parameters.Get(4))
	assert.Equal(t, "field2", parameters.Get(5))
	assert.Equal(t, "field3", parameters.Get(6))
}
