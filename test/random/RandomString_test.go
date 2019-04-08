package test_random

import (
	"strings"
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/random"
	"github.com/stretchr/testify/assert"
)

const symbols = "_,.:-/.[].{},#-!,$=%.+^.&*-() "
const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "01234956789"

func TestPickString(t *testing.T) {
	assert.True(t, random.RandomString.PickChar("") == 0)
	charVariable := random.RandomString.PickChar(chars)
	assert.True(t, strings.IndexByte(chars, charVariable) >= 0)

	valuesEmpty := []string{}
	assert.True(t, random.RandomString.Pick(valuesEmpty) == "")

	values := []string{"ab", "cd"}
	result := random.RandomString.Pick(values)
	assert.True(t, result == "ab" || result == "cd")
}

func TestDistort(t *testing.T) {
	value := random.RandomString.Distort("abc")
	assert.True(t, len(value) == 3 || len(value) == 4)
	value1 := value[:3]
	assert.True(t, strings.Compare(value1, "Abc") == 0 || strings.Compare(value1, "abc") == 0)

	if len(value) == 4 {
		assert.True(t, strings.IndexByte(symbols, value[3]) >= 0)
	}
}

func TestNextAlpaChar(t *testing.T) {
	assert.True(t, strings.IndexByte(chars, random.RandomString.NextAlphaChar()) >= 0)
}

func TestNextString(t *testing.T) {
	value := random.RandomString.NextString(3, 5)
	assert.True(t, len(value) <= 5 && len(value) >= 3)

	for i := 0; i < len(value); i++ {
		assert.True(t, strings.IndexByte(chars, value[i]) >= 0 || strings.IndexByte(symbols, value[i]) >= 0 || strings.IndexByte(digits, value[i]) >= 0)
	}
}
