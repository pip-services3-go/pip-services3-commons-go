package run

import (
	"github.com/pip-services-go/pip-services-commons-go/convert"
	"github.com/pip-services-go/pip-services-commons-go/data"
)

type Parameters struct {
	data.AnyValueMap
}

func NewEmptyParameters() *Parameters {
	return &Parameters{
		AnyValueMap: *data.NewEmptyAnyValueMap(),
	}
}

func NewParameters(values map[string]interface{}) *Parameters {
	c := &Parameters{
		AnyValueMap: *data.NewEmptyAnyValueMap(),
	}
	c.Append(values)
	return c
}

func (c *Parameters) Get(key string) interface{} {
	// Todo: Make this method recursive
	// if key == "" {
	// 	return nil
	// } else if strings.Index(key, ".") > 0 {
	// 	return RecursiveObjectReader.GetProperty(c, key)
	// } else {
	// 	return (*c)[key]
	// }
	return c.AnyValueMap.Get(key)
}

func (c *Parameters) Put(key string, value interface{}) {
	// Todo: Make this method recursive
	c.AnyValueMap.Put(key, value)
}

func (c *Parameters) Remove(key string) {
	// Todo: Make this method recursive
	c.AnyValueMap.Remove(key)
}

func (c *Parameters) Clone() interface{} {
	return NewParametersFromValue(c.GetAsSingleObject())
}

func NewParametersFromValue(value interface{}) *Parameters {
	result := NewEmptyParameters()
	result.SetAsSingleObject(value)
	return result
}

func NewParametersFromTuples(tuples ...interface{}) *Parameters {
	return NewParametersFromTuplesArray(tuples)
}

func NewParametersFromTuplesArray(tuples []interface{}) *Parameters {
	result := NewEmptyParameters()
	if len(tuples) == 0 {
		return result
	}

	for index := 0; index < len(tuples); index = index + 2 {
		if index+1 >= len(tuples) {
			break
		}

		name := convert.StringConverter.ToString(tuples[index])
		value := tuples[index+1]

		result.SetAsObject(name, value)
	}

	return result
}

func NewParametersFromMaps(maps ...map[string]interface{}) *Parameters {
	result := NewEmptyParameters()
	if len(maps) > 0 {
		for index := 0; index < len(maps); index++ {
			result.Append(maps[index])
		}
	}
	return result
}
