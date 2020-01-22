package data

import (
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
)

// Cross-language implementation of dynamic object what can hold value of any type.
// The stored value can be converted to different types using variety of accessor methods.
//
// Example:
//
// 	value := data.NewAnyValue("123.456")
//
// 	value.GetAsInteger() // Result: 123
// 	value.GetAsString()  // Result: "123.456"
// 	value.GetAsFloat()   // Result: 123.456
//
type AnyValue struct {
	value interface{}
}

// Creates a new empty instance of the object
// Returns: new empty object
func NewEmptyAnyValue() *AnyValue {
	return &AnyValue{value: nil}
}

// Creates a new instance of the object and assigns its value.
// Parameters: "value" - value to initialize this object.
// Returns: new object.
func NewAnyValue(value interface{}) *AnyValue {
	v, ok := value.(*AnyValue)
	if ok {
		return v
	} else {
		return &AnyValue{value: value}
	}
}

// Gets the value stored in this object without any conversions.
// Returns: the object value.
func (c *AnyValue) InnerValue() interface{} {
	return c.value
}

// Gets the value stored in this object without any conversions.
// Returns: the object value.
func (c *AnyValue) Value() interface{} {
	return c.value
}

// Gets type code for the value stored in this object.
// Returns: type code of the object value.
func (c *AnyValue) TypeCode() convert.TypeCode {
	return convert.TypeConverter.ToTypeCode(c.value)
}

// Gets the value stored in this object without any conversions.
// Returns: the object value.
func (c *AnyValue) GetAsObject() interface{} {
	return c.value
}

// Sets a new value for this object.
// Parameters: "value" - the new object value.
func (c *AnyValue) SetAsObject(value interface{}) {
	c.value = value
}

// Converts object value into a string or returns null if conversion is not possible.
// Returns: string value or null if conversion is not supported.
func (c *AnyValue) GetAsNullableString() *string {
	return convert.StringConverter.ToNullableString(c.value)
}

// Converts object value into a string or returns "" if conversion is not possible.
// Returns: string value or "" if conversion is not supported.
func (c *AnyValue) GetAsString() string {
	return c.GetAsStringWithDefault("")
}

// Converts object value into a string or returns default value if conversion is not possible.
// Parameters: "defaultValue" - the default value.
// Returns: string value or default if conversion is not supported.
func (c *AnyValue) GetAsStringWithDefault(defaultValue string) string {
	return convert.StringConverter.ToStringWithDefault(c.value, defaultValue)
}

// Converts object value into a boolean or returns null if conversion is not possible.
// Returns: boolean value or null if conversion is not supported.
func (c *AnyValue) GetAsNullableBoolean() *bool {
	return convert.BooleanConverter.ToNullableBoolean(c.value)
}

// Converts object value into a boolean or returns false if conversion is not possible.
// Returns: string value or false if conversion is not supported.
func (c *AnyValue) GetAsBoolean() bool {
	return c.GetAsBooleanWithDefault(false)
}

// Converts object value into a boolean or returns default value if conversion is not possible.
// Parameters: "defaultValue" - the default value.
// Returns: boolean value or default if conversion is not supported.
func (c *AnyValue) GetAsBooleanWithDefault(defaultValue bool) bool {
	return convert.BooleanConverter.ToBooleanWithDefault(c.value, defaultValue)
}

// Converts object value into an integer or returns null if conversion is not possible.
// Returns: integer value or null if conversion is not supported.
func (c *AnyValue) GetAsNullableInteger() *int {
	return convert.IntegerConverter.ToNullableInteger(c.value)
}

// Converts object value into an integer or returns 0 if conversion is not possible.
// Returns: integer value or 0 if conversion is not supported.
func (c *AnyValue) GetAsInteger() int {
	return c.GetAsIntegerWithDefault(0)
}

// Converts object value into a integer or returns default value if conversion is not possible.
// Parameters: "defaultValue" - the default value
// Returns: integer value or default if conversion is not supported.
func (c *AnyValue) GetAsIntegerWithDefault(defaultValue int) int {
	return convert.IntegerConverter.ToIntegerWithDefault(c.value, defaultValue)
}

