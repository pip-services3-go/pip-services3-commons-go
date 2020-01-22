package data

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
)

/*
Cross-language implementation of a map (dictionary) where all keys and values are strings.
The stored values can be converted to different types using variety of accessor methods.
The string map is highly versatile. It can be converted into many formats, stored and sent over the wire.

This class is widely used in Pip.Services as a basis for variety of classes, such as ConfigParams, ConnectionParams,
CredentialParams and others.

Example
value1 := NewStringValueMapFromString("key1=1;key2=123.456;key3=2018-01-01");

value1.getAsBoolean("key1");   // Result: true
value1.getAsInteger("key2");   // Result: 123
value1.getAsFloat("key2");     // Result: 123.456
value1.getAsDateTime("key3");  // Result: new Date(2018,0,1)
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
*/

type StringValueMap struct {
	value map[string]string
}

// Creates a new instance of the map.
func NewEmptyStringValueMap() *StringValueMap {
	c := &StringValueMap{}
	c.value = map[string]string{}
	return c
}

// Creates a new instance of the map and assigns its value.
// Parameters
// 			- value map[string]string
// Returns *StringValueMap
func NewStringValueMap(value map[string]string) *StringValueMap {
	c := &StringValueMap{}
	c.value = map[string]string{}
	c.Append(value)
	return c
}

//Return inner values of map as interface{}
func (c *StringValueMap) InnerValue() interface{} {
	return c.value
}

//Returns map of elements as map[string]interface{}
func (c *StringValueMap) Value() map[string]string {
	return c.value
}

// Gets a map element specified by its key.
// Parameters:
// 			-key string
// 			a key of the element to get.
// Returns string
// the value of the map element.
func (c *StringValueMap) Get(key string) string {
	return c.value[key]
}

// Gets keys of all elements stored in this map.
// Returns []string
// a list with all map keys.
func (c *StringValueMap) Keys() []string {
	keys := []string{}
	for key := range c.value {
		keys = append(keys, key)
	}
	return keys
}

// Puts a new value into map element specified by its key.
// Parameters:
// 			- key string
// 			a key of the element to put.
// 			- value interface{}
// a new value for map element.
// Returns interface{}
func (c *StringValueMap) Put(key string, value interface{}) {
	c.value[key] = convert.StringConverter.ToString(value)
}

// Removes a map element specified by its key
// Parameters:
// 			- key string
// 			a key of the element to remove.
func (c *StringValueMap) Remove(key string) {
	delete(c.value, key)
}

// Checks if this map contains a key. The check uses direct comparison between key and the specified key value.
// Parameters
// 			- key string
// 			a value to be checked
// Returns bool
// true if this map contains the key or false otherwise.
func (c *StringValueMap) Contains(key string) bool {
	_, ok := c.value[key]
	return ok
}

// Appends new elements to this map.
// Parameters:
// 		- values map[string]string
// 		a map with elements to be added.
func (c *StringValueMap) Append(values map[string]string) {
	if values == nil {
		return
	}

	for key := range values {
		c.value[key] = values[key]
	}
}

// Appends new elements to this map.
// Parameters:
// 		- values map[string]interface{}
// 		a map with elements to be added.
func (c *StringValueMap) AppendAny(values map[string]interface{}) {
	if values == nil {
		return
	}

	for key := range values {
		c.value[key] = convert.StringConverter.ToString(values[key])
	}
}

// Clears this map by removing all its elements.
func (c *StringValueMap) Clear() {
	c.value = map[string]string{}
}

// Gets a number of elements stored in this map.
// Returns int
// the number of elements in this map.
func (c *StringValueMap) Len() int {
	return len(c.value)
}

// Gets the value stored in map element without any conversions.
// When element index is not defined it returns the entire array value.
// Returns interface{}
// the element value or value of the array when index is not defined.
func (c *StringValueMap) GetAsSingleObject() interface{} {
	return *c
}

// Sets a new value to map.
// Parameters:
// 			 - value interface{}
// a new element or array value.
func (c *StringValueMap) SetAsSingleObject(value interface{}) {
	a := convert.ToMap(value)
	//*c = a
	c.Clear()
	c.AppendAny(a)
}

// Gets the value stored in map element without any conversions.
// When element key is not defined it returns the entire map value.
// Parameters:
// 			 - key string
// 			 a key of the element to get
// Returns interface{}
// the element value or value of the map when index is not defined.
func (c *StringValueMap) GetAsObject(key string) interface{} {
	return c.Get(key)
}

