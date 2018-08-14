package validate

import (
	"regexp"
	"strings"

	"github.com/pip-services-go/pip-services-commons-go/convert"
)

type TObjectComparator struct{}

var ObjectComparator = &TObjectComparator{}

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

func (c *TObjectComparator) AreEqual(value1 interface{}, value2 interface{}) bool {
	if value1 == nil && value2 == nil {
		return true
	}
	if value1 == nil || value2 == nil {
		return false
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

func (c *TObjectComparator) AreNotEqual(value1 interface{}, value2 interface{}) bool {
	return !c.AreEqual(value1, value2)
}

func (c *TObjectComparator) IsLess(value1 interface{}, value2 interface{}) bool {
	number1 := convert.DoubleConverter.ToNullableDouble(value1)
	number2 := convert.DoubleConverter.ToNullableDouble(value2)

	if number1 == nil || number2 == nil {
		return false
	}

	return *number1 < *number2
}

func (c *TObjectComparator) IsGreater(value1 interface{}, value2 interface{}) bool {
	number1 := convert.DoubleConverter.ToNullableDouble(value1)
	number2 := convert.DoubleConverter.ToNullableDouble(value2)

	if number1 == nil || number2 == nil {
		return false
	}

	return *number1 > *number2
}

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