// Converts object value into a long or returns null if conversion is not possible.
// Returns: long value or null if conversion is not supported.
func (c *AnyValue) GetAsNullableLong() *int64 {
	return convert.LongConverter.ToNullableLong(c.value)
}

// Converts object value into a long or returns 0 if conversion is not possible.
// Returns: string value or 0 if conversion is not supported.
func (c *AnyValue) GetAsLong() int64 {
	return c.GetAsLongWithDefault(0)
}

// Converts object value into a long or returns default value if conversion is not possible.
// Parameters: "defaultValue" - the default value
// Returns: long value or default if conversion is not supported.
func (c *AnyValue) GetAsLongWithDefault(defaultValue int64) int64 {
	return convert.LongConverter.ToLongWithDefault(c.value, defaultValue)
}

// Converts object value into a float or returns null if conversion is not possible.
// Returns: float value or null if conversion is not supported.
func (c *AnyValue) GetAsNullableFloat() *float32 {
	return convert.FloatConverter.ToNullableFloat(c.value)
}

// Converts object value into a float or returns 0 if conversion is not possible.
// Returns: float value or 0 if conversion is not supported.
func (c *AnyValue) GetAsFloat() float32 {
	return c.GetAsFloatWithDefault(0)
}

// Converts object value into a float or returns default value if conversion is not possible.
// Parameters: "defaultValue" - the default value
// Returns: float value or default if conversion is not supported.
func (c *AnyValue) GetAsFloatWithDefault(defaultValue float32) float32 {
	return convert.FloatConverter.ToFloatWithDefault(c.value, defaultValue)
}

// Converts object value into a double or returns null if conversion is not possible.
// Returns: double value or null if conversion is not supported.
func (c *AnyValue) GetAsNullableDouble() *float64 {
	return convert.DoubleConverter.ToNullableDouble(c.value)
}

// Converts object value into a double or returns 0 if conversion is not possible.
// Returns: double value or 0 if conversion is not supported.
func (c *AnyValue) GetAsDouble() float64 {
	return c.GetAsDoubleWithDefault(0)
}

// Converts object value into a double or returns default value if conversion is not possible.
// Parameters: "defaultValue" - the default value
// Returns: double value or default if conversion is not supported.
func (c *AnyValue) GetAsDoubleWithDefault(defaultValue float64) float64 {
	return convert.DoubleConverter.ToDoubleWithDefault(c.value, defaultValue)
}

// Converts object value into a Date or returns null if conversion is not possible.
// Returns: DateTime value or null if conversion is not supported.
func (c *AnyValue) GetAsNullableDateTime() *time.Time {
	return convert.DateTimeConverter.ToNullableDateTime(c.value)
}

// Converts object value into a Date or returns current date if conversion is not possible.
// Returns: DateTime value or current date if conversion is not supported.
func (c *AnyValue) GetAsDateTime() time.Time {
	return c.GetAsDateTimeWithDefault(time.Time{})
}

// Converts object value into a Date or returns default value if conversion is not possible.
// Parameters: "defaultValue" - the default value
// Returns: DateTime value or default if conversion is not supported.
func (c *AnyValue) GetAsDateTimeWithDefault(defaultValue time.Time) time.Time {
	return convert.DateTimeConverter.ToDateTimeWithDefault(c.value, defaultValue)
}

// Converts object value into a Duration or returns null if conversion is not possible.
// Returns: Duration value or null if conversion is not supported.
func (c *AnyValue) GetAsNullableDuration() *time.Duration {
	return convert.DurationConverter.ToNullableDuration(c.value)
}

// Converts object value into a Duration or returns current date if conversion is not possible.
// Returns: Duration value or current date if conversion is not supported.
func (c *AnyValue) GetAsDuration() time.Duration {
	return c.GetAsDurationWithDefault(0 * time.Millisecond)
}

