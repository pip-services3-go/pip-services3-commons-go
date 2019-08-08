package data

import (
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
)

/*
Cross-language implementation of dynamic object what can hold value of any type.
The stored value can be converted to different types using variety of accessor methods.

### Example ###

	value := data.NewAnyValue("123.456")

	value.GetAsInteger() // Result: 123
	value.GetAsString()  // Result: "123.456"
	value.GetAsFloat()   // Result: 123.456

see: StringConverter
see: TypeConverter
see: BooleanConverter
see: IntegerConverter
see: LongConverter
see: DoubleConverter
see: FloatConverter
see: DateTimeConverter
see: ICloneable
*/
type AnyValue struct {
	value interface{}
}

// Creates a new empty instance of the object.
// returns: new empty object
func NewEmptyAnyValue() *AnyValue {
	return &AnyValue{value: nil}
}

// Creates a new instance of the object and assigns its value.
// param "value": value to initialize this object.
// returns: new object.
func NewAnyValue(value interface{}) *AnyValue {
	v, ok := value.(*AnyValue)
	if ok {
		return v
	} else {
		return &AnyValue{value: value}
	}
}

// Gets the value stored in this object without any conversions.
// returns: the object value.
func (c *AnyValue) InnerValue() interface{} {
	return c.value
}

// Gets the value stored in this object without any conversions.
// returns: the object value.
func (c *AnyValue) Value() interface{} {
	return c.value
}

// Gets type code for the value stored in this object.
// returns: type code of the object value.
// see:	TypeConverter.ToTypeCode
func (c *AnyValue) TypeCode() convert.TypeCode {
	return convert.TypeConverter.ToTypeCode(c.value)
}

// Gets the value stored in this object without any conversions.
// returns: the object value.
func (c *AnyValue) GetAsObject() interface{} {
	return c.value
}

// Sets a new value for this object.
// param "value": the new object value.
func (c *AnyValue) SetAsObject(value interface{}) {
	c.value = value
}

// Converts object value into a string or returns null if conversion is not possible.
// returns: string value or null if conversion is not supported.
// see: StringConverter.ToNullableString
func (c *AnyValue) GetAsNullableString() *string {
	return convert.StringConverter.ToNullableString(c.value)
}

// Converts object value into a string or returns "" if conversion is not possible.
// returns: string value or "" if conversion is not supported.
// see: GetAsStringWithDefault
func (c *AnyValue) GetAsString() string {
	return c.GetAsStringWithDefault("")
}

// Converts object value into a string or returns default value if conversion is not possible.
// param "defaultValue": the default value.
// returns: string value or default if conversion is not supported.
// see: StringConverter.ToStringWithDefault
func (c *AnyValue) GetAsStringWithDefault(defaultValue string) string {
	return convert.StringConverter.ToStringWithDefault(c.value, defaultValue)
}

// Converts object value into a boolean or returns null if conversion is not possible.
// returns: boolean value or null if conversion is not supported.
// see: BooleanConverter.ToNullableBoolean
func (c *AnyValue) GetAsNullableBoolean() *bool {
	return convert.BooleanConverter.ToNullableBoolean(c.value)
}

// Converts object value into a boolean or returns false if conversion is not possible.
// returns: string value or false if conversion is not supported.
// see: GetAsBooleanWithDefault
func (c *AnyValue) GetAsBoolean() bool {
	return c.GetAsBooleanWithDefault(false)
}

// Converts object value into a boolean or returns default value if conversion is not possible.
// param "defaultValue": the default value.
// returns: boolean value or default if conversion is not supported.
// see: BooleanConverter.ToBooleanWithDefault
func (c *AnyValue) GetAsBooleanWithDefault(defaultValue bool) bool {
	return convert.BooleanConverter.ToBooleanWithDefault(c.value, defaultValue)
}

// Converts object value into an integer or returns null if conversion is not possible.
// returns: integer value or null if conversion is not supported.
// see: IntegerConverter.ToNullableInteger
func (c *AnyValue) GetAsNullableInteger() *int {
	return convert.IntegerConverter.ToNullableInteger(c.value)
}

