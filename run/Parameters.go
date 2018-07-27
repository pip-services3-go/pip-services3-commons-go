package data

import (
	"fmt"
	"time"

	"github.com/pip-services-go/pip-services-commons-go/convert"
)

type Parameters map[string]interface{}

func NewEmptyParameters() *Parameters {
	return &Parameters{}
}

func NewParameters(values map[string]interface{}) *Parameters {
	c := &Parameters{}
	c.Append(values)
	return c
}

func (c *Parameters) Get(key string) interface{} {
	// Todo: Make this method recursive
	// if key == "" {
	// 	return nil
	// } else if strings.Index(key, ".") > 0 {
	// 	return RecursiveObjectReader.GetProperty(c, key)
	// } else {
	// 	return (*c)[key]
	// }
	return (*c)[key]
}

func (c *Parameters) Keys() []string {
	keys := []string{}
	for key := range *c {
		keys = append(keys, key)
	}
	return keys
}

func (c *Parameters) Put(key string, value interface{}) {
	// Todo: Make this method recursive
	(*c)[key] = value
}

func (c *Parameters) Remove(key string) {
	// Todo: Make this method recursive
	delete(*c, key)
}

func (c *Parameters) Append(values map[string]interface{}) {
	if values == nil {
		return
	}

	for key := range values {
		(*c)[key] = values[key]
	}
}

func (c *Parameters) Clear() {
	*c = Parameters{}
}

func (c *Parameters) Length() int {
	return len(*c)
}

func (c *Parameters) GetAsSingleObject() interface{} {
	return *c
}

func (c *Parameters) SetAsSingleObject(value interface{}) {
	a := convert.ToMap(value)
	*c = a
}

func (c *Parameters) GetAsObject(key string) interface{} {
	return c.Get(key)
}

func (c *Parameters) SetAsObject(key string, value interface{}) {
	c.Put(key, value)
}

func (c *Parameters) GetAsNullableString(key string) *string {
	value := c.Get(key)
	return convert.StringConverter.ToNullableString(value)
}

func (c *Parameters) GetAsString(key string) string {
	return c.GetAsStringWithDefault(key, "")
}

func (c *Parameters) GetAsStringWithDefault(key string, defaultValue string) string {
	value := c.Get(key)
	return convert.StringConverter.ToStringWithDefault(value, defaultValue)
}

func (c *Parameters) GetAsNullableBoolean(key string) *bool {
	value := c.Get(key)
	return convert.BooleanConverter.ToNullableBoolean(value)
}

func (c *Parameters) GetAsBoolean(key string) bool {
	return c.GetAsBooleanWithDefault(key, false)
}

func (c *Parameters) GetAsBooleanWithDefault(key string, defaultValue bool) bool {
	value := c.Get(key)
	return convert.BooleanConverter.ToBooleanWithDefault(value, defaultValue)
}

func (c *Parameters) GetAsNullableInteger(key string) *int {
	value := c.Get(key)
	return convert.IntegerConverter.ToNullableInteger(value)
}

func (c *Parameters) GetAsInteger(key string) int {
	return c.GetAsIntegerWithDefault(key, 0)
}

func (c *Parameters) GetAsIntegerWithDefault(key string, defaultValue int) int {
	value := c.Get(key)
	return convert.IntegerConverter.ToIntegerWithDefault(value, defaultValue)
}

func (c *Parameters) GetAsNullableLong(key string) *int64 {
	value := c.Get(key)
	return convert.LongConverter.ToNullableLong(value)
}

func (c *Parameters) GetAsLong(key string) int64 {
	return c.GetAsLongWithDefault(key, 0)
}

func (c *Parameters) GetAsLongWithDefault(key string, defaultValue int64) int64 {
	value := c.Get(key)
	return convert.LongConverter.ToLongWithDefault(value, defaultValue)
}

func (c *Parameters) GetAsNullableFloat(key string) *float32 {
	value := c.Get(key)
	return convert.FloatConverter.ToNullableFloat(value)
}

func (c *Parameters) GetAsFloat(key string) float32 {
	return c.GetAsFloatWithDefault(key, 0)
}

func (c *Parameters) GetAsFloatWithDefault(key string, defaultValue float32) float32 {
	value := c.Get(key)
	return convert.FloatConverter.ToFloatWithDefault(value, defaultValue)
}

