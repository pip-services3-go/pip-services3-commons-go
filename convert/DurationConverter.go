package convert

import (
	"time"
)

// Converts arbitrary values into time.Duration values.
//
// Example:
//
//  value1 := convert.DurationConverter.ToNullableDuration("123")
//  value2 := convert.DurationConverter.ToNullableDuration(123)
//  value3 := convert.DurationConverter.ToNullableDuration(123 * time.Second)
//  fmt.Println(value1) // 123ms
//  fmt.Println(value2) // 123ms
//  fmt.Println(value3) // 2m3s
type TDurationConverter struct{}

var DurationConverter *TDurationConverter = &TDurationConverter{}

// Converts value into time.Duration or returns null when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: time.Duration value or null when conversion is not supported.
func (c *TDurationConverter) ToNullableDuration(value interface{}) *time.Duration {
	return ToNullableDuration(value)
}

// Converts value into time.Duration or returns current when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: time.Duration value or current when conversion is not supported.
func (c *TDurationConverter) ToDuration(value interface{}) time.Duration {
	return ToDuration(value)
}

// Converts value into time.Duration or returns default when conversion is not possible.
// Parameters: 
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: time.Duration value or default when conversion is not supported.
func (c *TDurationConverter) ToDurationWithDefault(value interface{}, defaultValue time.Duration) time.Duration {
	return ToDurationWithDefault(value, defaultValue)
}

// Converts value into time.Duration or returns null when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: time.Duration value or null when conversion is not supported.
func ToNullableDuration(value interface{}) *time.Duration {
	if value == nil {
		return nil
	}

	var r time.Duration

	switch value.(type) {
	case int8:
		r = (time.Duration)(value.(int8)) * time.Millisecond
	case uint8:
		r = (time.Duration)(value.(uint8)) * time.Millisecond
	case int:
		r = (time.Duration)(value.(int)) * time.Millisecond
	case int16:
		r = (time.Duration)(value.(int16)) * time.Millisecond
	case uint16:
		r = (time.Duration)(value.(uint16)) * time.Millisecond
	case int32:
		r = (time.Duration)(value.(int32)) * time.Millisecond
	case uint32:
		r = (time.Duration)(value.(uint32)) * time.Millisecond
	case int64:
		r = (time.Duration)(value.(int64)) * time.Millisecond
	case uint64:
		r = (time.Duration)(value.(uint64)) * time.Millisecond
	case float32:
		r = (time.Duration)(value.(float32)) * time.Millisecond
	case float64:
		r = (time.Duration)(value.(float64)) * time.Millisecond

	case time.Duration:
		r = value.(time.Duration)

	case string:
		v := value.(string)
		var ok error
		r, ok = time.ParseDuration(v)
		if ok != nil {
			r = (time.Duration)(ToLong(value)) * time.Millisecond
		}

	default:
		return nil
	}

	return &r
}

// Converts value into time.Duration or returns current when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: time.Duration value or current when conversion is not supported.
func ToDuration(value interface{}) time.Duration {
	return ToDurationWithDefault(value, 0*time.Millisecond)
}

// Converts value into time.Duration or returns default when conversion is not possible.
// Parameters: 
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: time.Duration value or default when conversion is not supported.
func ToDurationWithDefault(value interface{}, defaultValue time.Duration) time.Duration {
	r := ToNullableDuration(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
