package convert

import (
	"reflect"
	"strings"
)

// Converts arbitrary values into array objects.
//
// Example:
//
//  value1 := convert.ArrayConverter.ToArray([...]int{1, 2})
//  value2 := convert.ArrayConverter.ToArray(1)
//  value3 := convert.ArrayConverter.ListToArray("1,2,3")
//  fmt.Println(value1) // [1 2]
//  fmt.Println(value2) // [1]
//  fmt.Println(value3) // [1 2 3]
type TArrayConverter struct{}

var ArrayConverter *TArrayConverter = &TArrayConverter{}

// Converts value into array object. Single values are converted into arrays with a single element.
// Parameters: "value" - the value to convert.
// Returns: array object or null when value is null.
func (c *TArrayConverter) ToNullableArray(value interface{}) *[]interface{} {
	return ToNullableArray(value)
}

// Converts value into array object with empty array as default. Single values
// are converted into arrays with single element.
// Parameters: "value" - the value to convert.
// Returns: array object or empty array when value is null.
func (c *TArrayConverter) ToArray(value interface{}) []interface{} {
	return ToArray(value)
}

// Converts value into array object with empty array as default. Single values
// are converted into arrays with single element.
// Parameters:
// "value" - the value to convert.
// "defaultValue" - default array object.
// Returns: array object or empty array when value is null.
func (c *TArrayConverter) ToArrayWithDefault(value interface{}, defaultValue []interface{}) []interface{} {
	return ToArrayWithDefault(value, defaultValue)
}

// Converts value into array object with empty array as default.
// Strings with comma-delimited values are split into array of strings.
// Parameters: "value" - the list to convert.
// Returns: array object or empty array when value is null
func (c *TArrayConverter) ListToArray(value interface{}) []interface{} {
	return ListToArray(value)
}

// Converts value into array object. Single values are converted into arrays with a single element.
// Parameters: "value" - the value to convert.
// Returns: array object or null when value is null.
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

// Converts value into array object with empty array as default. Single values
// are converted into arrays with single element.
// Parameters: "value" - the value to convert.
// Returns: array object or empty array when value is null.
func ToArray(value interface{}) []interface{} {
	return ToArrayWithDefault(value, []interface{}{})
}

// Converts value into array object with empty array as default. Single values
// are converted into arrays with single element.
// Parameters:
// "value" - the value to convert.
// "defaultValue" - default array object.
// Returns: array object or empty array when value is null.
func ToArrayWithDefault(value interface{}, defaultValue []interface{}) []interface{} {
	if m := ToNullableArray(value); m != nil {
		return *m
	}
	return []interface{}{}
}

// Converts value into array object with empty array as default.
// Strings with comma-delimited values are split into array of strings.
// Parameters: "value" - the list to convert.
// Returns: array object or empty array when value is null
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