package data

import (
	"time"

	"github.com/pip-services-go/pip-services-commons-go/convert"
)

type AnyValue struct {
	value interface{}
}

func NewEmptyAnyValue() *AnyValue {
	return &AnyValue{value: nil}
}

func NewAnyValue(value interface{}) *AnyValue {
	v, ok := value.(*AnyValue)
	if ok {
		return v
	} else {
		return &AnyValue{value: value}
	}
}

func (c *AnyValue) InnerValue() interface{} {
	return c.value
}

func (c *AnyValue) Value() interface{} {
	return c.value
}

func (c *AnyValue) TypeCode() convert.TypeCode {
	return convert.TypeConverter.ToTypeCode(c.value)
}

func (c *AnyValue) GetAsObject() interface{} {
	return c.value
}

func (c *AnyValue) SetAsObject(value interface{}) {
	c.value = value
}

func (c *AnyValue) GetAsNullableString() *string {
	return convert.StringConverter.ToNullableString(c.value)
}

func (c *AnyValue) GetAsString() string {
	return c.GetAsStringWithDefault("")
}

func (c *AnyValue) GetAsStringWithDefault(defaultValue string) string {
	return convert.StringConverter.ToStringWithDefault(c.value, defaultValue)
}

func (c *AnyValue) GetAsNullableBoolean() *bool {
	return convert.BooleanConverter.ToNullableBoolean(c.value)
}

func (c *AnyValue) GetAsBoolean() bool {
	return c.GetAsBooleanWithDefault(false)
}

func (c *AnyValue) GetAsBooleanWithDefault(defaultValue bool) bool {
	return convert.BooleanConverter.ToBooleanWithDefault(c.value, defaultValue)
}

func (c *AnyValue) GetAsNullableInteger() *int {
	return convert.IntegerConverter.ToNullableInteger(c.value)
}

func (c *AnyValue) GetAsInteger() int {
	return c.GetAsIntegerWithDefault(0)
}

func (c *AnyValue) GetAsIntegerWithDefault(defaultValue int) int {
	return convert.IntegerConverter.ToIntegerWithDefault(c.value, defaultValue)
}

func (c *AnyValue) GetAsNullableLong() *int64 {
	return convert.LongConverter.ToNullableLong(c.value)
}

func (c *AnyValue) GetAsLong() int64 {
	return c.GetAsLongWithDefault(0)
}

func (c *AnyValue) GetAsLongWithDefault(defaultValue int64) int64 {
	return convert.LongConverter.ToLongWithDefault(c.value, defaultValue)
}

func (c *AnyValue) GetAsNullableFloat() *float32 {
	return convert.FloatConverter.ToNullableFloat(c.value)
}

func (c *AnyValue) GetAsFloat() float32 {
	return c.GetAsFloatWithDefault(0)
}

func (c *AnyValue) GetAsFloatWithDefault(defaultValue float32) float32 {
	return convert.FloatConverter.ToFloatWithDefault(c.value, defaultValue)
}

func (c *AnyValue) GetAsNullableDouble() *float64 {
	return convert.DoubleConverter.ToNullableDouble(c.value)
}

func (c *AnyValue) GetAsDouble() float64 {
	return c.GetAsDoubleWithDefault(0)
}

func (c *AnyValue) GetAsDoubleWithDefault(defaultValue float64) float64 {
	return convert.DoubleConverter.ToDoubleWithDefault(c.value, defaultValue)
}

func (c *AnyValue) GetAsNullableDateTime() *time.Time {
	return convert.DateTimeConverter.ToNullableDateTime(c.value)
}

func (c *AnyValue) GetAsDateTime() time.Time {
	return c.GetAsDateTimeWithDefault(time.Time{})
}

func (c *AnyValue) GetAsDateTimeWithDefault(defaultValue time.Time) time.Time {
	return convert.DateTimeConverter.ToDateTimeWithDefault(c.value, defaultValue)
}

func (c *AnyValue) GetAsNullableDuration() *time.Duration {
	return convert.DurationConverter.ToNullableDuration(c.value)
}

func (c *AnyValue) GetAsDuration() time.Duration {
	return c.GetAsDurationWithDefault(0 * time.Millisecond)
}

func (c *AnyValue) GetAsDurationWithDefault(defaultValue time.Duration) time.Duration {
	return convert.DurationConverter.ToDurationWithDefault(c.value, defaultValue)
}

func (c *AnyValue) GetAsNullableType(typ convert.TypeCode) interface{} {
	return convert.TypeConverter.ToNullableType(typ, c.value)
}

func (c *AnyValue) GetAsType(typ convert.TypeCode) interface{} {
	return c.GetAsTypeWithDefault(typ, nil)
}

func (c *AnyValue) GetAsTypeWithDefault(typ convert.TypeCode, defaultValue interface{}) interface{} {
	return convert.TypeConverter.ToTypeWithDefault(typ, c.value, defaultValue)
}

func (c *AnyValue) GetAsArray() *AnyValueArray {
	return NewAnyValueArrayFromValue(c.value)
}

func (c *AnyValue) GetAsMap() *AnyValueMap {
	return NewAnyValueMapFromValue(c.value)
}

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

func (c *AnyValue) Clone() interface{} {
	return NewAnyValue(c.value)
}

func (c *AnyValue) String() string {
	return convert.StringConverter.ToString(c.value)
}
