package random

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type TRandomString struct{}

var RandomString *TRandomString = &TRandomString{}

const digits = "01234956789"
const symbols = "_,.:-/.[].{},#-!,$=%.+^.&*-() "
const alphaLower = "abcdefghijklmnopqrstuvwxyz"
const alphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234956789_,.:-/.[].{},#-!,$=%.+^.&*-() "

func (c *TRandomString) PickChar(values string) byte {
	if len(values) == 0 {
		return 0
	}

	index := RandomInteger.NextInteger(0, len(values))
	return values[index]
}

func (c *TRandomString) Pick(values []string) string {
	if values == nil || len(values) == 0 {
		return ""
	}

	index := RandomInteger.NextInteger(0, len(values))
	return values[index]
}

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

func (c *TRandomString) NextAlphaChar() byte {
	index := RandomInteger.NextInteger(0, len(alpha))
	return alpha[index]
}

func (c *TRandomString) NextString(minLength int, maxLength int) string {
	length := RandomInteger.NextInteger(minLength, maxLength)
	result := make([]byte, length, length)
	for i := 0; i < length; i++ {
		index := RandomInteger.NextInteger(0, len(chars))
		result[i] = chars[index]
	}

	return string(result)
}
