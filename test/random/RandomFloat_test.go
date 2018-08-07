package test_random

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/random"
	"github.com/stretchr/testify/assert"
)

func TestNextFloat(t *testing.T) {
	value := random.RandomFloat.NextFloat(0, 5)
	assert.True(t, value <= 5)

	value = random.RandomFloat.NextFloat(2, 5)
	assert.True(t, value <= 5 && value >= 2)
}

func TestUpdateFloat(t *testing.T) {
	value := random.RandomFloat.UpdateFloat(0, 5)
	assert.True(t, value <= 5 && value >= -5)

	value = random.RandomFloat.UpdateFloat(5, 0)

	value = random.RandomFloat.UpdateFloat(0, 0)
	assert.True(t, value == 0)
}
