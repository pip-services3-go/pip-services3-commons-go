package test_random

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/random"
	"github.com/stretchr/testify/assert"
)

func TestNextLong(t *testing.T) {
	value := random.RandomLong.NextLong(0, 5)
	assert.True(t, value <= 5)

	value = random.RandomLong.NextLong(2, 5)
	assert.True(t, value <= 5 && value >= 2)
}

func TestUpdateLong(t *testing.T) {
	value := random.RandomLong.UpdateLong(0, 5)
	assert.True(t, value <= 5 && value >= -5)

	value = random.RandomLong.UpdateLong(5, 0)

	value = random.RandomLong.UpdateLong(0, 0)
	assert.True(t, value == 0)
}

func TestLongSequence(t *testing.T) {
	list := random.RandomLong.Sequence(1, 5)
	assert.True(t, len(list) <= 5 && len(list) >= 1)

	list = random.RandomLong.Sequence(-1, 0)
	assert.True(t, len(list) == 0)

	list = random.RandomLong.Sequence(-1, -4)
	assert.True(t, len(list) == 0)

	list = random.RandomLong.Sequence(4, 4)
	assert.True(t, len(list) == 4)

	list = random.RandomLong.Sequence(5, 5)
	assert.True(t, len(list) == 5)
}
