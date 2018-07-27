package convert

import (
	"reflect"
	"strings"
)

type TArrayConverter struct{}

var ArrayConverter *TArrayConverter = &TArrayConverter{}

func (c *TArrayConverter) ToNullableArray(value interface{}) *[]interface{} {
	return ToNullableArray(value)
}

func (c *TArrayConverter) ToArray(value interface{}) []interface{} {
	return ToArray(value)
}

func (c *TArrayConverter) ToArrayWithDefault(value interface{}, defaultValue []interface{}) []interface{} {
	return ToArrayWithDefault(value, defaultValue)
}

func (c *TArrayConverter) ListToArray(value interface{}) []interface{} {
	return ListToArray(value)
}

func ToNullableArray(value interface{}) *[]interface{} {
	if value == nil {
		return nil
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {

	case reflect.Map:
		r := mapToArray(v)
		return &r

	case reflect.Array, reflect.Slice:
		r := arrayToArray(v)
		return &r

	default:
		value = valueToInterface(v)
		r := []interface{}{value}
		return &r
	}
}

func ToArray(value interface{}) []interface{} {
	return ToArrayWithDefault(value, []interface{}{})
}

func ToArrayWithDefault(value interface{}, defaultValue []interface{}) []interface{} {
	if m := ToNullableArray(value); m != nil {
		return *m
	}
	return []interface{}{}
}

func ListToArray(value interface{}) []interface{} {
	if value == nil {
		return []interface{}{}
	}

	v := reflect.ValueOf(value)

	if v.Kind() == reflect.String {
		value = strings.Split(value.(string), ",")
	}

	return ToArray(value)
}
