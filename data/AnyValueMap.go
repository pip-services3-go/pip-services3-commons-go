package data

import (
	"fmt"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
)

/*
Cross-language implementation of dynamic object map (dictionary) what can hold values of any type. The stored values can be converted to different types using variety of accessor methods.

Example
 value1 := AnyValueMap({ key1: 1, key2: "123.456", key3: "2018-01-01" });

 value1.GetAsBoolean("key1");   // Result: true
 value1.GetAsInteger("key2");   // Result: 123
 value1.getAsFloat("key2");     // Result: 123.456
 value1.GetAsDateTime("key3");  // Result: new Date(2018,0,1)
see
StringConverter
see
TypeConverter
see
BooleanConverter
see
IntegerConverter
see
LongConverter
see
DoubleConverter
see
FloatConverter
see
DateTimeConverter
see
ICloneable
*/
type AnyValueMap struct {
	value map[string]interface{}
	base  IMap
}

// Creates a new empty instance of the map.
// Returns *AnyValueMap
func NewEmptyAnyValueMap() *AnyValueMap {
	c := &AnyValueMap{
		value: map[string]interface{}{},
	}
	c.base = c
	return c
}

// Creates a new instance of the map and assigns base methods from interface.
// Parameters:
//  - base IMap
// Returns *AnyValueMap
func InheritAnyValueMap(base IMap) *AnyValueMap {
	c := &AnyValueMap{
		value: map[string]interface{}{},
	}
	c.base = base
	return c
}

// Creates a new instance of the map and assigns its value.
// Parameters:
//  - values map[string]interface{}
// Returns *AnyValueMap
func NewAnyValueMap(value map[string]interface{}) *AnyValueMap {
	c := &AnyValueMap{
		value: map[string]interface{}{},
	}
	c.base = c
	c.Append(value)
	return c
}

//Return inner values of map as interface{}
func (c *AnyValueMap) InnerValue() interface{} {
	return c.value
}

//Returns map of elements as map[string]interface{}
func (c *AnyValueMap) Value() map[string]interface{} {
	return c.value
}

// Gets a map element specified by its key.
// Parameters:
//  - key string
//  a key of the element to get.
// Returns interface {}
// the value of the map element.
func (c *AnyValueMap) Get(key string) interface{} {
	return c.value[key]
}

// Gets keys of all elements stored in this map.
// Returns []string
// a list with all map keys.
func (c *AnyValueMap) Keys() []string {
	keys := []string{}
	for key := range c.value {
		keys = append(keys, key)
	}
	return keys
}

// Puts a new value into map element specified by its key.
// Parameters:
//  - key string
//  a key of the element to put.
//  - value interface{}
//  a new value for map element.
// Returns interface{}
func (c *AnyValueMap) Put(key string, value interface{}) {
	c.value[key] = value
}

// Removes a map element specified by its key
// Parameters:
//  - key string
//  a key of the element to remove.
func (c *AnyValueMap) Remove(key string) {
	delete(c.value, key)
}

// Checks if this map contains a key. The check uses direct comparison between key and the specified key value.
// Parameters
//  - key string
//  a value to be checked
// Returns bool
// true if this map contains the key or false otherwise.
func (c *AnyValueMap) Contains(key string) bool {
	_, ok := c.value[key]
	return ok
}

// Appends new elements to this map.
// Parameters:
//  value: map[string]interface{}
// a map of elements to be added.
func (c *AnyValueMap) Append(value map[string]interface{}) {
	if value == nil {
		return
	}

	for key := range value {
		c.value[key] = value[key]
	}
}

//Clears this map by removing all its elements.
func (c *AnyValueMap) Clear() {
	c.value = map[string]interface{}{}
}

//Gets a number of elements stored in this map.
// Returns int
// the number of elements in this map.
func (c *AnyValueMap) Len() int {
	return len(c.value)
}

// Gets the value stored in map element without any conversions.
// When element index is not defined it returns the entire array value.
// Returns interface{}
// the element value or value of the array when index is not defined.
func (c *AnyValueMap) GetAsSingleObject() interface{} {
	return c.value
}

// Sets a new value to map.
// see
// ArrayConverter.toMap
// Parameters:
//  - value interface{}
// a new element or array value.
func (c *AnyValueMap) SetAsSingleObject(value interface{}) {
	a := convert.ToMap(value)
	c.value = a
}

// Gets the value stored in map element without any conversions.
// When element key is not defined it returns the entire map value.
// Parameters:
//  - key string
//  a key of the element to get
// Returns interface{}
// the element value or value of the map when index is not defined.
func (c *AnyValueMap) GetAsObject(key string) interface{} {
	return c.base.Get(key)
}

