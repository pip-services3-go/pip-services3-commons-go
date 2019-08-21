package convert

import (
	"reflect"
	"strconv"
)

// Converts arbitrary values into interface
// Parameters: "value" - the reflect.Value to convert.
// Returns: the interface of specific type.
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
}

// Converts array values into array interface
// Parameters: "value" - the array of reflect.Value to convert.
// Returns: the interface array of specific type.
func arrayToArray(value reflect.Value) []interface{} {
	r := []interface{}{}

	for i := 0; i < value.Len(); i++ {
		v := value.Index(i)
		r = append(r, valueToInterface(v))
	}

	return r
}

// Converts array values into map interface
// Parameters: "value" - the array of reflect.Value to convert.
// Returns: the map with values of specific types.
func arrayToMap(value reflect.Value) map[string]interface{} {
	r := map[string]interface{}{}

	for i := 0; i < value.Len(); i++ {
		k := strconv.FormatInt(int64(i), 10)
		v := valueToInterface(value.Index(i))
		r[k] = v
	}

	return r
}

// Converts map values into array interface
// Parameters: "value" - the map to convert.
// Returns: the interface array of specific type.
func mapToArray(value reflect.Value) []interface{} {
	r := []interface{}{}

	for _, key := range value.MapKeys() {
		v := valueToInterface(value.MapIndex(key))
		r = append(r, v)
	}

	return r
}

// Converts map values into map interface
// Parameters: "value" - the map to convert.
// Returns: the map with values of specific types.
func mapToMap(value reflect.Value) map[string]interface{} {
	r := map[string]interface{}{}

	for _, key := range value.MapKeys() {
		k := ToString(valueToInterface(key))
		v := valueToInterface(value.MapIndex(key))
		r[k] = v
	}

	return r
}

// Converts struct values into map interface
// Parameters: "value" - the struct to convert.
// Returns: the map with values of specific types.
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