// Sets a new value to map element specified by its index. When the index is not defined, it resets the entire map value.
// This method has double purpose because method overrides are not supported in JavaScript.
// see
// MapConverter.toMap
// Parameters:
// 			- key any
// 			 a key of the element to set
//  		- value interface{}
// a new element or map value.
func (c *StringValueMap) SetAsObject(key string, value interface{}) {
	c.Put(key, value)
}

// Converts map element into a string or returns null if conversion is not possible.
// see
// StringConverter.toNullableString
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns *string
// string value of the element or null if conversion is not supported.
func (c *StringValueMap) GetAsNullableString(key string) *string {
	value := c.Get(key)
	return convert.StringConverter.ToNullableString(value)
}

// Converts map element into a string or returns "" if conversion is not possible.
// see
// getAsStringWithDefault
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns string
// string value of the element or "" if conversion is not supported.
func (c *StringValueMap) GetAsString(key string) string {
	return c.GetAsStringWithDefault(key, "")
}

// Converts map element into a string or returns default value if conversion is not possible.
// see
// StringConverter.toStringWithDefault
// Parameters:
// 			 - key string
// 			 a key of element to get.
// 			 - defaultValue string
// 			 the default value
// Returns string
// string value of the element or default value if conversion is not supported.
func (c *StringValueMap) GetAsStringWithDefault(key string, defaultValue string) string {
	value := c.Get(key)

	// A special case for Golang strings
	if value == "" {
		return defaultValue
	}

	return convert.StringConverter.ToStringWithDefault(value, defaultValue)
}

// Converts map element into a boolean or returns null if conversion is not possible.
// see
// BooleanConverter.toNullableBoolean
// Parameters:
// 			 - key string
// 			 a key of element to get.
// Returns *bool
// boolean value of the element or null if conversion is not supported.
func (c *StringValueMap) GetAsNullableBoolean(key string) *bool {
	value := c.Get(key)
	return convert.BooleanConverter.ToNullableBoolean(value)
}

// Converts map element into a boolean or returns false if conversion is not possible.
// see
// getAsBooleanWithDefault
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns bool
// boolean value of the element or false if conversion is not supported.
func (c *StringValueMap) GetAsBoolean(key string) bool {
	return c.GetAsBooleanWithDefault(key, false)
}

// Converts map element into a boolean or returns default value if conversion is not possible.
// see
// BooleanConverter.toBooleanWithDefault
// Parameters
// 			- key string
// 			a key of element to get.
// 			- defaultValue bool
// 			the default value
// Returns bool
// boolean value of the element or default value if conversion is not supported.
func (c *StringValueMap) GetAsBooleanWithDefault(key string, defaultValue bool) bool {
	value := c.Get(key)
	return convert.BooleanConverter.ToBooleanWithDefault(value, defaultValue)
}

// Converts map element into an integer or returns null if conversion is not possible.
// see
// IntegerConverter.toNullableInteger
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns *int
// integer value of the element or null if conversion is not supported.
func (c *StringValueMap) GetAsNullableInteger(key string) *int {
	value := c.Get(key)
	return convert.IntegerConverter.ToNullableInteger(value)
}

// Converts map element into an integer or returns 0 if conversion is not possible.
// see
// getAsIntegerWithDefault
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns int
func (c *StringValueMap) GetAsInteger(key string) int {
	return c.GetAsIntegerWithDefault(key, 0)
}

// Converts map element into an integer or returns default value if conversion is not possible.
// see
// IntegerConverter.toIntegerWithDefault
// Parameters:
// 			- key string
// 			a key of element to get.
// 			- defaultValue int
// 			the default value
// Returns int
// integer value of the element or default value if conversion is not supported.
func (c *StringValueMap) GetAsIntegerWithDefault(key string, defaultValue int) int {
	value := c.Get(key)
	return convert.IntegerConverter.ToIntegerWithDefault(value, defaultValue)
}

// Converts map element into a int64 or returns null if conversion is not possible.
// see
// LongConverter.toNullableLong
// Parameters:
// 			 - key string
// 			a key of element to get.
// Returns *int64
// int64 value of the element or null if conversion is not supported.
func (c *StringValueMap) GetAsNullableLong(key string) *int64 {
	value := c.Get(key)
	return convert.LongConverter.ToNullableLong(value)
}

// Converts map element into a int64 or returns 0 if conversion is not possible.
// see
// getAsLongWithDefault
// Parameters:
// 			 - key string
// 			a key of element to get.
// Returns int64
// loint64ng value of the element or 0 if conversion is not supported.
func (c *StringValueMap) GetAsLong(key string) int64 {
	return c.GetAsLongWithDefault(key, 0)
}