func (c *Parameters) GetAsNullableDouble(key string) *float64 {
	value := c.Get(key)
	return convert.DoubleConverter.ToNullableDouble(value)
}

func (c *Parameters) GetAsDouble(key string) float64 {
	return c.GetAsDoubleWithDefault(key, 0)
}

func (c *Parameters) GetAsDoubleWithDefault(key string, defaultValue float64) float64 {
	value := c.Get(key)
	return convert.DoubleConverter.ToDoubleWithDefault(value, defaultValue)
}

func (c *Parameters) GetAsNullableDateTime(key string) *time.Time {
	value := c.Get(key)
	return convert.DateTimeConverter.ToNullableDateTime(value)
}

func (c *Parameters) GetAsDateTime(key string) time.Time {
	return c.GetAsDateTimeWithDefault(key, time.Time{})
}

func (c *Parameters) GetAsDateTimeWithDefault(key string, defaultValue time.Time) time.Time {
	value := c.Get(key)
	return convert.DateTimeConverter.ToDateTimeWithDefault(value, defaultValue)
}

func (c *Parameters) GetAsNullableDuration(key string) *time.Duration {
	value := c.Get(key)
	return convert.DurationConverter.ToNullableDuration(value)
}

func (c *Parameters) GetAsDuration(key string) time.Duration {
	return c.GetAsDurationWithDefault(key, 0*time.Millisecond)
}

func (c *Parameters) GetAsDurationWithDefault(key string, defaultValue time.Duration) time.Duration {
	value := c.Get(key)
	return convert.DurationConverter.ToDurationWithDefault(value, defaultValue)
}

func (c *Parameters) GetAsNullableType(typ convert.TypeCode, key string) interface{} {
	value := c.Get(key)
	return convert.TypeConverter.ToNullableType(typ, value)
}

func (c *Parameters) GetAsType(typ convert.TypeCode, key string) interface{} {
	return c.GetAsTypeWithDefault(typ, key, nil)
}

func (c *Parameters) GetAsTypeWithDefault(typ convert.TypeCode, key string, defaultValue interface{}) interface{} {
	value := c.Get(key)
	return convert.TypeConverter.ToTypeWithDefault(typ, value, defaultValue)
}

func (c *Parameters) GetAsValue(key string) *AnyValue {
	value := c.Get(key)
	return NewAnyValue(value)
}

func (c *Parameters) GetAsNullableArray(key string) *AnyValueArray {
	value := c.Get(key)
	if value != nil {
		return NewAnyValueArrayFromValue(value)
	} else {
		return nil
	}
}

func (c *Parameters) GetAsArray(key string) *AnyValueArray {
	value := c.Get(key)
	return NewAnyValueArrayFromValue(value)
}

func (c *Parameters) GetAsArrayWithDefault(key string, defaultValue *AnyValueArray) *AnyValueArray {
	result := c.GetAsNullableArray(key)
	if result != nil {
		return result
	} else {
		return defaultValue
	}
}

func (c *Parameters) GetAsNullableMap(key string) *Parameters {
	value := c.Get(key)
	if value != nil {
		return NewParametersFromValue(value)
	} else {
		return nil
	}
}

func (c *Parameters) GetAsMap(key string) *Parameters {
	value := c.Get(key)
	return NewParametersFromValue(value)
}

func (c *Parameters) GetAsMapWithDefault(key string, defaultValue *Parameters) *Parameters {
	result := c.GetAsNullableMap(key)
	if result != nil {
		return result
	} else {
		return defaultValue
	}
}

func (c *Parameters) String() string {
	builder := ""

	// Todo: User encoder
	for key := range *c {
		value := (*c)[key]

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

func (c *Parameters) Clone() interface{} {
	return NewParameters(*c)
}

func NewParametersFromValue(value interface{}) *Parameters {
	result := NewEmptyParameters()
	result.SetAsSingleObject(value)
	return result
}

func NewParametersFromTuples(tuples ...interface{}) *Parameters {
	return NewParametersFromTuplesArray(tuples)
}

func NewParametersFromTuplesArray(tuples []interface{}) *Parameters {
	result := NewEmptyParameters()
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

func NewParametersFromMaps(maps ...map[string]interface{}) *Parameters {
	result := NewEmptyParameters()
	if len(maps) > 0 {
		for index := 0; index < len(maps); index++ {
			result.Append(maps[index])
		}
	}
	return result
}
