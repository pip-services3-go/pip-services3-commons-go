package convert

import (
	"reflect"
	"time"
)

// Converts arbitrary values into objects specific by TypeCodes.
// For each TypeCode this class calls corresponding converter
// which applies extended conversion rules to convert the values.
//
// Example:
//
//  value1 := convert.TypeConverter.ToType(convert.Integer, "123.456")
//  value2 := convert.TypeConverter.ToType(convert.DateTime, 123)
//  value3 := convert.TypeConverter.ToType(convert.Boolean, "F")
//  fmt.Println(value1) // 123
//  fmt.Println(value2) // 1970-01-01 02:02:03 +0200 EET
//  fmt.Println(value3) // false
type TTypeConverter struct{}

var TypeConverter *TTypeConverter = &TTypeConverter{}

// Gets TypeCode for specific value.
// Parameters: "value" - value whose TypeCode is to be resolved.
// Returns: the TypeCode that corresponds to the passed object's type.
func (c *TTypeConverter) ToTypeCode(value interface{}) TypeCode {
	return ToTypeCode(value)
}

// Converts value into an object type specified by Type Code or returns null
// when conversion is not possible.
// Parameters:
//  "typ" - the TypeCode for the data type.
//  "value" - the value to convert.
// Returns: object value of type corresponding to TypeCode, or null when
// conversion is not supported.
func (c *TTypeConverter) ToNullableType(typ TypeCode, value interface{}) interface{} {
	return ToNullableType(typ, value)
}

// Converts value into an object type specified by Type Code
// or returns default value when conversion is not possible.
// Parameters:
//  "typ" - the TypeCode for the data type into which 'value' is to be converted.
//  "value" - the value to convert.
// Returns: object value of type corresponding to TypeCode, or default value when
// conversion is not supported
func (c *TTypeConverter) ToType(typ TypeCode, value interface{}) interface{} {
	return ToType(typ, value)
}

// Converts value into an object type specified by Type Code
// or returns default value when conversion is not possible.
// Parameters:
//  "typ" - the TypeCode for the data type into which 'value' is to be converted.
//  "value" - the value to convert.
//  "defaultValue" - the default value to return if conversion is not possible
//  (returns null).
// Returns: object value of type corresponding to TypeCode, or default value when
// conversion is not supported
func (c *TTypeConverter) ToTypeWithDefault(typ TypeCode, value interface{}, defaultValue interface{}) interface{} {
	return ToTypeWithDefault(typ, value, defaultValue)
}

// Converts a TypeCode into its string name.
// Parameters: "typ" - the TypeCode to convert into a string.
// Returns: the name of the TypeCode passed as a string value.
func (c *TTypeConverter) ToString(typ TypeCode) string {
	return TypeCodeToString(typ)
}

// Gets TypeCode for specific value.
// Parameters: "value" - value whose TypeCode is to be resolved.
// Returns: the TypeCode that corresponds to the passed object's type.
func ToTypeCode(value interface{}) TypeCode {
	if value == nil {
		return Unknown
	}

	switch value.(type) {
	case string:
		return String

	case bool:
		return Boolean

	case byte, uint, int, int32:
		return Integer

	case uint32, uint64, int64:
		return Long

	case float32:
		return Float

	case float64:
		return Double

	case time.Time:
		return DateTime

	case time.Duration:
		return Duration
	}

	rt, ok := value.(reflect.Type)
	if !ok {
		rt = reflect.TypeOf(value)
	}
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt == reflect.TypeOf((*time.Time)(nil)).Elem() {
		return DateTime
	}

	if rt == reflect.TypeOf((*time.Duration)(nil)).Elem() {
		return Duration
	}

	switch rt.Kind() {
	case reflect.String:
		return String

	case reflect.Bool:
		return Boolean

	case reflect.Int8, reflect.Uint8, reflect.Int16, reflect.Uint16,
		reflect.Int32, reflect.Uint32, reflect.Int, reflect.Uint:
		return Integer

	case reflect.Int64, reflect.Uint64:
		return Long

	case reflect.Float32:
		return Float

	case reflect.Float64:
		return Double

	case reflect.Struct:
		return Object

	case reflect.Map:
		return Map

	case reflect.Array, reflect.Slice:
		return Array

	default:
		return Unknown
	}
}

