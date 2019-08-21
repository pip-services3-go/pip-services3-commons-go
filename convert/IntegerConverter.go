package convert

// Converts arbitrary values into integer using extended conversion rules:
// - Strings are converted to integer values
// - DateTime: total number of milliseconds since unix epoсh
// - Boolean: 1 for true and 0 for false
//
// Example:
//
//  value1 := convert.IntegerConverter.ToNullableInteger("ABC")
//  value2 := convert.IntegerConverter.ToNullableInteger("123.456")
//  value3 := convert.IntegerConverter.ToNullableInteger(true)
//  value4 := convert.IntegerConverter.ToNullableInteger(time.Now())
//  fmt.Println(value1)  // <nil>
//  fmt.Println(*value2) // 123
//  fmt.Println(*value3) // 1
//  fmt.Println(*value4) // current milliseconds (e.g. 1566333428)
type TIntegerConverter struct{}

var IntegerConverter *TIntegerConverter = &TIntegerConverter{}

// Converts value into integer or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value or null when conversion is not supported.
func (c *TIntegerConverter) ToNullableInteger(value interface{}) *int {
	return ToNullableInteger(value)
}

// Converts value into integer or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value or 0 when conversion is not supported.
func (c *TIntegerConverter) ToInteger(value interface{}) int {
	return ToInteger(value)
}

// Converts value into integer or returns default when conversion is not possible.
// Parameters: 
// "value" - the value to convert.
// "defaultValue" - the default value.
// Returns: integer value or default when conversion is not supported.
func (c *TIntegerConverter) ToIntegerWithDefault(value interface{}, defaultValue int) int {
	return ToIntegerWithDefault(value, defaultValue)
}

// Converts value into integer or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value or null when conversion is not supported.
func ToNullableInteger(value interface{}) *int {
	v := ToNullableLong(value)
	if v == nil {
		return nil
	}
	r := int(*v)
	return &r
}

// Converts value into integer or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value or 0 when conversion is not supported.
func ToInteger(value interface{}) int {
	return ToIntegerWithDefault(value, 0)
}

// Converts value into integer or returns default when conversion is not possible.
// Parameters: 
// "value" - the value to convert.
// "defaultValue" - the default value.
// Returns: integer value or default when conversion is not supported.
func ToIntegerWithDefault(value interface{}, defaultValue int) int {
	r := ToNullableInteger(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
