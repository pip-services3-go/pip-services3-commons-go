package test_random

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/random"
	"github.com/stretchr/testify/assert"
)

func TestArrayPick(t *testing.T) {
	array1 := []interface{}{}
	value1 := random.RandomArray.Pick(array1)
	assert.Nil(t, value1)

	array2 := []interface{}{nil, nil}
	value2 := random.RandomArray.Pick(array2)
	assert.Nil(t, value2)

	array3 := []int{}
	assert.Nil(t, random.RandomArray.Pick(array3))

	array4 := []int{1, 2}
	value4 := random.RandomArray.Pick(array4).(int)
	assert.True(t, value4 == 1 || value4 == 2)
}
