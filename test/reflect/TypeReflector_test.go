package test_reflect

import (
	refl "reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pip-services-go/pip-services-commons-go/reflect"
)

func TestTypeReflectorCreate(t *testing.T) {
	typ := refl.TypeOf(TestClass{})
	obj, err := reflect.TypeReflector.CreateInstanceByType(typ)
	assert.NotNil(t, obj)
	assert.Nil(t, err)

	typ = refl.TypeOf((*TestClass)(nil))
	obj, err = reflect.TypeReflector.CreateInstanceByType(typ)
	assert.NotNil(t, obj)
	assert.Nil(t, err)
}
