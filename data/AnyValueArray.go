package data

import (
	"strings"
	"time"

	"github.com/pip-services-go/pip-services-commons-go/convert"
)

type AnyValueArray struct {
	values []interface{}
}

func NewEmptyAnyValueArray() *AnyValueArray {
	return &AnyValueArray{
		values: make([]interface{}, 0, 10),
	}
}

func NewAnyValueArray(values []interface{}) *AnyValueArray {
	c := &AnyValueArray{
		values: make([]interface{}, len(values)),
	}
	copy((*c).values, values)
	return c
}

func (c *AnyValueArray) Values() []interface{} {
	return (*c).values
}

func (c *AnyValueArray) Length() int {
	return len((*c).values)
}

func (c *AnyValueArray) Get(index int) interface{} {
	return (*c).values[index]
}

func (c *AnyValueArray) Put(index int, value interface{}) {
	if cap((*c).values)+1 < index {
		a := make([]interface{}, index+1, (index+1)*2)
		copy(a, (*c).values)
		(*c).values = a
	}

	(*c).values[index] = value
}

func (c *AnyValueArray) Remove(index int) {
	(*c).values = append((*c).values[:index], (*c).values[index+1:]...)
}

func (c *AnyValueArray) Push(value interface{}) {
	(*c).values = append((*c).values, value)
}

func (c *AnyValueArray) Append(elements []interface{}) {
	if elements != nil {
		(*c).values = append((*c).values, elements...)
	}
}

func (c *AnyValueArray) Clear() {
	(*c).values = make([]interface{}, 0, 10)
}

func (c *AnyValueArray) GetAsSingleObject() interface{} {
	return *c
}

func (c *AnyValueArray) SetAsSingleObject(value interface{}) {
	a := convert.ToArray(value)
	(*c).values = a
}

func (c *AnyValueArray) GetAsObject(index int) interface{} {
	return c.Get(index)
}

func (c *AnyValueArray) SetAsObject(index int, value interface{}) {
	c.Put(index, value)
}

func (c *AnyValueArray) GetAsNullableString(index int) *string {
	value := c.Get(index)
	return convert.StringConverter.ToNullableString(value)
}

func (c *AnyValueArray) GetAsString(index int) string {
	return c.GetAsStringWithDefault(index, "")
}

func (c *AnyValueArray) GetAsStringWithDefault(index int, defaultValue string) string {
	value := c.Get(index)
	return convert.StringConverter.ToStringWithDefault(value, defaultValue)
}

func (c *AnyValueArray) GetAsNullableBoolean(index int) *bool {
	value := c.Get(index)
	return convert.BooleanConverter.ToNullableBoolean(value)
}

func (c *AnyValueArray) GetAsBoolean(index int) bool {
	return c.GetAsBooleanWithDefault(index, false)
}

func (c *AnyValueArray) GetAsBooleanWithDefault(index int, defaultValue bool) bool {
	value := c.Get(index)
	return convert.BooleanConverter.ToBooleanWithDefault(value, defaultValue)
}

func (c *AnyValueArray) GetAsNullableInteger(index int) *int {
	value := c.Get(index)
	return convert.IntegerConverter.ToNullableInteger(value)
}

func (c *AnyValueArray) GetAsInteger(index int) int {
	return c.GetAsIntegerWithDefault(index, 0)
}

func (c *AnyValueArray) GetAsIntegerWithDefault(index int, defaultValue int) int {
	value := c.Get(index)
	return convert.IntegerConverter.ToIntegerWithDefault(value, defaultValue)
}

func (c *AnyValueArray) GetAsNullableLong(index int) *int64 {
	value := c.Get(index)
	return convert.LongConverter.ToNullableLong(value)
}

func (c *AnyValueArray) GetAsLong(index int) int64 {
	return c.GetAsLongWithDefault(index, 0)
}

func (c *AnyValueArray) GetAsLongWithDefault(index int, defaultValue int64) int64 {
	value := c.Get(index)
	return convert.LongConverter.ToLongWithDefault(value, defaultValue)
}

func (c *AnyValueArray) GetAsNullableFloat(index int) *float32 {
	value := c.Get(index)
	return convert.FloatConverter.ToNullableFloat(value)
}

func (c *AnyValueArray) GetAsFloat(index int) float32 {
	return c.GetAsFloatWithDefault(index, 0)
}

func (c *AnyValueArray) GetAsFloatWithDefault(index int, defaultValue float32) float32 {
	value := c.Get(index)
	return convert.FloatConverter.ToFloatWithDefault(value, defaultValue)
}

func (c *AnyValueArray) GetAsNullableDouble(index int) *float64 {
	value := c.Get(index)
	return convert.DoubleConverter.ToNullableDouble(value)
}

func (c *AnyValueArray) GetAsDouble(index int) float64 {
	return c.GetAsDoubleWithDefault(index, 0)
}

func (c *AnyValueArray) GetAsDoubleWithDefault(index int, defaultValue float64) float64 {
	value := c.Get(index)
	return convert.DoubleConverter.ToDoubleWithDefault(value, defaultValue)
}

