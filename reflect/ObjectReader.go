package reflect

import (
	refl "reflect"
	"strconv"
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
)

type TObjectReader struct{}

var ObjectReader = &TObjectReader{}

func (c *TObjectReader) GetValue(obj interface{}) interface{} {
	wrap, ok := obj.(IValueWrapper)
	if ok {
		obj = wrap.InnerValue()
	}

	return obj
}

func (c *TObjectReader) HasProperty(obj interface{}, name string) bool {
	if obj == nil || name == "" {
		return false
	}

	obj = c.GetValue(obj)
	val := refl.ValueOf(obj)

	if val.Kind() == refl.Map {
		name = strings.ToLower(name)
		for _, v := range val.MapKeys() {
			key := convert.StringConverter.ToString(v.Interface())
			key = strings.ToLower(key)
			if name == key {
				return true
			}
		}
		return false
	}

	if val.Kind() == refl.Slice || val.Kind() == refl.Array {
		index := convert.IntegerConverter.ToIntegerWithDefault(name, -1)
		return index >= 0 && index < val.Len()
	}

	return PropertyReflector.HasProperty(obj, name)
}

func (c *TObjectReader) GetProperty(obj interface{}, name string) interface{} {
	if obj == nil || name == "" {
		return nil
	}

	obj = c.GetValue(obj)
	val := refl.ValueOf(obj)

	if val.Kind() == refl.Map {
		name = strings.ToLower(name)
		for _, v := range val.MapKeys() {
			key := convert.StringConverter.ToString(v.Interface())
			key = strings.ToLower(key)
			if name == key {
				return val.MapIndex(v).Interface()
			}
		}
		return nil
	}

	if val.Kind() == refl.Slice || val.Kind() == refl.Array {
		index := convert.IntegerConverter.ToIntegerWithDefault(name, -1)
		if index >= 0 && index < val.Len() {
			return val.Index(index).Interface()
		}
		return nil
	}

	return PropertyReflector.GetProperty(obj, name)
}

func (c *TObjectReader) GetPropertyNames(obj interface{}) []string {
	if obj == nil {
		return nil
	}

	obj = c.GetValue(obj)
	val := refl.ValueOf(obj)
	properties := []string{}

	if val.Kind() == refl.Map {
		for _, v := range val.MapKeys() {
			key := convert.StringConverter.ToString(v.Interface())
			properties = append(properties, key)
		}
		return properties
	}

	if val.Kind() == refl.Slice || val.Kind() == refl.Array {
		for index := 0; index < val.Len(); index++ {
			properties = append(properties, strconv.Itoa(index))
		}
		return properties
	}

	return PropertyReflector.GetPropertyNames(obj)
}

func (c *TObjectReader) GetProperties(obj interface{}) map[string]interface{} {
	if obj == nil {
		return nil
	}

	obj = c.GetValue(obj)
	val := refl.ValueOf(obj)
	values := map[string]interface{}{}

	if val.Kind() == refl.Map {
		for _, v := range val.MapKeys() {
			key := convert.StringConverter.ToString(v.Interface())
			values[key] = val.MapIndex(v).Interface()
		}
		return values
	}

	if val.Kind() == refl.Slice || val.Kind() == refl.Array {
		for index := 0; index < val.Len(); index++ {
			values[strconv.Itoa(index)] = val.Index(index).Interface()
		}
		return values
	}

	return PropertyReflector.GetProperties(obj)
}
