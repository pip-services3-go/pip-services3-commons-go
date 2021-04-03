package convert

import (
	"encoding/json"
)

// Converts value from JSON string
// Parameters: "value" - the JSON string to convert.
// Returns: converted object value or null when value is null.
func FromJson(value string) (interface{}, error) {
	if value == "" {
		return nil, nil
	}

	var m interface{}
	return FromJsonAs(m, value)
}

// Converts value from JSON string to an object with specified type
// Parameters:
//   result - a references to the object that will receive a converted value.
//   value - the JSON string to convert.
// Returns: converted object value or null when value is null.
func FromJsonAs(result interface{}, value string) (interface{}, error) {
	if value == "" {
		return nil, nil
	}

	if err := json.Unmarshal([]byte(value), &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Converts value into JSON string.
// Parameters: "value" - the value to convert.
// Returns: JSON string or null when value is null.
func ToJson(value interface{}) (string, error) {
	if value == nil {
		return "", nil
	}

	b, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}

// Converts arbitrary values from and to JSON (JavaScript Object Notation) strings.
//
// Example:
//
//  value1, _ := convert.FromJson("{\"key\":123}")
//  value2 := convert.JsonConverter.ToMap("{\"key\":123}")
//  value3, _ := convert.ToJson(map[string]int{"key": 123})
//  fmt.Println(value1) // map[key:123]
//  fmt.Println(value2) // map[key:123]
//  fmt.Println(value3) // {"key":123}
type TJsonConverter struct{}

var JsonConverter *TJsonConverter = &TJsonConverter{}

// Converts JSON string into map object or returns null when conversion is not possible.
// Parameters: "value" - the JSON string to convert.
// Returns: Map object value or null when conversion is not supported.
func (c *TJsonConverter) ToNullableMap(value string) *map[string]interface{} {
	return JsonToNullableMap(value)
}

// Converts JSON string into map object or returns empty map when conversion is not possible.
// Parameters: "value" - the JSON string to convert.
// Returns: Map object value or empty map when conversion is not supported.
func (c *TJsonConverter) ToMap(value string) map[string]interface{} {
	return JsonToMap(value)
}

// Converts JSON string into map object or returns default map when conversion is not possible.
// Parameters:
//  "value" - the JSON string to convert.
//  "defaultValue" - the default value.
// Returns: Map object value or default map when conversion is not supported.
func (c *TJsonConverter) ToMapWithDefault(value string, defaultValue map[string]interface{}) map[string]interface{} {
	return JsonToMapWithDefault(value, defaultValue)
}

// Converts JSON string into an object.
// Parameters: "value" - the JSON string to convert.
// Returns: Map object value or empty map when conversion is not supported.
func (c *TJsonConverter) ToObject(value string) interface{} {
	result, _ := FromJson(value)
	return result
}

// Converts JSON string into an object with specified type.
// Parameters:
//   result - a references to an object that will receive the result
//   value - the JSON string to convert.
// Returns: Map object value or empty map when conversion is not supported.
func (c *TJsonConverter) ToObjectAs(result interface{}, value string) (interface{}, error) {
	return FromJsonAs(result, value)
}

// Converts JSON string into map object or returns null when conversion is not possible.
// Parameters: "value" - the JSON string to convert.
// Returns: Map object value or null when conversion is not supported.
func JsonToNullableMap(value string) *map[string]interface{} {
	v, _ := FromJson(value)
	if v == nil {
		return nil
	}
	return ToNullableMap(v)
}

// Converts JSON string into map object or returns empty map when conversion is not possible.
// Parameters: "value" - the JSON string to convert.
// Returns: Map object value or empty map when conversion is not supported.
func JsonToMap(value string) map[string]interface{} {
	return JsonToMapWithDefault(value, map[string]interface{}{})
}

// Converts JSON string into map object or returns default map when conversion is not possible.
// Parameters:
//  "value" - the JSON string to convert.
//  "defaultValue" - the default value.
// Returns: Map object value or default map when conversion is not supported.
func JsonToMapWithDefault(value string, defaultValue map[string]interface{}) map[string]interface{} {
	if m := JsonToNullableMap(value); m != nil {
		return *m
	}
	return map[string]interface{}{}
}