func (c *AnyValueArray) GetAsNullableDateTime(index int) *time.Time {
	value := c.Get(index)
	return convert.DateTimeConverter.ToNullableDateTime(value)
}

func (c *AnyValueArray) GetAsDateTime(index int) time.Time {
	return c.GetAsDateTimeWithDefault(index, time.Time{})
}

func (c *AnyValueArray) GetAsDateTimeWithDefault(index int, defaultValue time.Time) time.Time {
	value := c.Get(index)
	return convert.DateTimeConverter.ToDateTimeWithDefault(value, defaultValue)
}

func (c *AnyValueArray) GetAsNullableDuration(index int) *time.Duration {
	value := c.Get(index)
	return convert.DurationConverter.ToNullableDuration(value)
}

func (c *AnyValueArray) GetAsDuration(index int) time.Duration {
	return c.GetAsDurationWithDefault(index, 0*time.Millisecond)
}

func (c *AnyValueArray) GetAsDurationWithDefault(index int, defaultValue time.Duration) time.Duration {
	value := c.Get(index)
	return convert.DurationConverter.ToDurationWithDefault(value, defaultValue)
}

func (c *AnyValueArray) GetAsNullableType(typ convert.TypeCode, index int) interface{} {
	value := c.Get(index)
	return convert.TypeConverter.ToNullableType(typ, value)
}

func (c *AnyValueArray) GetAsType(typ convert.TypeCode, index int) interface{} {
	return c.GetAsTypeWithDefault(typ, index, nil)
}

func (c *AnyValueArray) GetAsTypeWithDefault(typ convert.TypeCode, index int, defaultValue interface{}) interface{} {
	value := c.Get(index)
	return convert.TypeConverter.ToTypeWithDefault(typ, value, defaultValue)
}

func (c *AnyValueArray) GetAsValue(index int) *AnyValue {
	value := c.Get(index)
	return NewAnyValue(value)
}

func (c *AnyValueArray) GetAsNullableArray(index int) *AnyValueArray {
	value := c.Get(index)
	if value != nil {
		return NewAnyValueArrayFromValue(value)
	} else {
		return nil
	}
}

func (c *AnyValueArray) GetAsArray(index int) *AnyValueArray {
	value := c.Get(index)
	return NewAnyValueArrayFromValue(value)
}

func (c *AnyValueArray) GetAsArrayWithDefault(index int, defaultValue *AnyValueArray) *AnyValueArray {
	result := c.GetAsNullableArray(index)
	if result != nil {
		return result
	} else {
		return defaultValue
	}
}

func (c *AnyValueArray) GetAsNullableMap(index int) *AnyValueMap {
	value := c.Get(index)
	if value != nil {
		return NewAnyValueMapFromValue(value)
	} else {
		return nil
	}
}

func (c *AnyValueArray) GetAsMap(index int) *AnyValueMap {
	value := c.Get(index)
	return NewAnyValueMapFromValue(value)
}

func (c *AnyValueArray) GetAsMapWithDefault(index int, defaultValue *AnyValueMap) *AnyValueMap {
	result := c.GetAsNullableMap(index)
	if result != nil {
		return NewAnyValueMapFromValue(result)
	} else {
		return defaultValue
	}
}

func (c *AnyValueArray) Contains(value interface{}) bool {
	for index := 0; index < c.Length(); index++ {
		element := c.Get(index)

		if value == nil && element == nil {
			return true
		}
		if value == nil || element == nil {
			continue
		}
		if value == element {
			return true
		}
	}

	return false
}

func (c *AnyValueArray) ContainsAsType(typ convert.TypeCode, value interface{}) bool {
	typedValue := convert.TypeConverter.ToType(typ, value)

	for index := 0; index < c.Length(); index++ {
		thisTypedValue := convert.TypeConverter.ToType(typ, c.Get(index))

		if typedValue == thisTypedValue {
			return true
		}
	}

	return false
}

func (c *AnyValueArray) Clone() interface{} {
	return NewAnyValueArray((*c).values)
}

func (c *AnyValueArray) String() string {
	builder := ""
	for index := 0; index < c.Length(); index++ {
		if index > 0 {
			builder += ","
		}
		builder = builder + c.GetAsStringWithDefault(index, "")
	}
	return builder
}

func NewAnyValueArrayFromValues(values ...interface{}) *AnyValueArray {
	return NewAnyValueArray(values)
}

func NewAnyValueArrayFromValue(value interface{}) *AnyValueArray {
	values := convert.ArrayConverter.ToArray(value)
	return NewAnyValueArray(values)
}

func NewAnyValueArrayFromString(values string, separator string, removeDuplicates bool) *AnyValueArray {
	result := NewEmptyAnyValueArray()

	if values == "" {
		return result
	}

	items := strings.Split(values, separator)
	for index := 0; index < len(items); index++ {
		item := items[index]
		if item != "" || removeDuplicates == false {
			result.Push(item)
		}
	}

	return result
}
