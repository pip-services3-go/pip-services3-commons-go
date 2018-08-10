package convert

import (
	"reflect"
	"strconv"
)

func valueToInterface(value reflect.Value) interface{} {
	switch value.Kind() {
	case reflect.Invalid:
		return nil
	case reflect.String:
		return value.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int64(value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(value.Uint())
	case reflect.Float32, reflect.Float64:
		return float64(value.Float())
	case reflect.Bool:
		return value.Bool()
	case reflect.Map:
		return mapToMap(value)
	case reflect.Array, reflect.Slice:
		return arrayToArray(value)
	case reflect.Struct:
		return structToMap(value)
	case reflect.Interface, reflect.Ptr:
		if value.IsNil() {
			return nil
		}
		return valueToInterface(value.Elem())
	}

	return value.Interface()

	// return nil

	// if value.IsNil() {
	// 	return nil
	// }
	// return value.Interface()
}

func arrayToArray(value reflect.Value) []interface{} {
	r := []interface{}{}

	for i := 0; i < value.Len(); i++ {
		v := value.Index(i)
		r = append(r, valueToInterface(v))
	}

	return r
}

func arrayToMap(value reflect.Value) map[string]interface{} {
	r := map[string]interface{}{}

	for i := 0; i < value.Len(); i++ {
		k := strconv.FormatInt(int64(i), 10)
		v := valueToInterface(value.Index(i))
		r[k] = v
	}

	return r
}

func mapToArray(value reflect.Value) []interface{} {
	r := []interface{}{}

	for _, key := range value.MapKeys() {
		v := valueToInterface(value.MapIndex(key))
		r = append(r, v)
	}

	return r
}

func mapToMap(value reflect.Value) map[string]interface{} {
	r := map[string]interface{}{}

	for _, key := range value.MapKeys() {
		k := ToString(valueToInterface(key))
		v := valueToInterface(value.MapIndex(key))
		r[k] = v
	}

	return r
}

func structToMap(value reflect.Value) map[string]interface{} {
	t := value.Type()
	r := map[string]interface{}{}

	for i := 0; i < value.NumField(); i++ {
		k := t.Field(i).Name
		v := valueToInterface(value.Field(i))
		r[k] = v
	}

	return r
}
