package convert

import "reflect"

type TMapConverter struct{}

var MapConverter *TMapConverter = &TMapConverter{}

func (c *TMapConverter) ToNullableMap(value interface{}) *map[string]interface{} {
	return ToNullableMap(value)
}

func (c *TMapConverter) ToMap(value interface{}) map[string]interface{} {
	return ToMap(value)
}

func (c *TMapConverter) ToMapWithDefault(value interface{}, defaultValue map[string]interface{}) map[string]interface{} {
	return ToMapWithDefault(value, defaultValue)
}

func ToNullableMap(value interface{}) *map[string]interface{} {
	if value == nil {
		return nil
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {

	case reflect.Map:
		r := mapToMap(v)
		return &r

	case reflect.Array, reflect.Slice:
		r := arrayToMap(v)
		return &r

	case reflect.Struct:
		r := structToMap(v)
		return &r

	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		value = valueToInterface(v.Elem())
		return ToNullableMap(value)
	}

	return nil
}

func ToMap(value interface{}) map[string]interface{} {
	return ToMapWithDefault(value, map[string]interface{}{})
}

func ToMapWithDefault(value interface{}, defaultValue map[string]interface{}) map[string]interface{} {
	if m := ToNullableMap(value); m != nil {
		return *m
	}
	return map[string]interface{}{}
}
