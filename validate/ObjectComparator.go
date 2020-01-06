package validate

import (
	"regexp"
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/v3/data"
)

/*
Helper class to perform comparison operations over arbitrary values.

Example
ObjectComparator.Compare(2, "GT", 1);        // Result: true
ObjectComparator.AreEqual("A", "B");         // Result: false
*/

type TObjectComparator struct{}

var ObjectComparator = &TObjectComparator{}

// Perform comparison operation over two arguments. The operation can be performed over values of any type.
// Parameters:
// 			- value1 interface{}
// 			the first argument to compare
// 			operation string
// 			the comparison operation: "==" ("=", "EQ"), "!= " ("<>", "NE"); "<"/">" ("LT"/"GT"), "<="/">=" ("LE"/"GE"); "LIKE".
// 			- value2 interface{}
// 			the second argument to compare
// Returns bool
// result of the comparison operation
func (c *TObjectComparator) Compare(value1 interface{}, operation string, value2 interface{}) bool {
	operation = strings.ToUpper(operation)

	if operation == "=" || operation == "==" || operation == "EQ" {
		return c.AreEqual(value1, value2)
	}
	if operation == "!=" || operation == "<>" || operation == "NE" {
		return c.AreNotEqual(value1, value2)
	}
	if operation == "<" || operation == "LT" {
		return c.IsLess(value1, value2)
	}
	if operation == "<=" || operation == "LE" || operation == "LTE" {
		return c.AreEqual(value1, value2) || c.IsLess(value1, value2)
	}
	if operation == ">" || operation == "GT" {
		return c.IsGreater(value1, value2)
	}
	if operation == ">=" || operation == "GE" || operation == "GTE" {
		return c.AreEqual(value1, value2) || c.IsGreater(value1, value2)
	}
	if operation == "LIKE" {
		return c.Match(value1, value2)
	}

	return false
}

// Checks if two values are equal. The operation can be performed over values of any type.
// Parameters:
// 			- value1 interface
// 			the first value to compare
// 			value2 interface{}
// 			the second value to compare
// Returns bool
// true if values are equal and false otherwise
func (c *TObjectComparator) AreEqual(value1 interface{}, value2 interface{}) bool {
	if value1 == nil && value2 == nil {
		return true
	}
	if value1 == nil || value2 == nil {
		return false
	}

	equatable, ok := value1.(data.IEquatable)
	if ok {
		return equatable.Equals(value2)
	}
	equatable, ok = value2.(data.IEquatable)
	if ok {
		return equatable.Equals(value1)
	}

	number1 := convert.DoubleConverter.ToNullableDouble(value1)
	number2 := convert.DoubleConverter.ToNullableDouble(value2)
	if number1 != nil && number2 != nil {
		return *number1 == *number2
	}

	str1 := convert.StringConverter.ToNullableString(value1)
	str2 := convert.StringConverter.ToNullableString(value1)
	if str1 == nil && str2 == nil {
		return *str1 == *str2
	}

	return value1 == value2
}

// Checks if two values are NOT equal The operation can be performed over values of any type.
// Parameters:
// 			 - value1 interface{}
// 			 the first value to compare
// 			 - value2 interface{}
// 			 the second value to compare
// Returns bool
// true if values are NOT equal and false otherwise
func (c *TObjectComparator) AreNotEqual(value1 interface{}, value2 interface{}) bool {
	return !c.AreEqual(value1, value2)
}

// Checks if first value is less than the second one. The operation can be performed over numbers or strings.
// Parameters:
// 			- value1 interface{}
// 			the first value to compare
// 			- value2 interface{}
// 			the second value to compare
// Returns bool
// true if the first value is less than second and false otherwise.
func (c *TObjectComparator) IsLess(value1 interface{}, value2 interface{}) bool {
	number1 := convert.DoubleConverter.ToNullableDouble(value1)
	number2 := convert.DoubleConverter.ToNullableDouble(value2)

	if number1 == nil || number2 == nil {
		return false
	}

	return *number1 < *number2
}

// Checks if first value is greater than the second one. The operation can be performed over numbers or strings.
// Parameters:
// 			- value1 interface{}
// 			the first value to compare
// 			- value2 interface{}
// 			the second value to compare
// Returns bool
// true if the first value is greater than second and false otherwise.
func (c *TObjectComparator) IsGreater(value1 interface{}, value2 interface{}) bool {
	number1 := convert.DoubleConverter.ToNullableDouble(value1)
	number2 := convert.DoubleConverter.ToNullableDouble(value2)

	if number1 == nil || number2 == nil {
		return false
	}

	return *number1 > *number2
}

// Checks if string  views are matches
// Parameters:
// 			 - value1 interface{}
// 			 a string value to match
// 			 - value1 interface{}
// 			 a string value to match
// Returns bool
// true if the value matches regular expression and false otherwise.
func (c *TObjectComparator) Match(value1 interface{}, value2 interface{}) bool {
	if value1 == nil && value2 == nil {
		return true
	}
	if value1 == nil || value2 == nil {
		return false
	}

	str1 := convert.StringConverter.ToString(value1)
	str2 := convert.StringConverter.ToString(value2)

	matched, _ := regexp.MatchString(str2, str1)
	return matched
}
