package convert

import (
	"strconv"
	"time"
)

// Converts arbitrary values into double using extended conversion rules:
// - Strings are converted to double values
// - DateTime: total number of milliseconds since unix epoсh
// - Boolean: 1 for true and 0 for false
//
// Example:
//
//  value1 := convert.DoubleConverter.ToNullableDouble("ABC")
//  value2 := convert.DoubleConverter.ToNullableDouble("123.456")
//  value3 := convert.DoubleConverter.ToNullableDouble(true)
//  value4 := convert.DoubleConverter.ToNullableDouble(time.Now())
//  fmt.Println(value1)  // <nil>
//  fmt.Println(*value2) // 123.456
//  fmt.Println(*value3) // 1
//  fmt.Println(*value4) // current milliseconds (e.g. 1.566333114e+09)
type TDoubleConverter struct{}

var DoubleConverter *TDoubleConverter = &TDoubleConverter{}

// Converts value into doubles or returns null when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: double value or null when conversion is not supported.
func (c *TDoubleConverter) ToNullableDouble(value interface{}) *float64 {
	return ToNullableDouble(value)
}

// Converts value into doubles or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: double value or 0 when conversion is not supported.
func (c *TDoubleConverter) ToDouble(value interface{}) float64 {
	return ToDouble(value)
}

// Converts value into doubles or returns default when conversion is not possible.
// Parameters: 
//  "value" - the value to convert.
//  "defaultValue" - the default value
// Returns: double value or default when conversion is not supported.
func (c *TDoubleConverter) ToDoubleWithDefault(value interface{}, defaultValue float64) float64 {
	return ToDoubleWithDefault(value, defaultValue)
}

// Converts value into doubles or returns null when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: double value or null when conversion is not supported.
func ToNullableDouble(value interface{}) *float64 {
	if value == nil {
		return nil
	}

	var r float64 = 0

	switch value.(type) {
	case int8:
		r = (float64)(value.(int8))
	case uint8:
		r = (float64)(value.(uint8))
	case int:
		r = (float64)(value.(int))
	case int16:
		r = (float64)(value.(int16))
	case uint16:
		r = (float64)(value.(uint16))
	case int32:
		r = (float64)(value.(int32))
	case uint32:
		r = (float64)(value.(uint32))
	case int64:
		r = (float64)(value.(int64))
	case uint64:
		r = (float64)(value.(uint64))
	case float32:
		r = (float64)(value.(float32))
	case float64:
		r = (float64)(value.(float64))

	case bool:
		v := value.(bool)
		if v == true {
			r = 1.0
		}

	case time.Time:
		r = float64(value.(time.Time).Unix())

	case time.Duration:
		r = float64(value.(time.Duration).Nanoseconds() / 1000000)

	case string:
		var ok error
		r, ok = strconv.ParseFloat(value.(string), 0)
		if ok != nil {
			return nil
		}

	default:
		return nil
	}

	return &r
}

// Converts value into doubles or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: double value or 0 when conversion is not supported.
func ToDouble(value interface{}) float64 {
	return ToDoubleWithDefault(value, 0)
}

// Converts value into doubles or returns default when conversion is not possible.
// Parameters: 
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: double value or default when conversion is not supported.
func ToDoubleWithDefault(value interface{}, defaultValue float64) float64 {
	r := ToNullableDouble(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
