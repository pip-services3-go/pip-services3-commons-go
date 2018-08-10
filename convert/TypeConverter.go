package convert

import (
	"reflect"
	"time"
)

type TTypeConverter struct{}

var TypeConverter *TTypeConverter = &TTypeConverter{}

func (c *TTypeConverter) ToTypeCode(value interface{}) TypeCode {
	return ToTypeCode(value)
}

func (c *TTypeConverter) ToNullableType(typ TypeCode, value interface{}) interface{} {
	return ToNullableType(typ, value)
}

func (c *TTypeConverter) ToType(typ TypeCode, value interface{}) interface{} {
	return ToType(typ, value)
}

func (c *TTypeConverter) ToTypeWithDefault(typ TypeCode, value interface{}, defaultValue interface{}) interface{} {
	return ToTypeWithDefault(typ, value, defaultValue)
}

func (c *TTypeConverter) ToString(typ TypeCode) string {
	return TypeCodeToString(typ)
}

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

func ToNullableType(typ TypeCode, value interface{}) interface{} {
	if value == nil {
		return nil
	}

	// Convert to known types
	if typ == String {
		return StringConverter.ToNullableString(value)
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

func ToType(typ TypeCode, value interface{}) interface{} {
	if value == nil {
		return nil
	}

	// Convert to known types
	if typ == String {
		return StringConverter.ToString(value)
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

func ToTypeWithDefault(typ TypeCode, value interface{}, defaultValue interface{}) interface{} {
	if value == nil {
		return defaultValue
	}

	// Convert to known types
	if typ == String {
		return ToStringWithDefault(value, defaultValue.(string))
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
