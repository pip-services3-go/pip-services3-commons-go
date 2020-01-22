package test_random

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/random"
	"github.com/stretchr/testify/assert"
)

func TestNextDouble(t *testing.T) {
	value := random.RandomDouble.NextDouble(0, 5)
	assert.True(t, value <= 5)

	value = random.RandomDouble.NextDouble(2, 5)
	assert.True(t, value <= 5 && value >= 2)
}

func TestUpdateDouble(t *testing.T) {
	value := random.RandomDouble.UpdateDouble(0, 5)
	assert.True(t, value <= 5 && value >= -5)

	value = random.RandomDouble.UpdateDouble(5, 0)

	value = random.RandomDouble.UpdateDouble(0, 0)
	assert.True(t, value == 0)
}