// Sets a new value to map element specified by its index. When the index is not defined, it resets the entire map value.
// see
// MapConverter.toMap
// Parameters:
//  - key string
//   a key of the element to set
//  - value interface{}
// a new element or map value.

func (c *AnyValueMap) SetAsObject(key string, value interface{}) {
	c.base.Put(key, value)
}

// Converts map element into a string or returns null if conversion is not possible.
// see
// StringConverter.toNullableString
// Parameters:
//  - key string
//  a key of element to get.
// Returns string
// string value of the element or null if conversion is not supported.
func (c *AnyValueMap) GetAsNullableString(key string) *string {
	value := c.base.Get(key)
	return convert.StringConverter.ToNullableString(value)
}

// Converts map element into a string or returns "" if conversion is not possible.
// see
// getAsStringWithDefault
// Parameters:
//  - key string
//  a key of element to get.
// Returns string
// string value of the element or "" if conversion is not supported.
func (c *AnyValueMap) GetAsString(key string) string {
	return c.GetAsStringWithDefault(key, "")
}

// Converts map element into a string or returns default value if conversion is not possible.
// see
// StringConverter.toStringWithDefault
// Parameters:
//  - key string
//  a key of element to get.
//  - defaultValue string
//  the default value
// Returns string
// string value of the element or default value if conversion is not supported.
func (c *AnyValueMap) GetAsStringWithDefault(key string, defaultValue string) string {
	value := c.base.Get(key)
	return convert.StringConverter.ToStringWithDefault(value, defaultValue)
}

// Converts map element into a boolean or returns null if conversion is not possible.
// see
// BooleanConverter.toNullableBoolean
// Parameters:
//  - key string
//  a key of element to get.
// Returns bool
// bool value of the element or null if conversion is not supported.
func (c *AnyValueMap) GetAsNullableBoolean(key string) *bool {
	value := c.base.Get(key)
	return convert.BooleanConverter.ToNullableBoolean(value)
}

// Converts map element into a boolean or returns false if conversion is not possible.
// see
// getAsBooleanWithDefault
// Parameters:
//  - key: string
//  a key of element to get.
// Returns bool
// bool value of the element or false if conversion is not supported.
func (c *AnyValueMap) GetAsBoolean(key string) bool {
	return c.GetAsBooleanWithDefault(key, false)
}

// Converts map element into a boolean or returns default value if conversion is not possible.
// see
// BooleanConverter.toBooleanWithDefault
// Parameters:
//  - key string
//  a key of element to get.
//  - defaultValue  bool
//  the default value
// Returns bool
// bool value of the element or default value if conversion is not supported.
func (c *AnyValueMap) GetAsBooleanWithDefault(key string, defaultValue bool) bool {
	value := c.base.Get(key)
	return convert.BooleanConverter.ToBooleanWithDefault(value, defaultValue)
}

// Converts map element into an integer or returns null if conversion is not possible.
// see
// IntegerConverter.toNullableInteger
// Parameters:
//  - key string
//  a key of element to get.
// Returns *int
// integer value of the element or null if conversion is not supported.
func (c *AnyValueMap) GetAsNullableInteger(key string) *int {
	value := c.base.Get(key)
	return convert.IntegerConverter.ToNullableInteger(value)
}

// Converts map element into an integer or returns 0 if conversion is not possible.
// see
// getAsIntegerWithDefault
// Parameters:
//  - key string
//  a key of element to get.
// Returns int
// integer value of the element or 0 if conversion is not supported.
func (c *AnyValueMap) GetAsInteger(key string) int {
	return c.GetAsIntegerWithDefault(key, 0)
}

// Converts map element into an integer or returns default value if conversion is not possible.
// see
// IntegerConverter.toIntegerWithDefault
// Parameters:
//  - key string
//  a key of element to get.
//  -defaultValue int
//  the default value
// Returns int
// integer value of the element or default value if conversion is not supported.
func (c *AnyValueMap) GetAsIntegerWithDefault(key string, defaultValue int) int {
	value := c.base.Get(key)
	return convert.IntegerConverter.ToIntegerWithDefault(value, defaultValue)
}

// Converts map element into a long or returns null if conversion is not possible.
// see
// LongConverter.toNullableLong
// Parameters:
//  - key string
//  a key of element to get.
// Returns *int64
// int64 value of the element or null if conversion is not supported.
func (c *AnyValueMap) GetAsNullableLong(key string) *int64 {
	value := c.base.Get(key)
	return convert.LongConverter.ToNullableLong(value)
}

