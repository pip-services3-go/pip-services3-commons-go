package run

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/pip-services3-go/pip-services3-commons-go/reflect"
)

type Parameters struct {
	data.AnyValueMap
}

func NewEmptyParameters() *Parameters {
	c := &Parameters{}
	c.AnyValueMap = *data.InheritAnyValueMap(c)
	return c
}

func NewParameters(values map[string]interface{}) *Parameters {
	c := &Parameters{}
	c.AnyValueMap = *data.InheritAnyValueMap(c)
	c.Append(values)
	return c
}

func (c *Parameters) Get(key string) interface{} {
	if key == "" {
		return nil
	} else if strings.Index(key, ".") > 0 {
		return reflect.RecursiveObjectReader.GetProperty(c.InnerValue(), key)
	} else {
		return c.AnyValueMap.Get(key)
	}
}

func (c *Parameters) Put(key string, value interface{}) {
	if key == "" {
		// Do nothing...
	} else if strings.Index(key, ".") > 0 {
		reflect.RecursiveObjectWriter.SetProperty(c.InnerValue(), key, value)
	} else {
		c.AnyValueMap.Put(key, value)
	}
}

func (c *Parameters) Remove(key string) {
	// Todo: Make this method recursive
	c.AnyValueMap.Remove(key)
}

func (c *Parameters) Contains(key string) bool {
	return reflect.RecursiveObjectReader.HasProperty(c.InnerValue(), key)
}

func (c *Parameters) GetAsNullableParameters(key string) *Parameters {
	value := c.GetAsNullableMap(key)
	if value != nil {
		return NewParametersFromValue(value)
	}
	return nil
}

func (c *Parameters) GetAsParameters(key string) *Parameters {
	value := c.GetAsMap(key)
	return NewParametersFromValue(value)
}

func (c *Parameters) GetAsParametersWithDefault(key string, defaultValue *Parameters) *Parameters {
	result := c.GetAsNullableParameters(key)
	if result == nil {
		return defaultValue
	}
	return result
}

func (c *Parameters) Clone() interface{} {
	return NewParametersFromValue(c.GetAsSingleObject())
}

func (c *Parameters) Override(parameters *Parameters, recursive bool) *Parameters {
	if parameters == nil {
		return c
	}

	result := NewEmptyParameters()
	if recursive {
		reflect.RecursiveObjectWriter.CopyProperties(result.InnerValue(), c.InnerValue())
		reflect.RecursiveObjectWriter.CopyProperties(result.InnerValue(), parameters.InnerValue())
	} else {
		reflect.ObjectWriter.SetProperties(result.InnerValue(), c.InnerValue().(map[string]interface{}))
		reflect.ObjectWriter.SetProperties(result.InnerValue(), parameters.InnerValue().(map[string]interface{}))
	}
	return result
}

func (c *Parameters) SetDefaults(defaultParameters *Parameters, recursive bool) *Parameters {
	if defaultParameters == nil {
		return c
	}

	result := NewEmptyParameters()
	if recursive {
		reflect.RecursiveObjectWriter.CopyProperties(result.InnerValue(), defaultParameters.InnerValue())
		reflect.RecursiveObjectWriter.CopyProperties(result.InnerValue(), c.InnerValue())
	} else {
		reflect.ObjectWriter.SetProperties(result.InnerValue(), defaultParameters.InnerValue().(map[string]interface{}))
		reflect.ObjectWriter.SetProperties(result.InnerValue(), c.InnerValue().(map[string]interface{}))
	}
	return result
}

func (c *Parameters) AssignTo(value interface{}) {
	if value == nil {
		return
	}
	reflect.RecursiveObjectWriter.CopyProperties(value, c.InnerValue())
}

func (c *Parameters) Pick(paths ...string) *Parameters {
	result := NewEmptyParameters()
	for _, path := range paths {
		if c.Contains(path) {
			result.Put(path, c.Get(path))
		}
	}
	return result
}

func (c *Parameters) Omit(paths ...string) *Parameters {
	result := NewParametersFromValue(c.InnerValue())
	for _, path := range paths {
		result.Remove(path)
	}
	return result
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

func NewParametersFromConfig(config *config.ConfigParams) *Parameters {
	result := NewEmptyParameters()
	values := config.InnerValue().(map[string]string)
	for key, value := range values {
		result.Put(key, value)
	}
	return result
}
