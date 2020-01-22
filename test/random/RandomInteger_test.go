package test_random

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/random"
	"github.com/stretchr/testify/assert"
)

func TestNextInteger(t *testing.T) {
	value := random.RandomInteger.NextInteger(0, 5)
	assert.True(t, value <= 5)

	value = random.RandomInteger.NextInteger(2, 5)
	assert.True(t, value <= 5 && value >= 2)
}

func TestUpdateInteger(t *testing.T) {
	value := random.RandomInteger.UpdateInteger(0, 5)
	assert.True(t, value <= 5 && value >= -5)

	value = random.RandomInteger.UpdateInteger(5, 0)

	value = random.RandomInteger.UpdateInteger(0, 0)
	assert.True(t, value == 0)
}

func TestIntegerSequence(t *testing.T) {
	list := random.RandomInteger.Sequence(1, 5)
	assert.True(t, len(list) <= 5 && len(list) >= 1)

	list = random.RandomInteger.Sequence(-1, 0)
	assert.True(t, len(list) == 0)

	list = random.RandomInteger.Sequence(-1, -4)
	assert.True(t, len(list) == 0)

	list = random.RandomInteger.Sequence(4, 4)
	assert.True(t, len(list) == 4)

	list = random.RandomInteger.Sequence(5, 5)
	assert.True(t, len(list) == 5)
}
