package data

import (
	"fmt"
	"time"

	"github.com/pip-services-go/pip-services-commons-go/convert"
)

type AnyValueMap struct {
	values map[string]interface{}
}

func NewEmptyAnyValueMap() *AnyValueMap {
	return &AnyValueMap{
		values: map[string]interface{}{},
	}
}

func NewAnyValueMap(values map[string]interface{}) *AnyValueMap {
	c := &AnyValueMap{
		values: map[string]interface{}{},
	}
	c.Append(values)
	return c
}

func (c *AnyValueMap) Get(key string) interface{} {
	return (*c).values[key]
}

func (c *AnyValueMap) Keys() []string {
	keys := []string{}
	for key := range (*c).values {
		keys = append(keys, key)
	}
	return keys
}

func (c *AnyValueMap) Values() map[string]interface{} {
	return (*c).values
}

func (c *AnyValueMap) Put(key string, value interface{}) {
	(*c).values[key] = value
}

func (c *AnyValueMap) Remove(key string) {
	delete((*c).values, key)
}

func (c *AnyValueMap) Append(values map[string]interface{}) {
	if values == nil {
		return
	}

	for key := range values {
		(*c).values[key] = values[key]
	}
}

func (c *AnyValueMap) Clear() {
	(*c).values = map[string]interface{}{}
}

func (c *AnyValueMap) Length() int {
	return len((*c).values)
}

func (c *AnyValueMap) GetAsSingleObject() interface{} {
	return *c
}

func (c *AnyValueMap) SetAsSingleObject(value interface{}) {
	a := convert.ToMap(value)
	(*c).values = a
}

func (c *AnyValueMap) GetAsObject(key string) interface{} {
	return c.Get(key)
}

func (c *AnyValueMap) SetAsObject(key string, value interface{}) {
	c.Put(key, value)
}

func (c *AnyValueMap) GetAsNullableString(key string) *string {
	value := c.Get(key)
	return convert.StringConverter.ToNullableString(value)
}

func (c *AnyValueMap) GetAsString(key string) string {
	return c.GetAsStringWithDefault(key, "")
}

func (c *AnyValueMap) GetAsStringWithDefault(key string, defaultValue string) string {
	value := c.Get(key)
	return convert.StringConverter.ToStringWithDefault(value, defaultValue)
}

func (c *AnyValueMap) GetAsNullableBoolean(key string) *bool {
	value := c.Get(key)
	return convert.BooleanConverter.ToNullableBoolean(value)
}

func (c *AnyValueMap) GetAsBoolean(key string) bool {
	return c.GetAsBooleanWithDefault(key, false)
}

func (c *AnyValueMap) GetAsBooleanWithDefault(key string, defaultValue bool) bool {
	value := c.Get(key)
	return convert.BooleanConverter.ToBooleanWithDefault(value, defaultValue)
}

func (c *AnyValueMap) GetAsNullableInteger(key string) *int {
	value := c.Get(key)
	return convert.IntegerConverter.ToNullableInteger(value)
}

func (c *AnyValueMap) GetAsInteger(key string) int {
	return c.GetAsIntegerWithDefault(key, 0)
}

func (c *AnyValueMap) GetAsIntegerWithDefault(key string, defaultValue int) int {
	value := c.Get(key)
	return convert.IntegerConverter.ToIntegerWithDefault(value, defaultValue)
}

func (c *AnyValueMap) GetAsNullableLong(key string) *int64 {
	value := c.Get(key)
	return convert.LongConverter.ToNullableLong(value)
}

func (c *AnyValueMap) GetAsLong(key string) int64 {
	return c.GetAsLongWithDefault(key, 0)
}

func (c *AnyValueMap) GetAsLongWithDefault(key string, defaultValue int64) int64 {
	value := c.Get(key)
	return convert.LongConverter.ToLongWithDefault(value, defaultValue)
}

func (c *AnyValueMap) GetAsNullableFloat(key string) *float32 {
	value := c.Get(key)
	return convert.FloatConverter.ToNullableFloat(value)
}

func (c *AnyValueMap) GetAsFloat(key string) float32 {
	return c.GetAsFloatWithDefault(key, 0)
}

func (c *AnyValueMap) GetAsFloatWithDefault(key string, defaultValue float32) float32 {
	value := c.Get(key)
	return convert.FloatConverter.ToFloatWithDefault(value, defaultValue)
}

func (c *AnyValueMap) GetAsNullableDouble(key string) *float64 {
	value := c.Get(key)
	return convert.DoubleConverter.ToNullableDouble(value)
}

func (c *AnyValueMap) GetAsDouble(key string) float64 {
	return c.GetAsDoubleWithDefault(key, 0)
}

func (c *AnyValueMap) GetAsDoubleWithDefault(key string, defaultValue float64) float64 {
	value := c.Get(key)
	return convert.DoubleConverter.ToDoubleWithDefault(value, defaultValue)
}

