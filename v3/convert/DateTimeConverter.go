package convert

import (
	"time"
)

// Converts arbitrary values into Date values using extended conversion rules:
// - Strings: converted using ISO time format
// - Numbers: converted using milliseconds since unix epoch
//
// Example:
//
//  value1 := convert.DateTimeConverter.ToNullableDateTime("ABC")
//  value2 := convert.DateTimeConverter.ToNullableDateTime("2019-01-01T11:30:00.0Z")
//  value3 := convert.DateTimeConverter.ToNullableDateTime(123)
//  fmt.Println(value1) // <nil>
//  fmt.Println(value2) // 2019-01-01 11:30:00 +0000 UTC
//  fmt.Println(value3) // 1970-01-01 02:02:03 +0200 EET
type TDateTimeConverter struct{}

var DateTimeConverter *TDateTimeConverter = &TDateTimeConverter{}

// Converts value into Date or returns null when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: Date value or null when conversion is not supported.
func (c *TDateTimeConverter) ToNullableDateTime(value interface{}) *time.Time {
	return ToNullableDateTime(value)
}

// Converts value into Date or returns current when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: Date value or current when conversion is not supported.
func (c *TDateTimeConverter) ToDateTime(value interface{}) time.Time {
	return ToDateTime(value)
}

// Converts value into Date or returns default when conversion is not possible.
// Parameters: 
// "value" - the value to convert.
// "defaultValue" - the default value.
// Returns: Date value or default when conversion is not supported.
func (c *TDateTimeConverter) ToDateTimeWithDefault(value interface{}, defaultValue time.Time) time.Time {
	return ToDateTimeWithDefault(value, defaultValue)
}

// Converts value into Date or returns null when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: Date value or null when conversion is not supported.
func ToNullableDateTime(value interface{}) *time.Time {
	if value == nil {
		return nil
	}

	var r time.Time

	switch value.(type) {
	case int8:
		r = time.Unix((int64)(value.(int8)), 0)
	case uint8:
		r = time.Unix((int64)(value.(uint8)), 0)
	case int:
		r = time.Unix((int64)(value.(int)), 0)
	case int16:
		r = time.Unix((int64)(value.(int16)), 0)
	case uint16:
		r = time.Unix((int64)(value.(uint16)), 0)
	case int32:
		r = time.Unix((int64)(value.(int32)), 0)
	case uint32:
		r = time.Unix((int64)(value.(uint32)), 0)
	case int64:
		r = time.Unix((int64)(value.(int64)), 0)
	case uint64:
		r = time.Unix((int64)(value.(uint64)), 0)
	case float32:
		r = time.Unix((int64)(value.(float32)), 0)
	case float64:
		r = time.Unix((int64)(value.(float64)), 0)

	case time.Time:
		r = value.(time.Time)

	case string:
		v := value.(string)
		var ok error
		r, ok = time.Parse(time.RFC3339, v)
		if ok != nil {
			r, ok = time.Parse(time.RFC3339Nano, v)
		}
		if ok != nil {
			return nil
		}

	default:
		return nil
	}

	return &r
}

// Converts value into Date or returns current when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: Date value or current when conversion is not supported.
func ToDateTime(value interface{}) time.Time {
	return ToDateTimeWithDefault(value, time.Time{})
}

// Converts value into Date or returns default when conversion is not possible.
// Parameters:
// "value" - the value to convert.
// "defaultValue" - the default value.
// Returns: Date value or default when conversion is not supported.
func ToDateTimeWithDefault(value interface{}, defaultValue time.Time) time.Time {
	r := ToNullableDateTime(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