// Converts object value into a Duration or returns default value if conversion is not possible.
// Parameters: "defaultValue" - the default value
// Returns: Duration value or default if conversion is not supported.
func (c *AnyValue) GetAsDurationWithDefault(defaultValue time.Duration) time.Duration {
	return convert.DurationConverter.ToDurationWithDefault(c.value, defaultValue)
}

// Converts object value into a value defined by specied typecode. If conversion is not possible it returns null.
// Parameters: "typ" - the TypeCode that defined the type of the result.
// Returns: value defined by the typecode or null if conversion is not supported.
func (c *AnyValue) GetAsNullableType(typ convert.TypeCode) interface{} {
	return convert.TypeConverter.ToNullableType(typ, c.value)
}

// Converts object value into a value defined by specied typecode. If conversion
// is not possible it returns default value for the specified type.
// Parameters: "typ" - the TypeCode that defined the type of the result.
// Returns: value defined by the typecode or type default value if conversion is not supported.
func (c *AnyValue) GetAsType(typ convert.TypeCode) interface{} {
	return c.GetAsTypeWithDefault(typ, nil)
}

// Converts object value into a value defined by specied typecode. If conversion
// is not possible it returns default value.
// Parameters:
// "typ" - the TypeCode that defined the type of the result;
// "defaultValue" - the default value.
// Returns: value defined by the typecode or type default value if conversion is not supported.
func (c *AnyValue) GetAsTypeWithDefault(typ convert.TypeCode, defaultValue interface{}) interface{} {
	return convert.TypeConverter.ToTypeWithDefault(typ, c.value, defaultValue)
}

// Converts object value into an AnyArray or returns empty AnyArray if conversion is not possible.
// Returns: AnyArray value or empty AnyArray if conversion is not supported.
func (c *AnyValue) GetAsArray() *AnyValueArray {
	return NewAnyValueArrayFromValue(c.value)
}

// Converts object value into AnyMap or returns empty AnyMap if conversion is not possible.
// Returns: AnyMap value or empty AnyMap if conversion is not supported.
func (c *AnyValue) GetAsMap() *AnyValueMap {
	return NewAnyValueMapFromValue(c.value)
}

// Compares this object value to specified specified value. When direct
// comparison gives negative results it tries to compare values as strings.
// Parameters: "obj" - the value to be compared with.
// Returns: true when objects are equal and false otherwise.
func (c *AnyValue) Equals(obj interface{}) bool {
	if obj == nil && c.value == nil {
		return true
	}
	if obj == nil || c.value == nil {
		return false
	}

	v, ok := obj.(*AnyValue)
	if ok {
		obj = v.value
	}

	strThisValue := convert.StringConverter.ToNullableString(c.value)
	strValue := convert.StringConverter.ToNullableString(obj)

	if strThisValue == nil && strValue == nil {
		return true
	}
	if strThisValue == nil || strValue == nil {
		return false
	}
	return (*strThisValue) == (*strValue)
}

// Compares this object value to specified specified value. When direct
// comparison gives negative results it converts values to type specified by
// type code and compare them again.
// Parameters:
// "typ" - the TypeCode that defined the type of the result.
// "obj" - the value to be compared with.
// Returns: true when objects are equal and false otherwise.
func (c *AnyValue) EqualsAsType(typ convert.TypeCode, obj interface{}) bool {
	if obj == nil && c.value == nil {
		return true
	}
	if obj == nil || c.value == nil {
		return false
	}

	v, ok := obj.(*AnyValue)
	if ok {
		obj = v.value
	}

	typedThisValue := convert.TypeConverter.ToType(typ, c.value)
	typedValue := convert.TypeConverter.ToType(typ, obj)

	return typedThisValue == typedValue
}

// Creates a binary clone of this object.
// Returns: a clone of this object.
func (c *AnyValue) Clone() interface{} {
	return NewAnyValue(c.value)
}

// Gets a string representation of the object.
// Returns: a string representation of the object.
func (c *AnyValue) String() string {
	return convert.StringConverter.ToString(c.value)
}
