package data

import (
	"strings"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
)

/*
Cross-language implementation of dynamic object array what can hold values of any type. The stored values can be converted to different types using variety of accessor methods.

Example
value1 :=  NewAnyValueArray([1, "123.456", "2018-01-01"]);

value1.GetAsBoolean(0);   // Result: true
value1.GetAsInteger(1);   // Result: 123
value1.GetAsFloat(1);     // Result: 123.456
value1.GetAsDateTime(2);  // Result: new Date(2018,0,1)

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
type AnyValueArray struct {
	value []interface{}
}

// Creates a new instance of the empty array.
// Returns *AnyValueArray
func NewEmptyAnyValueArray() *AnyValueArray {
	return &AnyValueArray{
		value: make([]interface{}, 0, 10),
	}
}

// Creates a new instance of the array and assigns its value.
// Parameters:
// 			- values []interface{}
// Returns *AnyValueArray
func NewAnyValueArray(values []interface{}) *AnyValueArray {
	c := &AnyValueArray{
		value: make([]interface{}, len(values)),
	}
	copy(c.value, values)
	return c
}

//Return inner value of array as interface{}
func (c *AnyValueArray) InnerValue() interface{} {
	return c.value
}

//Returns array of elements []interface{}
func (c *AnyValueArray) Value() []interface{} {
	return c.value
}

// Returns length of array
func (c *AnyValueArray) Len() int {
	return len(c.value)
}

// Gets an array element specified by its index.
// Parameters:
//  		- index int
// 			an index of the element to get.
// Returns interface {}
// the value of the array element.
func (c *AnyValueArray) Get(index int) interface{} {
	return c.value[index]
}

// Puts a new value into array element specified by its index.
// Parameters:
// 			- index int
// 			an index of the element to put.
// 			value: interface {}
// 			a new value for array element.

func (c *AnyValueArray) Put(index int, value interface{}) {
	if cap(c.value)+1 < index {
		a := make([]interface{}, index+1, (index+1)*2)
		copy(a, c.value)
		c.value = a
	}

	c.value[index] = value
}

// Removes an array element specified by its index
// Parameters:
// 			- index int
// 			an index of the element to remove.
func (c *AnyValueArray) Remove(index int) {
	c.value = append(c.value[:index], c.value[index+1:]...)
}

// Push element in the end of array
// Parameters:
//          - value interface{}
//			an value what need to insert
func (c *AnyValueArray) Push(value interface{}) {
	c.value = append(c.value, value)
}

// Appends new elements to this array.
// Parameters:
// 			- elements []interface{}
// a list of elements to be added.
func (c *AnyValueArray) Append(elements []interface{}) {
	if elements != nil {
		c.value = append(c.value, elements...)
	}
}

//Clears this array by removing all its elements.
func (c *AnyValueArray) Clear() {
	c.value = make([]interface{}, 0, 10)
}

// Inflate AnyValueArray as single object
// Return interface{}
func (c *AnyValueArray) GetAsSingleObject() interface{} {
	return *c
}

// Set AnyValueArray from input object
// Parameters:
// 			- value interface{}
// 			input object
func (c *AnyValueArray) SetAsSingleObject(value interface{}) {
	a := convert.ToArray(value)
	c.value = a
}

// Gets the value stored in array element without any conversions.
// When element index is not defined it returns the entire array value.
// Parameters
// 			- indexint
// 			 an index of the element to get
// Returns interface{}
// the element value or value of the array when index is not defined.
func (c *AnyValueArray) GetAsObject(index int) interface{} {
	return c.Get(index)
}

// Sets a new value to array element specified by its index.
// When the index is not defined, it resets the entire array value.
// see
// ArrayConverter.toArray
// Parameters:
// 			 - index int
// 			 an index of the element to set
// 			 - value interface{}
// a new element or array value.
func (c *AnyValueArray) SetAsObject(index int, value interface{}) {
	c.Put(index, value)
}

// Converts array element into a string or returns nil if conversion is not possible.
// see
// StringConverter.toNullableString
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns *string
// string value of the element or nil if conversion is not supported.
func (c *AnyValueArray) GetAsNullableString(index int) *string {
	value := c.Get(index)
	return convert.StringConverter.ToNullableString(value)
}

// Converts array element into a string or returns "" if conversion is not possible.
// see
// getAsStringWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns string
// string value ot the element or "" if conversion is not supported.
func (c *AnyValueArray) GetAsString(index int) string {
	return c.GetAsStringWithDefault(index, "")
}

// Converts array element into a string or returns default value if conversion is not possible.
// see
// StringConverter.toStringWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// 			- defaultValue: string
// 			the default value
// Returns string
// string value ot the element or default value if conversion is not supported.
func (c *AnyValueArray) GetAsStringWithDefault(index int, defaultValue string) string {
	value := c.Get(index)
	return convert.StringConverter.ToStringWithDefault(value, defaultValue)
}

// Converts array element into a boolean or returns nil if conversion is not possible.
// see
// BooleanConverter.toNullableBoolean
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns bool
// boolean value of the element or nil if conversion is not supported.
func (c *AnyValueArray) GetAsNullableBoolean(index int) *bool {
	value := c.Get(index)
	return convert.BooleanConverter.ToNullableBoolean(value)
}

// Converts array element into a boolean or returns false if conversion is not possible.
// see
// getAsBooleanWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns bool
// boolean value ot the element or false if conversion is not supported.
func (c *AnyValueArray) GetAsBoolean(index int) bool {
	return c.GetAsBooleanWithDefault(index, false)
}

// Converts array element into a boolean or returns default value if conversion is not possible.
// see
// BooleanConverter.toBooleanWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// 			- defaultValue: boolean
// 			the default value
// Returns bool
// boolean value ot the element or default value if conversion is not supported.
func (c *AnyValueArray) GetAsBooleanWithDefault(index int, defaultValue bool) bool {
	value := c.Get(index)
	return convert.BooleanConverter.ToBooleanWithDefault(value, defaultValue)
}

// Converts array element into an integer or returns nil if conversion is not possible.
// see
// IntegerConverter.toNullableInteger
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns *int
// integer value of the element or nil if conversion is not supported.
func (c *AnyValueArray) GetAsNullableInteger(index int) *int {
	value := c.Get(index)
	return convert.IntegerConverter.ToNullableInteger(value)
}

// Converts array element into an integer or returns 0 if conversion is not possible.
// see
// getAsIntegerWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns int
// integer value ot the element or 0 if conversion is not supported.
func (c *AnyValueArray) GetAsInteger(index int) int {
	return c.GetAsIntegerWithDefault(index, 0)
}

// Converts array element into an integer or returns default value if conversion is not possible.
// see
// IntegerConverter.toIntegerWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// 			- defaultValue int
// 			the default value
// Returns int
// integer value ot the element or default value if conversion is not supported.
func (c *AnyValueArray) GetAsIntegerWithDefault(index int, defaultValue int) int {
	value := c.Get(index)
	return convert.IntegerConverter.ToIntegerWithDefault(value, defaultValue)
}

// Converts array element into a long or returns nil if conversion is not possible.
// see
// LongConverter.toNullableLong
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns *int64
// int64 value of the element or nil if conversion is not supported.
func (c *AnyValueArray) GetAsNullableLong(index int) *int64 {
	value := c.Get(index)
	return convert.LongConverter.ToNullableLong(value)
}

// Converts array element into a long or returns 0 if conversion is not possible.
// see
// getAsLongWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns int64
// int64 value ot the element or 0 if conversion is not supported.
func (c *AnyValueArray) GetAsLong(index int) int64 {
	return c.GetAsLongWithDefault(index, 0)
}

// Converts array element into a long or returns default value if conversion is not possible.
// see
// LongConverter.toLongWithDefault
// Parameters:
// 			 - index int
// 			an index of element to get.
// 			- defaultValue int64
// 			the default value
// Returns int64
// int64 value ot the element or default value if conversion is not supported.

func (c *AnyValueArray) GetAsLongWithDefault(index int, defaultValue int64) int64 {
	value := c.Get(index)
	return convert.LongConverter.ToLongWithDefault(value, defaultValue)
}

// Converts array element into a float or returns nil if conversion is not possible.
// see
// FloatConverter.toNullableFloat
// Parameters:
//           - index int
// an index of element to get.
// Returns *float64
// float64 value of the element or nil if conversion is not supported.
func (c *AnyValueArray) GetAsNullableFloat(index int) *float32 {
	value := c.Get(index)
	return convert.FloatConverter.ToNullableFloat(value)
}

// Converts array element into a float or returns 0 if conversion is not possible.
// see
// getAsFloatWithDefault
// Parameters:
// 			 - index int
// 			an index of element to get.
// Returns float32
// float value ot the element or 0 if conversion is not supported.
func (c *AnyValueArray) GetAsFloat(index int) float32 {
	return c.GetAsFloatWithDefault(index, 0)
}

// Converts array element into a float or returns default value if conversion is not possible.
// see
// FloatConverter.toFloatWithDefault
// Parameters:
// 			- index  int
// 			an index of element to get.
// defaultValue: number
// the default value
// Returns number
// float value ot the element or default value if conversion is not supported.
func (c *AnyValueArray) GetAsFloatWithDefault(index int, defaultValue float32) float32 {
	value := c.Get(index)
	return convert.FloatConverter.ToFloatWithDefault(value, defaultValue)
}

// Converts array element into a double or returns nil if conversion is not possible.
// see
// DoubleConverter.toNullableDouble
// Parameters:
// 			  - index int
// 			  an index of element to get.
// Returns *float64
// float64 value of the element or nil if conversion is not supported.
func (c *AnyValueArray) GetAsNullableDouble(index int) *float64 {
	value := c.Get(index)
	return convert.DoubleConverter.ToNullableDouble(value)
}

// Converts array element into a double or returns 0 if conversion is not possible.
// see
// getAsDoubleWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns float64
// double value ot the element or 0 if conversion is not supported.
func (c *AnyValueArray) GetAsDouble(index int) float64 {
	return c.GetAsDoubleWithDefault(index, 0)
}

// Converts array element into a double or returns default value if conversion is not possible.
// see
// DoubleConverter.toDoubleWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// 			- defaultValue: float64
// 			the default value
// Returns float64
// double value ot the element or default value if conversion is not supported.
func (c *AnyValueArray) GetAsDoubleWithDefault(index int, defaultValue float64) float64 {
	value := c.Get(index)
	return convert.DoubleConverter.ToDoubleWithDefault(value, defaultValue)
}

// Converts array element into a time.Time or returns nil if conversion is not possible.
// see
// DateTimeConverter.toNullableDateTime
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns *time.Time
// time.Time value of the element or nil if conversion is not supported.
func (c *AnyValueArray) GetAsNullableDateTime(index int) *time.Time {
	value := c.Get(index)
	return convert.DateTimeConverter.ToNullableDateTime(value)
}

// Converts array element into a time.Time or returns the current date if conversion is not possible.
// see
// getAsDateTimeWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns time.Time
// time.Time value ot the element or the current date if conversion is not supported.
func (c *AnyValueArray) GetAsDateTime(index int) time.Time {
	return c.GetAsDateTimeWithDefault(index, time.Time{})
}

// Converts array element into a time.Time or returns default value if conversion is not possible.
// see
// DateTimeConverter.toDateTimeWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// 			- defaultValue: time.Time
// 			the default value
// Returns time.Time
// time.time value ot the element or default value if conversion is not supported.
func (c *AnyValueArray) GetAsDateTimeWithDefault(index int, defaultValue time.Time) time.Time {
	value := c.Get(index)
	return convert.DateTimeConverter.ToDateTimeWithDefault(value, defaultValue)
}

// Converts array element into a time.Duration or returns nil if conversion is not possible.
// see
// DateTimeConverter.toNullableDateTime
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns *time.Duration
// time.Duration value of the element or nil if conversion is not supported.
func (c *AnyValueArray) GetAsNullableDuration(index int) *time.Duration {
	value := c.Get(index)
	return convert.DurationConverter.ToNullableDuration(value)
}

// Converts array element into a time.Duration or returns the current date if conversion is not possible.
// see
// getAsDateTimeWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns time.Duration
// Date value ot the element or the current date if conversion is not supported.
func (c *AnyValueArray) GetAsDuration(index int) time.Duration {
	return c.GetAsDurationWithDefault(index, 0*time.Millisecond)
}

// Converts array element into a time.Duration or returns default value if conversion is not possible.
// see
// DateTimeConverter.toDateTimeWithDefault
// Parameters:
// 			- index int
// 			an index of element to get.
// 			- defaultValue: ttime.Duration
// 			the default value
// Returns time.Duration
// Date value ot the element or default value if conversion is not supported.
func (c *AnyValueArray) GetAsDurationWithDefault(index int, defaultValue time.Duration) time.Duration {
	value := c.Get(index)
	return convert.DurationConverter.ToDurationWithDefault(value, defaultValue)
}

// Converts array element into a value defined by specied typecode. If conversion is not possible it returns nil.
// see
// TypeConverter.toNullableType
// Parameters
// 			- type: TypeCode
// 			the TypeCode that defined the type of the result
// 			- index int
// 			an index of element to get.
// Returns interface{}
// element value defined by the typecode or nil if conversion is not supported.
func (c *AnyValueArray) GetAsNullableType(typ convert.TypeCode, index int) interface{} {
	value := c.Get(index)
	return convert.TypeConverter.ToNullableType(typ, value)
}

// Converts array element into a value defined by specied typecode.
// If conversion is not possible it returns default value for the specified type.
// see
// getAsTypeWithDefault
// Parameters:
// 			- type TypeCode
// 				the TypeCode that defined the type of the result
// 			- index int
// 				an index of element to get.
// Returns interface{}
// element value defined by the typecode or default if conversion is not supported.
func (c *AnyValueArray) GetAsType(typ convert.TypeCode, index int) interface{} {
	return c.GetAsTypeWithDefault(typ, index, nil)
}

// Converts array element into a value defined by specied typecode.
// If conversion is not possible it returns default value.
// see
// TypeConverter.toTypeWithDefault
// Parameters:
// 			- type TypeCode
// 			the TypeCode that defined the type of the result
// 			- index int
// 			an index of element to get.
// 			- defaultValue interface{}
// 			the default value
// Returns interface{}
// element value defined by the typecode or default value if conversion is not supported.
func (c *AnyValueArray) GetAsTypeWithDefault(typ convert.TypeCode, index int, defaultValue interface{}) interface{} {
	value := c.Get(index)
	return convert.TypeConverter.ToTypeWithDefault(typ, value, defaultValue)
}

// Converts array element into an AnyValue or returns an empty AnyValue if conversion is not possible.
// see
// AnyValue
// see
// AnyValue.constructor
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns *AnyValue
// AnyValue value of the element or empty AnyValue if conversion is not supported.
func (c *AnyValueArray) GetAsValue(index int) *AnyValue {
	value := c.Get(index)
	return NewAnyValue(value)
}

// Converts array element into an AnyValueArray or returns nil if conversion is not possible.
// see
// fromValue
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns *AnyValueArray
// AnyValueArray value of the element or nil if conversion is not supported.
func (c *AnyValueArray) GetAsNullableArray(index int) *AnyValueArray {
	value := c.Get(index)
	if value != nil {
		return NewAnyValueArrayFromValue(value)
	} else {
		return nil
	}
}

// Converts array element into an AnyValueArray or returns empty AnyValueArray if conversion is not possible.
// see
// fromValue
// Parameters:
// 			 - index int
// 			an index of element to get.
// Returns *AnyValueArray
// AnyValueArray value of the element or empty AnyValueArray if conversion is not supported.
func (c *AnyValueArray) GetAsArray(index int) *AnyValueArray {
	value := c.Get(index)
	return NewAnyValueArrayFromValue(value)
}

// Converts array element into an AnyValueArray or returns default value if conversion is not possible.
// see
// getAsNullableArray
// Parameters:
// 			 - index int
// 				an index of element to get.
// 			 - defaultValue *AnyValueArray
// 				the default value
// Returns *AnyValueArray
// AnyValueArray value of the element or default value if conversion is not supported.
func (c *AnyValueArray) GetAsArrayWithDefault(index int, defaultValue *AnyValueArray) *AnyValueArray {
	result := c.GetAsNullableArray(index)
	if result != nil {
		return result
	} else {
		return defaultValue
	}
}

// Converts array element into an AnyValueMap or returns nil if conversion is not possible.
// see
// AnyValueMap
// see
// AnyValueMap.fromValue
// Parameters:
// 			- index int
// 			an index of element to get.
// Returns *AnyValueMap
// AnyValueMap value of the element or nil if conversion is not supported.
func (c *AnyValueArray) GetAsNullableMap(index int) *AnyValueMap {
	value := c.Get(index)
	if value != nil {
		return NewAnyValueMapFromValue(value)
	} else {
		return nil
	}
}

// Converts array element into an AnyValueMap or returns empty AnyValueMap if conversion is not possible.
// see
// AnyValueMap
// see
// AnyValueMap.fromValue
// Parameters:
// 			- index int
// an index of element to get.

// Returns *AnyValueMap
func (c *AnyValueArray) GetAsMap(index int) *AnyValueMap {
	value := c.Get(index)
	return NewAnyValueMapFromValue(value)
}

// Converts array element into an AnyValueMap or returns default value if conversion is not possible.
// see
// getAsNullableMap
// Parameters
// 			- index int
// 			an index of element to get.
// 			- defaultValue *AnyValueMap
// 			the default value
// Returns *AnyValueMap
// AnyValueMap value of the element or default value if conversion is not supported.
func (c *AnyValueArray) GetAsMapWithDefault(index int, defaultValue *AnyValueMap) *AnyValueMap {
	result := c.GetAsNullableMap(index)
	if result != nil {
		return NewAnyValueMapFromValue(result)
	} else {
		return defaultValue
	}
}

// Checks if this array contains a value. The check uses direct comparison between elements and the specified value.
// Parameters:
// 			- value interface{}
// 			a value to be checked
// Returns bool
// true if this array contains the value or false otherwise.
func (c *AnyValueArray) Contains(value interface{}) bool {
	for index := 0; index < c.Len(); index++ {
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

// Checks if this array contains a value.
//The check before comparison converts elements and the value to type specified by type code.
// see
// TypeConverter.toType
// see
// TypeConverter.toNullableType
// Type parameters

// Parameters:
// 				- typeCode TypeCode
// 				a type code that defines a type to convert values before comparison
// 				- value interface{}
// 				a value to be checked
// Returns bool
// true if this array contains the value or false otherwise.
func (c *AnyValueArray) ContainsAsType(typ convert.TypeCode, value interface{}) bool {
	typedValue := convert.TypeConverter.ToType(typ, value)

	for index := 0; index < c.Len(); index++ {
		thisTypedValue := convert.TypeConverter.ToType(typ, c.Get(index))

		if typedValue == thisTypedValue {
			return true
		}
	}

	return false
}

// Creates a binary clone of this object.
// Returns interface{}
// a clone of this object.
func (c *AnyValueArray) Clone() interface{} {
	return NewAnyValueArray(c.value)
}

func (c *AnyValueArray) String() string {
	builder := ""
	for index := 0; index < c.Len(); index++ {
		if index > 0 {
			builder += ","
		}
		builder = builder + c.GetAsStringWithDefault(index, "")
	}
	return builder
}

// Creates a new AnyValueArray from a list of values
// Parameters:
// 				- values ...values interface{}
// 				a list of values to initialize the created AnyValueArray
// Returns *AnyValueArray
// a newly created AnyValueArray.
func NewAnyValueArrayFromValues(values ...interface{}) *AnyValueArray {
	return NewAnyValueArray(values)
}

// Converts specified value into AnyValueArray.
// see
// ArrayConverter.toNullableArray
// Parameters:
// 			- value interface{}
// value to be converted
// Returns *AnyValueArray
// a newly created AnyValueArray.
func NewAnyValueArrayFromValue(value interface{}) *AnyValueArray {
	values := convert.ArrayConverter.ToArray(value)
	return NewAnyValueArray(values)
}

// Splits specified string into elements using a separator and assigns the elements to a newly created AnyValueArray.
// Parameters:
// 			- values string
// 			a string value to be split and assigned to AnyValueArray
// 			separator string
// 			a separator to split the string
// 			- removeDuplicates bool
// 			true to remove duplicated elements
// Returns *AnyValueArray
// a newly created AnyValueArray.
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
