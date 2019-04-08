package reflect

import (
	refl "reflect"

	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

type TTypeReflector struct{}

var TypeReflector *TTypeReflector = &TTypeReflector{}

func (c *TTypeReflector) GetType(name string, pkg string) refl.Type {
	// Dynamic type discovery is not supported
	// Todo: Add type discovery
	return nil
}

func (c *TTypeReflector) GetTypeByDescriptor(typ *TypeDescriptor) refl.Type {
	if typ == nil {
		panic("Type descriptor cannot be nil")
	}

	return c.GetType(typ.Name(), typ.Package())
}

func (c *TTypeReflector) CreateInstanceByType(typ refl.Type, args ...interface{}) (interface{}, error) {
	if len(args) > 0 {
		err := errors.NewBadRequestError(
			"", "ARGS_NOT_SUPPORTED", "Constructors with arguments are not supported",
		)
		return nil, err
	}

	if typ.Kind() == refl.Ptr {
		obj := refl.New(typ.Elem()).Elem().Interface()
		return obj, nil
	}

	obj := refl.New(typ).Interface()
	return obj, nil
}

func (c *TTypeReflector) CreateInstance(name string, pkg string, args ...interface{}) (interface{}, error) {
	typ := c.GetType(name, pkg)

	if typ == nil {
		err := errors.NewNotFoundError(
			"", "TYPE_NOT_FOUND", "Type "+name+","+pkg+" was not found",
		).WithDetails("type", name).WithDetails("package", pkg)
		return nil, err
	}

	return c.CreateInstanceByType(typ, args)
}

func (c *TTypeReflector) CreateInstanceByDescriptor(typ *TypeDescriptor, args ...interface{}) (interface{}, error) {
	if typ == nil {
		err := errors.NewBadRequestError(
			"", "NO_TYPE_DESC", "Type descriptor cannot be nil",
		)
		return nil, err
	}

	return c.CreateInstance(typ.Name(), typ.Package(), args...)
}