// Converts object value into an integer or returns 0 if conversion is not possible.
// returns: integer value or 0 if conversion is not supported.
// see: GetAsIntegerWithDefault
func (c *AnyValue) GetAsInteger() int {
	return c.GetAsIntegerWithDefault(0)
}

// Converts object value into a integer or returns default value if conversion is not possible.
// param "defaultValue": the default value
// returns: integer value or default if conversion is not supported.
// see: IntegerConverter.ToIntegerWithDefault
func (c *AnyValue) GetAsIntegerWithDefault(defaultValue int) int {
	return convert.IntegerConverter.ToIntegerWithDefault(c.value, defaultValue)
}

// Converts object value into a long or returns null if conversion is not possible.
// returns: long value or null if conversion is not supported.
// see: LongConverter.ToNullableLong
func (c *AnyValue) GetAsNullableLong() *int64 {
	return convert.LongConverter.ToNullableLong(c.value)
}

// Converts object value into a long or returns 0 if conversion is not possible.
// returns: string value or 0 if conversion is not supported.
// see: GetAsLongWithDefault
func (c *AnyValue) GetAsLong() int64 {
	return c.GetAsLongWithDefault(0)
}

// Converts object value into a long or returns default value if conversion is not possible.
// param "defaultValue": the default value
// returns: long value or default if conversion is not supported.
// see: LongConverter.ToLongWithDefault
func (c *AnyValue) GetAsLongWithDefault(defaultValue int64) int64 {
	return convert.LongConverter.ToLongWithDefault(c.value, defaultValue)
}

// Converts object value into a float or returns null if conversion is not possible.
// returns: float value or null if conversion is not supported.
// see: FloatConverter.ToNullableFloat
func (c *AnyValue) GetAsNullableFloat() *float32 {
	return convert.FloatConverter.ToNullableFloat(c.value)
}

// Converts object value into a float or returns 0 if conversion is not possible.
// returns: float value or 0 if conversion is not supported.
// see: GetAsFloatWithDefault
func (c *AnyValue) GetAsFloat() float32 {
	return c.GetAsFloatWithDefault(0)
}

// Converts object value into a float or returns default value if conversion is not possible.
// param "defaultValue": the default value
// returns: float value or default if conversion is not supported.
// see: FloatConverter.ToFloatWithDefault
func (c *AnyValue) GetAsFloatWithDefault(defaultValue float32) float32 {
	return convert.FloatConverter.ToFloatWithDefault(c.value, defaultValue)
}

// Converts object value into a double or returns null if conversion is not possible.
// returns: double value or null if conversion is not supported.
// see: DoubleConverter.ToNullableDouble
func (c *AnyValue) GetAsNullableDouble() *float64 {
	return convert.DoubleConverter.ToNullableDouble(c.value)
}

// Converts object value into a double or returns 0 if conversion is not possible.
// returns: double value or 0 if conversion is not supported.
// see: GetAsDoubleWithDefault
func (c *AnyValue) GetAsDouble() float64 {
	return c.GetAsDoubleWithDefault(0)
}

// Converts object value into a double or returns default value if conversion is not possible.
// param "defaultValue": the default value
// returns: double value or default if conversion is not supported.
// see: DoubleConverter.ToDoubleWithDefault
func (c *AnyValue) GetAsDoubleWithDefault(defaultValue float64) float64 {
	return convert.DoubleConverter.ToDoubleWithDefault(c.value, defaultValue)
}

// Converts object value into a Date or returns null if conversion is not possible.
// returns: DateTime value or null if conversion is not supported.
// see: DateTimeConverter.ToNullableDateTime
func (c *AnyValue) GetAsNullableDateTime() *time.Time {
	return convert.DateTimeConverter.ToNullableDateTime(c.value)
}

// Converts object value into a Date or returns current date if conversion is not possible.
// returns: DateTime value or current date if conversion is not supported.
// see: GetAsDateTimeWithDefault
func (c *AnyValue) GetAsDateTime() time.Time {
	return c.GetAsDateTimeWithDefault(time.Time{})
}