// Converts map element into a long or returns 0 if conversion is not possible.
// see
// getAsLongWithDefault
// Parameters:
//  - key string
//  a key of element to get.
// Returns int64
// int64 value of the element or 0 if conversion is not supported.
func (c *AnyValueMap) GetAsLong(key string) int64 {
	return c.GetAsLongWithDefault(key, 0)
}

// Converts map element into a long or returns default value if conversion is not possible.
// see
// LongConverter.toLongWithDefault
// Parameters:
//  -key string
//  a key of element to get.
//  - defaultValue int64
//  the default value
// Returns int64
// int64 value of the element or default value if conversion is not supported.
func (c *AnyValueMap) GetAsLongWithDefault(key string, defaultValue int64) int64 {
	value := c.base.Get(key)
	return convert.LongConverter.ToLongWithDefault(value, defaultValue)
}

// Converts map element into a float or returns null if conversion is not possible.
// see
// FloatConverter.toNullableFloat
// Parameters:
//  - key string
//  a key of element to get.
// Returns *float32
// *float32 value of the element or null if conversion is not supported.

func (c *AnyValueMap) GetAsNullableFloat(key string) *float32 {
	value := c.base.Get(key)
	return convert.FloatConverter.ToNullableFloat(value)
}

// Converts map element into a float or returns 0 if conversion is not possible.
// see
// getAsFloatWithDefault
// Parameters:
//  - key string
//  a key of element to get.
// Returns float32
// float32 value of the element or 0 if conversion is not supported.
func (c *AnyValueMap) GetAsFloat(key string) float32 {
	return c.GetAsFloatWithDefault(key, 0)
}

// Converts map element into a flot or returns default value if conversion is not possible.
// see
// FloatConverter.toFloatWithDefault
// Parameters:
//  - key string
//  a key of element to get.
//  - defaultValue float32
//  the default value
// Returns float32
// float32 value of the element or default value if conversion is not supported.
func (c *AnyValueMap) GetAsFloatWithDefault(key string, defaultValue float32) float32 {
	value := c.base.Get(key)
	return convert.FloatConverter.ToFloatWithDefault(value, defaultValue)
}

// Converts map element into a double or returns null if conversion is not possible.
// see
// DoubleConverter.toNullableDouble
// Parameters:
//  - key string
//  a key of element to get.
// Returns float64
// float64 value of the element or null if conversion is not supported.

func (c *AnyValueMap) GetAsNullableDouble(key string) *float64 {
	value := c.base.Get(key)
	return convert.DoubleConverter.ToNullableDouble(value)
}

// Converts map element into a double or returns 0 if conversion is not possible.
// see
// getAsDoubleWithDefault
// Parameters:
//  - key string
//  a key of element to get.
// Returns float64
// float64 value of the element or 0 if conversion is not supported.
func (c *AnyValueMap) GetAsDouble(key string) float64 {
	return c.GetAsDoubleWithDefault(key, 0)
}

// Converts map element into a double or returns default value if conversion is not possible.
// see
// DoubleConverter.toDoubleWithDefault
// Parameters:
//  - key string
//  a key of element to get.
//  - defaultValue float64
//  the default value
// Returns float64
// float64 value of the element or default value if conversion is not supported.
func (c *AnyValueMap) GetAsDoubleWithDefault(key string, defaultValue float64) float64 {
	value := c.base.Get(key)
	return convert.DoubleConverter.ToDoubleWithDefault(value, defaultValue)
}

// Converts map element into a time.Time or returns null if conversion is not possible.
// see
// DateTimeConverter.toNullableDateTime
// Parameters:
//  - key string
//  a key of element to get.
// Returns *time.Time
// time.Time value of the element or null if conversion is not supported.

func (c *AnyValueMap) GetAsNullableDateTime(key string) *time.Time {
	value := c.base.Get(key)
	return convert.DateTimeConverter.ToNullableDateTime(value)
}

// Converts map element into a time.Time or returns the current date if conversion is not possible.
// see
// getAsDateTimeWithDefault
// Parameters:
//  - key string
//  a key of element to get.
// Returns time.Time
// time.Time value of the element or the current date if conversion is not supported.
func (c *AnyValueMap) GetAsDateTime(key string) time.Time {
	return c.GetAsDateTimeWithDefault(key, time.Time{})
}

