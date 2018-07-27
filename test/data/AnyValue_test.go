package data

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/data"
	"github.com/stretchr/testify/assert"
)

func TestAnyValueGetAndSet(t *testing.T) {
	value := data.NewEmptyAnyValue()
	assert.Nil(t, value.GetAsObject())

	value.SetAsObject(1)
	assert.Equal(t, 1, value.GetAsInteger())
	assert.True(t, 1.0-value.GetAsFloat() < 0.001)
	assert.Equal(t, "1", value.GetAsString())
}

func TestAnyValueEquals(t *testing.T) {
	value := data.NewAnyValue(1)

	assert.True(t, value.Equals(1))
	assert.True(t, value.Equals(1.0))
	assert.True(t, value.Equals("1"))
	//	assert.True(t, value.EqualsAsType(convert.Float, "1"))
}
