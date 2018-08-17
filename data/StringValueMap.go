package data

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/pip-services-go/pip-services-commons-go/convert"
)

type StringValueMap struct {
	value map[string]string
}

func NewEmptyStringValueMap() *StringValueMap {
	c := &StringValueMap{}
	c.value = map[string]string{}
	return c
}

func NewStringValueMap(value map[string]string) *StringValueMap {
	c := &StringValueMap{}
	c.value = map[string]string{}
	c.Append(value)
	return c
}

func (c *StringValueMap) InnerValue() interface{} {
	return c.value
}

func (c *StringValueMap) Value() map[string]string {
	return c.value
}

func (c *StringValueMap) Get(key string) string {
	return c.value[key]
}

func (c *StringValueMap) Keys() []string {
	keys := []string{}
	for key := range c.value {
		keys = append(keys, key)
	}
	return keys
}

func (c *StringValueMap) Put(key string, value interface{}) {
	c.value[key] = convert.StringConverter.ToString(value)
}

func (c *StringValueMap) Remove(key string) {
	delete(c.value, key)
}

func (c *StringValueMap) Contains(key string) bool {
	_, ok := c.value[key]
	return ok
}

func (c *StringValueMap) Append(values map[string]string) {
	if values == nil {
		return
	}

	for key := range values {
		c.value[key] = values[key]
	}
}

func (c *StringValueMap) AppendAny(values map[string]interface{}) {
	if values == nil {
		return
	}

	for key := range values {
		c.value[key] = convert.StringConverter.ToString(values[key])
	}
}

func (c *StringValueMap) Clear() {
	c.value = map[string]string{}
}

func (c *StringValueMap) Len() int {
	return len(c.value)
}

func (c *StringValueMap) GetAsSingleObject() interface{} {
	return *c
}

func (c *StringValueMap) SetAsSingleObject(value interface{}) {
	a := convert.ToMap(value)
	//*c = a
	c.Clear()
	c.AppendAny(a)
}

func (c *StringValueMap) GetAsObject(key string) interface{} {
	return c.Get(key)
}

func (c *StringValueMap) SetAsObject(key string, value interface{}) {
	c.Put(key, value)
}

func (c *StringValueMap) GetAsNullableString(key string) *string {
	value := c.Get(key)
	return convert.StringConverter.ToNullableString(value)
}

func (c *StringValueMap) GetAsString(key string) string {
	return c.GetAsStringWithDefault(key, "")
}

func (c *StringValueMap) GetAsStringWithDefault(key string, defaultValue string) string {
	value := c.Get(key)
	return convert.StringConverter.ToStringWithDefault(value, defaultValue)
}

func (c *StringValueMap) GetAsNullableBoolean(key string) *bool {
	value := c.Get(key)
	return convert.BooleanConverter.ToNullableBoolean(value)
}

func (c *StringValueMap) GetAsBoolean(key string) bool {
	return c.GetAsBooleanWithDefault(key, false)
}

func (c *StringValueMap) GetAsBooleanWithDefault(key string, defaultValue bool) bool {
	value := c.Get(key)
	return convert.BooleanConverter.ToBooleanWithDefault(value, defaultValue)
}

func (c *StringValueMap) GetAsNullableInteger(key string) *int {
	value := c.Get(key)
	return convert.IntegerConverter.ToNullableInteger(value)
}

func (c *StringValueMap) GetAsInteger(key string) int {
	return c.GetAsIntegerWithDefault(key, 0)
}

func (c *StringValueMap) GetAsIntegerWithDefault(key string, defaultValue int) int {
	value := c.Get(key)
	return convert.IntegerConverter.ToIntegerWithDefault(value, defaultValue)
}

func (c *StringValueMap) GetAsNullableLong(key string) *int64 {
	value := c.Get(key)
	return convert.LongConverter.ToNullableLong(value)
}

func (c *StringValueMap) GetAsLong(key string) int64 {
	return c.GetAsLongWithDefault(key, 0)
}

func (c *StringValueMap) GetAsLongWithDefault(key string, defaultValue int64) int64 {
	value := c.Get(key)
	return convert.LongConverter.ToLongWithDefault(value, defaultValue)
}

func (c *StringValueMap) GetAsNullableFloat(key string) *float32 {
	value := c.Get(key)
	return convert.FloatConverter.ToNullableFloat(value)
}

func (c *StringValueMap) GetAsFloat(key string) float32 {
	return c.GetAsFloatWithDefault(key, 0)
}

func (c *StringValueMap) GetAsFloatWithDefault(key string, defaultValue float32) float32 {
	value := c.Get(key)
	return convert.FloatConverter.ToFloatWithDefault(value, defaultValue)
}

func (c *StringValueMap) GetAsNullableDouble(key string) *float64 {
	value := c.Get(key)
	return convert.DoubleConverter.ToNullableDouble(value)
}

func (c *StringValueMap) GetAsDouble(key string) float64 {
	return c.GetAsDoubleWithDefault(key, 0)
}