// Converts map element into a time.Time or returns default value if conversion is not possible.
// see
// DateTimeConverter.toDateTimeWithDefault
// Parameters:
//  - key: string
//  a key of element to get.
//  - defaultValue: Date
//  the default value
// Returns time.Time
// time.Time value of the element or default value if conversion is not supported.
func (c *AnyValueMap) GetAsDateTimeWithDefault(key string, defaultValue time.Time) time.Time {
	value := c.base.Get(key)
	return convert.DateTimeConverter.ToDateTimeWithDefault(value, defaultValue)
}

// Converts map element into a time.Duration or returns null if conversion is not possible.
// see
// DateTimeConverter.toNullableDateTime
// Parameters:
//  - key string
//  a key of element to get.
// Returns *time.Duration
// time.Duration value of the element or null if conversion is not supported.
func (c *AnyValueMap) GetAsNullableDuration(key string) *time.Duration {
	value := c.base.Get(key)
	return convert.DurationConverter.ToNullableDuration(value)
}

// Converts map element into a time.Duration or returns the current date if conversion is not possible.
// see
// getAsDateTimeWithDefault
// Parameters:
//  - key string
//  a key of element to get.
// Returns time.Duration
// time.Duration value of the element or the current date if conversion is not supported.
func (c *AnyValueMap) GetAsDuration(key string) time.Duration {
	return c.GetAsDurationWithDefault(key, 0*time.Millisecond)
}

// Converts map element into a time.Duration or returns default value if conversion is not possible.
// see
// DateTimeConverter.toDateTimeWithDefault
// Parameters:
//  - key: string
//  a key of element to get.
//  - defaultValue: Date
//  the default value
// Returns time.Duration
// time.Duration value of the element or default value if conversion is not supported.
func (c *AnyValueMap) GetAsDurationWithDefault(key string, defaultValue time.Duration) time.Duration {
	value := c.base.Get(key)
	return convert.DurationConverter.ToDurationWithDefault(value, defaultValue)
}

// Converts map element into a value defined by specied typecode. If conversion is not possible it returns null.
// see
// TypeConverter.toNullableType
// Parameters:
//  - type TypeCode
//  the TypeCode that defined the type of the result
//  - key string
//  a key of element to get.
// Returns interface{}
// element value defined by the typecode or null if conversion is not supported.
func (c *AnyValueMap) GetAsNullableType(typ convert.TypeCode, key string) interface{} {
	value := c.base.Get(key)
	return convert.TypeConverter.ToNullableType(typ, value)
}

// Converts map element into a value defined by specied typecode.
// If conversion is not possible it returns default value for the specified type.
// see
// getAsTypeWithDefault
// Parameters:
//  - type TypeCode
//  the TypeCode that defined the type of the result
//  - key string
//  a key of element to get.
// Returns interface{}
// element value defined by the typecode or default if conversion is not supported.
func (c *AnyValueMap) GetAsType(typ convert.TypeCode, key string) interface{} {
	return c.GetAsTypeWithDefault(typ, key, nil)
}

// Converts map element into a value defined by specied typecode. If conversion is not possible it returns default value.
// see
// TypeConverter.toTypeWithDefault
// Parameters:
//  - type TypeCode
//  the TypeCode that defined the type of the result
//  - key string
//  a key of element to get.
//  defaultValue interface{}
//  the default value
// Returns interface {}
// element value defined by the typecode or default value if conversion is not supported.
func (c *AnyValueMap) GetAsTypeWithDefault(typ convert.TypeCode, key string, defaultValue interface{}) interface{} {
	value := c.base.Get(key)
	return convert.TypeConverter.ToTypeWithDefault(typ, value, defaultValue)
}

// Converts map element into an AnyValue or returns an empty AnyValue if conversion is not possible.
// see
// AnyValue
// see
// AnyValue.constructor
// Parameters:
//  - key string
//  a key of element to get.
// Returns *AnyValue
// AnyValue value of the element or empty AnyValue if conversion is not supported.
func (c *AnyValueMap) GetAsValue(key string) *AnyValue {
	value := c.base.Get(key)
	return NewAnyValue(value)
}

// Converts map element into an AnyValueArray or returns null if conversion is not possible.
// see
// AnyValueArray
// see
// AnyValueArray.fromValue
// Parameters:
//  - key string
//  a key of element to get.
// Returns *AnyValueArray
// AnyValueArray value of the element or null if conversion is not supported.
func (c *AnyValueMap) GetAsNullableArray(key string) *AnyValueArray {
	value := c.base.Get(key)
	if value != nil {
		return NewAnyValueArrayFromValue(value)
	} else {
		return nil
	}
}