// Converts map element into a int64 or returns default value if conversion is not possible.
// see
// LongConverter.toLongWithDefault
// Parameters:
// 			- key string
// 			a key of element to get.
// 			- defaultValue int64
// 			the default value
// Returns int64
// int64 value of the element or default value if conversion is not supported.
func (c *StringValueMap) GetAsLongWithDefault(key string, defaultValue int64) int64 {
	value := c.Get(key)
	return convert.LongConverter.ToLongWithDefault(value, defaultValue)
}

// Converts map element into a float32 or returns null if conversion is not possible.
// see
// FloatConverter.toNullableFloat
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns *float32
// float32 value of the element or null if conversion is not supported.
func (c *StringValueMap) GetAsNullableFloat(key string) *float32 {
	value := c.Get(key)
	return convert.FloatConverter.ToNullableFloat(value)
}

// Converts map element into a float32 or returns 0 if conversion is not possible.
// see
// getAsFloatWithDefault
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns float32
// float32 value of the element or 0 if conversion is not supported.
func (c *StringValueMap) GetAsFloat(key string) float32 {
	return c.GetAsFloatWithDefault(key, 0)
}

// Converts map element into a float32 or returns default value if conversion is not possible.
// see
// FloatConverter.toFloatWithDefault
// Parameters:
// 			- key string
// 			a key of element to get.
// 			- defaultValue: float32
// 			the default value
// Returns float32
// float32 value of the element or default value if conversion is not supported.
func (c *StringValueMap) GetAsFloatWithDefault(key string, defaultValue float32) float32 {
	value := c.Get(key)
	return convert.FloatConverter.ToFloatWithDefault(value, defaultValue)
}

// Converts map element into a float64 or returns null if conversion is not possible.
// see
// DoubleConverter.toNullableDouble
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns *float64
// float64 value of the element or null if conversion is not supported.
func (c *StringValueMap) GetAsNullableDouble(key string) *float64 {
	value := c.Get(key)
	return convert.DoubleConverter.ToNullableDouble(value)
}

// Converts map element into a float64 or returns 0 if conversion is not possible.
// see
// getAsDoubleWithDefault
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns float64
// float64 value of the element or 0 if conversion is not supported.
func (c *StringValueMap) GetAsDouble(key string) float64 {
	return c.GetAsDoubleWithDefault(key, 0)
}

// Converts map element into a float64 or returns default value if conversion is not possible.
// see
// DoubleConverter.toDoubleWithDefault
// Parameters:
// 			- key string
// 			a key of element to get.
// 			- defaultValue float64
// 			the default value
// Returns float64
// float64 value of the element or default value if conversion is not supported.
func (c *StringValueMap) GetAsDoubleWithDefault(key string, defaultValue float64) float64 {
	value := c.Get(key)
	return convert.DoubleConverter.ToDoubleWithDefault(value, defaultValue)
}

// Converts map element into a time.Time or returns null if conversion is not possible.
// see
// DateTimeConverter.toNullableDateTime
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns time.Time
// time.Time value of the element or null if conversion is not supported.
func (c *StringValueMap) GetAsNullableDateTime(key string) *time.Time {
	value := c.Get(key)
	return convert.DateTimeConverter.ToNullableDateTime(value)
}

// Converts map element into a time.Time or returns the current date if conversion is not possible.
// see
// getAsDateTimeWithDefault
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns time.Time
// time.Time value of the element or the current date if conversion is not supported.
func (c *StringValueMap) GetAsDateTime(key string) time.Time {
	return c.GetAsDateTimeWithDefault(key, time.Time{})
}

// Converts map element into a time.Time or returns default value if conversion is not possible.
// see
// DateTimeConverter.toDateTimeWithDefault
// Parameters:
// 			- key string
// 			a key of element to get.
// 			- defaultValue time.Time
// 			the default value
// Returns time.Time
// time.Time value of the element or default value if conversion is not supported.

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

// Converts map element into an AnyValue or returns an empty AnyValue if conversion is not possible.
// see
// AnyValue
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns *AnyValue
// AnyValue value of the element or empty AnyValue if conversion is not supported.
func (c *StringValueMap) GetAsValue(key string) *AnyValue {
	value := c.Get(key)
	return NewAnyValue(value)
}

