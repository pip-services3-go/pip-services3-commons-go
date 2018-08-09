package test_reflect

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/reflect"
	"github.com/stretchr/testify/assert"
)

func TestTypeDescriptorFromString(t *testing.T) {
	descriptor, ok := reflect.ParseTypeDescriptorFromString("")
	assert.Nil(t, descriptor)
	assert.False(t, ok)

	descriptor, ok = reflect.ParseTypeDescriptorFromString("xxx,yyy")
	assert.Equal(t, "xxx", descriptor.Name())
	assert.Equal(t, "yyy", descriptor.Package())
	assert.True(t, ok)

	descriptor, ok = reflect.ParseTypeDescriptorFromString("xxx")
	assert.Equal(t, "xxx", descriptor.Name())
	assert.Equal(t, "", descriptor.Package())

	descriptor, ok = reflect.ParseTypeDescriptorFromString("xxx,yyy,zzz")
	assert.False(t, ok)
}