func (c *AnyValueMap) GetAsNullableDateTime(key string) *time.Time {
	value := c.Get(key)
	return convert.DateTimeConverter.ToNullableDateTime(value)
}

func (c *AnyValueMap) GetAsDateTime(key string) time.Time {
	return c.GetAsDateTimeWithDefault(key, time.Time{})
}

func (c *AnyValueMap) GetAsDateTimeWithDefault(key string, defaultValue time.Time) time.Time {
	value := c.Get(key)
	return convert.DateTimeConverter.ToDateTimeWithDefault(value, defaultValue)
}

func (c *AnyValueMap) GetAsNullableDuration(key string) *time.Duration {
	value := c.Get(key)
	return convert.DurationConverter.ToNullableDuration(value)
}

func (c *AnyValueMap) GetAsDuration(key string) time.Duration {
	return c.GetAsDurationWithDefault(key, 0*time.Millisecond)
}

func (c *AnyValueMap) GetAsDurationWithDefault(key string, defaultValue time.Duration) time.Duration {
	value := c.Get(key)
	return convert.DurationConverter.ToDurationWithDefault(value, defaultValue)
}

func (c *AnyValueMap) GetAsNullableType(typ convert.TypeCode, key string) interface{} {
	value := c.Get(key)
	return convert.TypeConverter.ToNullableType(typ, value)
}

func (c *AnyValueMap) GetAsType(typ convert.TypeCode, key string) interface{} {
	return c.GetAsTypeWithDefault(typ, key, nil)
}

func (c *AnyValueMap) GetAsTypeWithDefault(typ convert.TypeCode, key string, defaultValue interface{}) interface{} {
	value := c.Get(key)
	return convert.TypeConverter.ToTypeWithDefault(typ, value, defaultValue)
}

func (c *AnyValueMap) GetAsValue(key string) *AnyValue {
	value := c.Get(key)
	return NewAnyValue(value)
}

func (c *AnyValueMap) GetAsNullableArray(key string) *AnyValueArray {
	value := c.Get(key)
	if value != nil {
		return NewAnyValueArrayFromValue(value)
	} else {
		return nil
	}
}

func (c *AnyValueMap) GetAsArray(key string) *AnyValueArray {
	value := c.Get(key)
	return NewAnyValueArrayFromValue(value)
}

func (c *AnyValueMap) GetAsArrayWithDefault(key string, defaultValue *AnyValueArray) *AnyValueArray {
	result := c.GetAsNullableArray(key)
	if result != nil {
		return result
	} else {
		return defaultValue
	}
}

func (c *AnyValueMap) GetAsNullableMap(key string) *AnyValueMap {
	value := c.Get(key)
	if value != nil {
		return NewAnyValueMapFromValue(value)
	} else {
		return nil
	}
}

func (c *AnyValueMap) GetAsMap(key string) *AnyValueMap {
	value := c.Get(key)
	return NewAnyValueMapFromValue(value)
}

func (c *AnyValueMap) GetAsMapWithDefault(key string, defaultValue *AnyValueMap) *AnyValueMap {
	result := c.GetAsNullableMap(key)
	if result != nil {
		return result
	} else {
		return defaultValue
	}
}

func (c *AnyValueMap) String() string {
	builder := ""

	// Todo: User encoder
	for key := range c.Values() {
		value := c.Get(key)

		if len(builder) > 0 {
			builder = builder + ";"
		}

		if value != nil {
			builder = builder + fmt.Sprintf("%s=%v", key, value)
		} else {
			builder = builder + key
		}
	}

	return builder
}

func (c *AnyValueMap) Clone() interface{} {
	return NewAnyValueMap((*c).values)
}

func NewAnyValueMapFromValue(value interface{}) *AnyValueMap {
	result := NewEmptyAnyValueMap()
	result.SetAsSingleObject(value)
	return result
}

func NewAnyValueMapFromTuples(tuples ...interface{}) *AnyValueMap {
	return NewAnyValueMapFromTuplesArray(tuples)
}

func NewAnyValueMapFromTuplesArray(tuples []interface{}) *AnyValueMap {
	result := NewEmptyAnyValueMap()
	if len(tuples) == 0 {
		return result
	}

	for index := 0; index < len(tuples); index = index + 2 {
		if index+1 >= len(tuples) {
			break
		}

		name := convert.StringConverter.ToString(tuples[index])
		value := tuples[index+1]

		result.SetAsObject(name, value)
	}

	return result
}

func NewAnyValueMapFromMaps(maps ...map[string]interface{}) *AnyValueMap {
	result := NewEmptyAnyValueMap()
	if len(maps) > 0 {
		for index := 0; index < len(maps); index++ {
			result.Append(maps[index])
		}
	}
	return result
}
