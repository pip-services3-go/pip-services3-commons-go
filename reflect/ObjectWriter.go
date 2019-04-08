package reflect

import (
	refl "reflect"
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
)

type TObjectWriter struct{}

var ObjectWriter = &TObjectWriter{}

func (c *TObjectWriter) GetValue(obj interface{}) interface{} {
	wrap, ok := obj.(IValueWrapper)
	if ok {
		obj = wrap.InnerValue()
	}

	return obj
}

func (c *TObjectWriter) SetProperty(obj interface{}, name string, value interface{}) {
	if obj == nil || name == "" {
		return
	}

	obj = c.GetValue(obj)
	val := refl.ValueOf(obj)

	if val.Kind() == refl.Map {
		name = strings.ToLower(name)
		for _, v := range val.MapKeys() {
			key := convert.StringConverter.ToString(v.Interface())
			key = strings.ToLower(key)
			if name == key {
				val.SetMapIndex(v, refl.ValueOf(value))
				return
			}
		}
		val.SetMapIndex(refl.ValueOf(name), refl.ValueOf(value))
		return
	}

	if val.Kind() == refl.Slice || val.Kind() == refl.Array {
		index := convert.IntegerConverter.ToIntegerWithDefault(name, -1)

		// Todo: Think how to resize slice

		// Set array element
		if index >= 0 && index < val.Len() {
			v := val.Index(index)
			v.Set(refl.ValueOf(value))
			return
		}
		return
	}

	PropertyReflector.SetProperty(obj, name, value)
}

func (c *TObjectWriter) SetProperties(obj interface{}, values map[string]interface{}) {
	if values == nil || len(values) == 0 {
		return
	}

	for key, value := range values {
		c.SetProperty(obj, key, value)
	}
}
