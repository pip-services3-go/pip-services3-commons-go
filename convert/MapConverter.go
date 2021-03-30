package convert

import "reflect"

// Converts arbitrary values into map objects using extended conversion rules:
// - Objects: property names as keys, property values as values
// - Arrays: element indexes as keys, elements as values
//
// Example:
//
//  value1 := convert.MapConverter.ToNullableMap("ABC")
//  value2 := convert.MapConverter.ToNullableMap(map[string]int{"key": 123})
//  value3 := convert.MapConverter.ToNullableMap([...]int{1, 2, 3})
//  fmt.Println(value1)  // <nil>
//  fmt.Println(*value2) // map[key:123]
//  fmt.Println(*value3) // map[0:1 1:2 2:3]
type TMapConverter struct{}

var MapConverter *TMapConverter = &TMapConverter{}

// Converts value into map object or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: map object or null when conversion is not supported.
func (c *TMapConverter) ToNullableMap(value interface{}) *map[string]interface{} {
	return ToNullableMap(value)
}

// Converts value into map object or returns empty map when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: map object or empty map when conversion is not supported.
func (c *TMapConverter) ToMap(value interface{}) map[string]interface{} {
	return ToMap(value)
}

// Converts value into map object or returns default map when conversion is not possible.
// Parameters:
//  "value" - the value to convert
//  "defaultValue" - the default value.
// Returns: map object or default map when conversion is not supported.
func (c *TMapConverter) ToMapWithDefault(value interface{}, defaultValue map[string]interface{}) map[string]interface{} {
	return ToMapWithDefault(value, defaultValue)
}

// Converts value into map object or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: map object or null when conversion is not supported.
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

// Converts value into map object or returns empty map when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: map object or empty map when conversion is not supported.
func ToMap(value interface{}) map[string]interface{} {
	return ToMapWithDefault(value, map[string]interface{}{})
}

// Converts value into map object or returns default map when conversion is not possible.
// Parameters:
//  "value" - the value to convert
//  "defaultValue" - the default value.
// Returns: map object or default map when conversion is not supported.
func ToMapWithDefault(value interface{}, defaultValue map[string]interface{}) map[string]interface{} {
	if m := ToNullableMap(value); m != nil {
		return *m
	}
	return map[string]interface{}{}
}
