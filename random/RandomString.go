package random

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

//Random generator for string values.
//
//Example:
//
//  value1 := RandomString.pickChar("ABC");     // Possible result: "C"
//  value2 := RandomString.pick(["A","B","C"]); // Possible result: "gBW"
type TRandomString struct{}

var RandomString *TRandomString = &TRandomString{}

const digits = "01234956789"
const symbols = "_,.:-/.[].{},#-!,$=%.+^.&*-() "
const alphaLower = "abcdefghijklmnopqrstuvwxyz"
const alphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234956789_,.:-/.[].{},#-!,$=%.+^.&*-() "

//
//Picks a random character from a string.
//
//Parameters:
//
//			- values: string to pick a char from
//
//Returnsa randomly picked char.
//
func (c *TRandomString) PickChar(values string) byte {
	if len(values) == 0 {
		return 0
	}

	index := RandomInteger.NextInteger(0, len(values))
	return values[index]
}

//
//Picks a random string from an array of string.
//
//Parameters:
//
//			- values: string[] strings to pick from.
//
//Returns a randomly picked string.
//
func (c *TRandomString) Pick(values []string) string {
	if values == nil || len(values) == 0 {
		return ""
	}

	index := RandomInteger.NextInteger(0, len(values))
	return values[index]
}

//
//Distorts a string by randomly replacing characters in it.
//
//Parameters:
//
//			-value: string - a string to distort.
//
//Returns a distored string.
//

func (c *TRandomString) Distort(value string) string {
	if value == "" {
		return ""
	}

	value = strings.ToLower(value)

	//Capitalize the first letter of the string 'value'.
	if RandomBoolean.Chance(1, 5) {
		r, n := utf8.DecodeRuneInString(value)
		value = string(unicode.ToUpper(r)) + value[n:]
	}

	//Add a symbol to the end of the string 'value'
	if RandomBoolean.Chance(1, 3) {
		value = value + string([]byte{RandomString.PickChar(symbols)})
	}

	return value
}

//
//Generates random alpha characted [A-Za-z]
//
//Returns a random characted.
//
func (c *TRandomString) NextAlphaChar() byte {
	index := RandomInteger.NextInteger(0, len(alpha))
	return alpha[index]
}

//
//Generates a random string, consisting of upper and lower case letters (of the English alphabet), digits (0-9), and symbols ("_,.:-/.[].{},#-!,$=%.+^.&*-() ").
//
//Parameters:
//
//			- minLength: int - minimum string length.
//			- maxLength: int - maximum string length.
//
//Returns a random string.
//
func (c *TRandomString) NextString(minLength int, maxLength int) string {
	length := RandomInteger.NextInteger(minLength, maxLength)
	result := make([]byte, length, length)
	for i := 0; i < length; i++ {
		index := RandomInteger.NextInteger(0, len(chars))
		result[i] = chars[index]
	}

	return string(result)
}