// Converts value into an object type specified by Type Code or returns null
// when conversion is not possible.
// Parameters:
//  "typ" - the TypeCode for the data type.
//  "value" - the value to convert.
// Returns: object value of type corresponding to TypeCode, or null when
// conversion is not supported.
func ToNullableType(typ TypeCode, value interface{}) interface{} {
	if value == nil {
		return nil
	}

	// Convert to known types
	if typ == String {
		return StringConverter.ToNullableString(value)
	} else if typ == Boolean {
		return BooleanConverter.ToNullableBoolean(value)
	} else if typ == Integer {
		return IntegerConverter.ToNullableInteger(value)
	} else if typ == Long {
		return LongConverter.ToNullableLong(value)
	} else if typ == Float {
		return FloatConverter.ToNullableFloat(value)
	} else if typ == Double {
		return DoubleConverter.ToNullableDouble(value)
	} else if typ == DateTime {
		return DateTimeConverter.ToNullableDateTime(value)
	} else if typ == Duration {
		return DurationConverter.ToNullableDuration(value)
	} else if typ == Array {
		return ArrayConverter.ToNullableArray(value)
	} else if typ == Map {
		return MapConverter.ToNullableMap(value)
	} else {
		return nil
	}
}

// Converts value into an object type specified by Type Code
// or returns default value when conversion is not possible.
// Parameters:
//  "typ" - the TypeCode for the data type into which 'value' is to be converted.
//  "value" - the value to convert.
// Returns: object value of type corresponding to TypeCode, or default value when
// conversion is not supported
func ToType(typ TypeCode, value interface{}) interface{} {
	if value == nil {
		return nil
	}

	// Convert to known types
	if typ == String {
		return StringConverter.ToString(value)
	} else if typ == Boolean {
		return BooleanConverter.ToBoolean(value)
	} else if typ == Integer {
		return IntegerConverter.ToInteger(value)
	} else if typ == Long {
		return LongConverter.ToLong(value)
	} else if typ == Float {
		return FloatConverter.ToFloat(value)
	} else if typ == Double {
		return DoubleConverter.ToDouble(value)
	} else if typ == DateTime {
		return DateTimeConverter.ToDateTime(value)
	} else if typ == Duration {
		return DurationConverter.ToDuration(value)
	} else if typ == Array {
		return ArrayConverter.ToArray(value)
	} else if typ == Map {
		return MapConverter.ToMap(value)
	} else {
		return value
	}
}

// Converts value into an object type specified by Type Code
// or returns default value when conversion is not possible.
// Parameters:
//  "typ" - the TypeCode for the data type into which 'value' is to be converted.
//  "value" - the value to convert.
//  "defaultValue" - the default value to return if conversion is not possible
//  (returns null).
// Returns: object value of type corresponding to TypeCode, or default value when
// conversion is not supported
func ToTypeWithDefault(typ TypeCode, value interface{}, defaultValue interface{}) interface{} {
	if value == nil {
		return defaultValue
	}

	// Convert to known types
	if typ == String {
		return ToStringWithDefault(value, defaultValue.(string))
	} else if typ == Boolean {
		return BooleanConverter.ToBooleanWithDefault(value, defaultValue.(bool))
	} else if typ == Integer {
		return IntegerConverter.ToIntegerWithDefault(value, defaultValue.(int))
	} else if typ == Long {
		return LongConverter.ToLongWithDefault(value, defaultValue.(int64))
	} else if typ == Float {
		return FloatConverter.ToFloatWithDefault(value, defaultValue.(float32))
	} else if typ == Double {
		return DoubleConverter.ToDoubleWithDefault(value, defaultValue.(float64))
	} else if typ == DateTime {
		return DateTimeConverter.ToDateTimeWithDefault(value, defaultValue.(time.Time))
	} else if typ == Duration {
		return DurationConverter.ToDurationWithDefault(value, defaultValue.(time.Duration))
	} else if typ == Array {
		return ArrayConverter.ToArrayWithDefault(value, defaultValue.([]interface{}))
	} else if typ == Map {
		return MapConverter.ToMapWithDefault(value, defaultValue.(map[string]interface{}))
	} else {
		return defaultValue
	}
}

// Converts a TypeCode into its string name.
// Parameters: "typ" - the TypeCode to convert into a string.
// Returns: the name of the TypeCode passed as a string value.
func TypeCodeToString(typ TypeCode) string {
	switch typ {
	case Unknown:
		return "unknown"
	case String:
		return "string"
	case Boolean:
		return "boolean"
	case Integer:
		return "integer"
	case Long:
		return "long"
	case Float:
		return "float"
	case Double:
		return "double"
	case DateTime:
		return "datetime"
	case Duration:
		return "duration"
	case Object:
		return "object"
	case Enum:
		return "enum"
	case Array:
		return "array"
	case Map:
		return "map"
	default:
		return "unknown"
	}
}