// Converts object value into a Date or returns default value if conversion is not possible.
// param "defaultValue": the default value
// returns: DateTime value or default if conversion is not supported.
// see: DateTimeConverter.ToDateTimeWithDefault
func (c *AnyValue) GetAsDateTimeWithDefault(defaultValue time.Time) time.Time {
	return convert.DateTimeConverter.ToDateTimeWithDefault(c.value, defaultValue)
}

// Converts object value into a Duration or returns null if conversion is not possible.
// returns: Duration value or null if conversion is not supported.
// see: DurationConverter.ToNullableDuration
func (c *AnyValue) GetAsNullableDuration() *time.Duration {
	return convert.DurationConverter.ToNullableDuration(c.value)
}

// Converts object value into a Duration or returns current date if conversion is not possible.
// returns: Duration value or current date if conversion is not supported.
// see: GetAsDurationWithDefault
func (c *AnyValue) GetAsDuration() time.Duration {
	return c.GetAsDurationWithDefault(0 * time.Millisecond)
}

// Converts object value into a Duration or returns default value if conversion is not possible.
// param "defaultValue": the default value
// returns: Duration value or default if conversion is not supported.
// see: DurationConverter.ToDurationWithDefault
func (c *AnyValue) GetAsDurationWithDefault(defaultValue time.Duration) time.Duration {
	return convert.DurationConverter.ToDurationWithDefault(c.value, defaultValue)
}

// Converts object value into a value defined by specied typecode. If conversion is not possible it returns null.
// param "typ": the TypeCode that defined the type of the result.
// returns: value defined by the typecode or null if conversion is not supported.
// see: TypeConverter.ToNullableType
func (c *AnyValue) GetAsNullableType(typ convert.TypeCode) interface{} {
	return convert.TypeConverter.ToNullableType(typ, c.value)
}

// Converts object value into a value defined by specied typecode. If conversion
// is not possible it returns default value for the specified type.
// param "typ": the TypeCode that defined the type of the result.
// returns: value defined by the typecode or type default value if conversion is not supported.
// see: GetAsTypeWithDefault
func (c *AnyValue) GetAsType(typ convert.TypeCode) interface{} {
	return c.GetAsTypeWithDefault(typ, nil)
}

// Converts object value into a value defined by specied typecode. If conversion
// is not possible it returns default value.
// param "typ": the TypeCode that defined the type of the result.
// param "defaultValue": the default value.
// returns: value defined by the typecode or type default value if conversion is not supported.
// see: TypeConverter.ToTypeWithDefault
func (c *AnyValue) GetAsTypeWithDefault(typ convert.TypeCode, defaultValue interface{}) interface{} {
	return convert.TypeConverter.ToTypeWithDefault(typ, c.value, defaultValue)
}

// Converts object value into an AnyArray or returns empty AnyArray if conversion is not possible.
// returns: AnyArray value or empty AnyArray if conversion is not supported.
// see: NewAnyValueArrayFromValue
func (c *AnyValue) GetAsArray() *AnyValueArray {
	return NewAnyValueArrayFromValue(c.value)
}

// Converts object value into AnyMap or returns empty AnyMap if conversion is not possible.
// returns: AnyMap value or empty AnyMap if conversion is not supported.
// see: NewAnyValueMapFromValue
func (c *AnyValue) GetAsMap() *AnyValueMap {
	return NewAnyValueMapFromValue(c.value)
}

// Compares this object value to specified specified value. When direct
// comparison gives negative results it tries to compare values as strings.
// param "obj": the value to be compared with.
// returns: true when objects are equal and false otherwise.
// see: StringConverter.ToNullableString
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
// param "typ": the TypeCode that defined the type of the result.
// param "obj": the value to be compared with.
// returns: true when objects are equal and false otherwise.
// see: TypeConverter.ToType
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
// returns: a clone of this object.
// see: NewAnyValue
func (c *AnyValue) Clone() interface{} {
	return NewAnyValue(c.value)
}

// Gets a string representation of the object.
// returns: a string representation of the object.
// see: StringConverter.ToString
func (c *AnyValue) String() string {
	return convert.StringConverter.ToString(c.value)
}