// Converts map element into an AnyValueArray or returns null if conversion is not possible.
// see
// AnyValueArray
// see
// AnyValueArray.fromValue
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns *AnyValueArray
// AnyValueArray value of the element or null if conversion is not supported.
func (c *StringValueMap) GetAsNullableArray(key string) *AnyValueArray {
	value := c.Get(key)
	if value != "" {
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
// 			- key string
// 			a key of element to get.
// Returns *AnyValueArray
// AnyValueArray value of the element or empty AnyValueArray if conversion is not supported.
func (c *StringValueMap) GetAsArray(key string) *AnyValueArray {
	value := c.Get(key)
	return NewAnyValueArrayFromValue(value)
}

// Converts map element into an AnyValueArray or returns default value if conversion is not possible.
// see
// AnyValueArray
// see
// getAsNullableArray
// Parameters
// 			- key string
// 			a key of element to get.
// 			- defaultValue *AnyValueArray
// 			the default value
// Returns *AnyValueArray
// AnyValueArray value of the element or default value if conversion is not supported.
func (c *StringValueMap) GetAsArrayWithDefault(key string, defaultValue *AnyValueArray) *AnyValueArray {
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
// 			 - key string
// 			a key of element to get.
// Returns *AnyValueMap
// AnyValueMap value of the element or null if conversion is not supported.
func (c *StringValueMap) GetAsNullableMap(key string) *AnyValueMap {
	value := c.Get(key)
	if value != "" {
		return NewAnyValueMapFromValue(value)
	} else {
		return nil
	}
}

// Converts map element into an AnyValueMap or returns empty AnyValueMap if conversion is not possible.
// see
// fromValue
// Parameters:
// 			- key string
// 			a key of element to get.
// Returns *AnyValueMap
// AnyValueMap value of the element or empty AnyValueMap if conversion is not supported.
func (c *StringValueMap) GetAsMap(key string) *AnyValueMap {
	value := c.Get(key)
	return NewAnyValueMapFromValue(value)
}

// Converts map element into an AnyValueMap or returns default value if conversion is not possible.
// see
// getAsNullableMap
// Parameters:
// 			 - key string
// 			a key of element to get.
// 			defaultValue *AnyValueMap
// 			the default value
// Returns *AnyValueMap
// AnyValueMap value of the element or default value if conversion is not supported.
func (c *StringValueMap) GetAsMapWithDefault(key string, defaultValue *AnyValueMap) *AnyValueMap {
	result := c.GetAsNullableMap(key)
	if result != nil {
		return result
	} else {
		return defaultValue
	}
}

// Gets a string representation of the object. The result is a semicolon-separated
// list of key-value pairs as "key1=value1;key2=value2;key=value3"
// Returns string
// a string representation of the object.
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

// Creates a binary clone of this object.
// Returns interface{}
// a clone of this object.
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

// Converts specified value into StringValueMap.
// see
// setAsObject
// Parameters:
// 			 - value interface{}
// 			 value to be converted
// Returns *StringValueMap
// a newly created StringValueMap.
func NewStringValueMapFromValue(value interface{}) *StringValueMap {
	result := NewEmptyStringValueMap()
	result.SetAsSingleObject(value)
	return result
}

// Creates a new StringValueMap from a list of key-value pairs called tuples.
// see
// fromTuplesArray
// Parameters
// 			- tuples ...interface{}
// 			a list of values where odd elements are keys and the following even elements are values
// Returns *StringValueMap
// a newly created StringValueMap.

func NewStringValueMapFromTuples(tuples ...interface{}) *StringValueMap {
	return NewStringValueMapFromTuplesArray(tuples)
}

// Creates a new StringValueMap from a list of key-value pairs called tuples.
// The method is similar to fromTuples but tuples are passed as array instead of parameters.
// Parameters:
// 			 - tuples: []interface{}
// 			a list of values where odd elements are keys and the following even elements are values
// Returns *StringValueMap
// a newly created StringValueMap.
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

// Parses semicolon-separated key-value pairs and returns them as a StringValueMap.
// Parameters:
// 			- line string
// 			semicolon-separated key-value list to initialize StringValueMap.
// Returns *StringValueMap
// a newly created StringValueMap.
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

// Creates a new AnyValueMap by merging two or more maps.
// Maps defined later in the list override values from previously defined maps.
// Parameters:
//  		- maps...map[string]string
// 			an array of maps to be merged
// Returns StringValueMap
// a newly created AnyValueMap.
func NewStringValueMapFromMaps(maps ...map[string]string) *StringValueMap {
	result := NewEmptyStringValueMap()
	if len(maps) > 0 {
		for index := 0; index < len(maps); index++ {
			result.Append(maps[index])
		}
	}
	return result
}