// Converts map element into an AnyValueArray or returns empty AnyValueArray if conversion is not possible.
// see
// AnyValueArray
// see
// AnyValueArray.fromValue
// Parameters:
//  key string
//  a key of element to get.
// Returns *AnyValueArray
// AnyValueArray value of the element or empty AnyValueArray if conversion is not supported.
func (c *AnyValueMap) GetAsArray(key string) *AnyValueArray {
	value := c.base.Get(key)
	return NewAnyValueArrayFromValue(value)
}

// Converts map element into an AnyValueArray or returns default value if conversion is not possible.
// see
// AnyValueArray
// see
// getAsNullableArray
// Parameters:
//  - key string
//  a key of element to get.
//  - defaultValue: *AnyValueArray
//  the default value
// Returns *AnyValueArray
// AnyValueArray value of the element or default value if conversion is not supported.
func (c *AnyValueMap) GetAsArrayWithDefault(key string, defaultValue *AnyValueArray) *AnyValueArray {
	result := c.GetAsNullableArray(key)
	if result != nil {
		return result
	} else {
		return defaultValue
	}
}

// Converts map element into an AnyValueMap or returns null if conversion is not possible.
// see
// fromValue
// Parameters:
//  - key string
//  a key of element to get.
// Returns *AnyValueMap
// *AnyValueMap value of the element or null if conversion is not supported.
func (c *AnyValueMap) GetAsNullableMap(key string) *AnyValueMap {
	value := c.base.Get(key)
	if value != nil {
		return NewAnyValueMapFromValue(value)
	} else {
		return nil
	}
}

// Converts map element into an AnyValueMap or returns empty AnyValueMap if conversion is not possible.
// see
// fromValue
// Parameters:
//  - key string
//  a key of element to get.
// Returns *AnyValueMap
// AnyValueMap value of the element or empty AnyValueMap if conversion is not supported.
func (c *AnyValueMap) GetAsMap(key string) *AnyValueMap {
	value := c.base.Get(key)
	return NewAnyValueMapFromValue(value)
}

// Converts map element into an AnyValueMap or returns default value if conversion is not possible.
// see
// getAsNullableMap
// Parameters:
//  - key string
//  a key of element to get.
//  - defaultValue *AnyValueMap
//  the default value
// Returns *AnyValueMap
// AnyValueMap value of the element or default value if conversion is not supported.
func (c *AnyValueMap) GetAsMapWithDefault(key string, defaultValue *AnyValueMap) *AnyValueMap {
	result := c.GetAsNullableMap(key)
	if result != nil {
		return result
	} else {
		return defaultValue
	}
}

// Gets a string representation of the object.
// The result is a semicolon-separated list of key-value pairs as "key1=value1;key2=value2;key=value3"
// Returns string
// a string representation of the object.
func (c *AnyValueMap) String() string {
	builder := ""

	// Todo: User encoder
	for key := range c.Value() {
		value := c.base.Get(key)

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

// Creates a binary clone of this object.
// Returns any
// a clone of this object.
func (c *AnyValueMap) Clone() interface{} {
	return NewAnyValueMap(c.value)
}

// Converts specified value into AnyValueMap.
// see
// setAsObject
// Parameters:
//  - value interface {}
//  value to be converted
// Returns *AnyValueMap
// a newly created AnyValueMap.
func NewAnyValueMapFromValue(value interface{}) *AnyValueMap {
	result := NewEmptyAnyValueMap()
	result.SetAsSingleObject(value)
	return result
}

// Creates a new AnyValueMap from a list of key-value pairs called tuples.
// see
// fromTuplesArray
// Parameters:
//  - tuples ...tuples: interface{}
//  a list of values where odd elements are keys and the following even elements are values
// Returns *AnyValueMap
// a newly created AnyValueArray.
func NewAnyValueMapFromTuples(tuples ...interface{}) *AnyValueMap {
	return NewAnyValueMapFromTuplesArray(tuples)
}

// Creates a new AnyValueMap from a list of key-value pairs called tuples.
//The method is similar to fromTuples but tuples are passed as array instead of parameters.
// Parameters:
//  - tuples: []interface{}
//  a list of values where odd elements are keys and the following even elements are values
// Returns *AnyValueMap
// a newly created AnyValueArray.

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

// Creates a new AnyValueMap by merging two or more maps. Maps defined later in the list override values from previously defined maps.
// Parameters:
//  maps ...maps: any[]
//  an array of maps to be merged
// Returns *AnyValueMap
// a newly created AnyValueMap.
func NewAnyValueMapFromMaps(maps ...map[string]interface{}) *AnyValueMap {
	result := NewEmptyAnyValueMap()
	if len(maps) > 0 {
		for index := 0; index < len(maps); index++ {
			result.Append(maps[index])
		}
	}
	return result
}