func (c *StringValueMap) GetAsDoubleWithDefault(key string, defaultValue float64) float64 {
	value := c.Get(key)
	return convert.DoubleConverter.ToDoubleWithDefault(value, defaultValue)
}

func (c *StringValueMap) GetAsNullableDateTime(key string) *time.Time {
	value := c.Get(key)
	return convert.DateTimeConverter.ToNullableDateTime(value)
}

func (c *StringValueMap) GetAsDateTime(key string) time.Time {
	return c.GetAsDateTimeWithDefault(key, time.Time{})
}

func (c *StringValueMap) GetAsDateTimeWithDefault(key string, defaultValue time.Time) time.Time {
	value := c.Get(key)
	return convert.DateTimeConverter.ToDateTimeWithDefault(value, defaultValue)
}

// func (c *StringValueMap) GetAsNullableType(typ convert.TypeCode, key string) interface{} {
// 	value := c.Get(key)
// 	return convert.TypeConverter.ToNullableType(typ, value)
// }

// func (c *StringValueMap) GetAsType(typ convert.TypeCode, key string) interface{} {
// 	return c.GetAsTypeWithDefault(typ, key, nil)
// }

// func (c *StringValueMap) GetAsTypeWithDefault(typ convert.TypeCode, key string, defaultValue interface{}) interface{} {
// 	value := c.Get(key)
// 	return convert.TypeConverter.ToTypeWithDefault(typ, value, defaultValue)
// }

func (c *StringValueMap) GetAsValue(key string) *AnyValue {
	value := c.Get(key)
	return NewAnyValue(value)
}

func (c *StringValueMap) GetAsNullableArray(key string) *AnyValueArray {
	value := c.Get(key)
	if value != "" {
		return NewAnyValueArrayFromValue(value)
	} else {
		return nil
	}
}

func (c *StringValueMap) GetAsArray(key string) *AnyValueArray {
	value := c.Get(key)
	return NewAnyValueArrayFromValue(value)
}

func (c *StringValueMap) GetAsArrayWithDefault(key string, defaultValue *AnyValueArray) *AnyValueArray {
	result := c.GetAsNullableArray(key)
	if result != nil {
		return result
	} else {
		return defaultValue
	}
}

func (c *StringValueMap) GetAsNullableMap(key string) *AnyValueMap {
	value := c.Get(key)
	if value != "" {
		return NewAnyValueMapFromValue(value)
	} else {
		return nil
	}
}

func (c *StringValueMap) GetAsMap(key string) *AnyValueMap {
	value := c.Get(key)
	return NewAnyValueMapFromValue(value)
}

func (c *StringValueMap) GetAsMapWithDefault(key string, defaultValue *AnyValueMap) *AnyValueMap {
	result := c.GetAsNullableMap(key)
	if result != nil {
		return result
	} else {
		return defaultValue
	}
}

func (c *StringValueMap) String() string {
	builder := ""

	// Todo: User encoder
	for key := range c.value {
		value := c.value[key]

		if len(builder) > 0 {
			builder = builder + ";"
		}

		if value != "" {
			builder = builder + fmt.Sprintf("%s=%s", key, value)
		} else {
			builder = builder + key
		}
	}

	return builder
}

func (c *StringValueMap) Clone() interface{} {
	return NewStringValueMap(c.value)
}

func (c *StringValueMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *StringValueMap) UnmarshalJSON(data []byte) error {
	var values map[string]interface{}
	err := json.Unmarshal(data, &values)
	if err == nil {
		c.Clear()
		c.AppendAny(values)
	}
	return err
}

// Other static methods

func NewStringValueMapFromValue(value interface{}) *StringValueMap {
	result := NewEmptyStringValueMap()
	result.SetAsSingleObject(value)
	return result
}

func NewStringValueMapFromTuples(tuples ...interface{}) *StringValueMap {
	return NewStringValueMapFromTuplesArray(tuples)
}

func NewStringValueMapFromTuplesArray(tuples []interface{}) *StringValueMap {
	result := NewEmptyStringValueMap()
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

func NewStringValueMapFromString(line string) *StringValueMap {
	result := NewEmptyStringValueMap()
	if line == "" {
		return result
	}

	// Todo: User tokenizer / decoder
	tokens := strings.Split(line, ";")

	for index := 0; index < len(tokens); index++ {
		token := tokens[index]
		if len(token) == 0 {
			continue
		}

		pos := strings.Index(token, "=")

		var key string
		if pos > 0 {
			key = token[0:pos]
			key = strings.TrimSpace(key)
		} else {
			key = strings.TrimSpace(token)
		}

		var value string
		if pos > 0 {
			value = token[pos+1:]
			value = strings.TrimSpace(value)
		} else {
			value = ""
		}

		result.Put(key, value)
	}

	return result
}

func NewStringValueMapFromMaps(maps ...map[string]string) *StringValueMap {
	result := NewEmptyStringValueMap()
	if len(maps) > 0 {
		for index := 0; index < len(maps); index++ {
			result.Append(maps[index])
		}
	}
	return result
}
