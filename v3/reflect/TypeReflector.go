package reflect

import (
	refl "reflect"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/errors"
)

/*
Helper class to perform object type introspection and object instantiation.

This class has symmetric implementation across all languages supported by Pip.Services toolkit and used to support dynamic data processing.

Because all languages have different casing and case sensitivity rules, this TypeReflector treats all type names as case insensitive.

see
TypeDescriptor

Example:
descriptor := NewTypeDescriptor("MyObject", "mylibrary");
TypeReflector.GetTypeByDescriptor(descriptor);
myObj = TypeReflector.CreateInstanceByDescriptor(descriptor);

*/
type TTypeReflector struct{}

var TypeReflector *TTypeReflector = &TTypeReflector{}

// Gets object type by its name and library where it is defined.
// Parameters:
// 			 - name string
// 			an object type name.
// 			pkg string
// 			a package where the type is defined
// Returns refl.Type
// the object type or nil is the type wasn't found.
func (c *TTypeReflector) GetType(name string, pkg string) refl.Type {
	// Dynamic type discovery is not supported
	// Todo: Add type discovery
	return nil
}

// Gets object type by type descriptor.
// Parameters:
// 			- descriptor *TypeDescriptor
// 			a type descriptor that points to an object type
// Returns refl.Type
// the object type or nil is the type wasn't found.
func (c *TTypeReflector) GetTypeByDescriptor(typ *TypeDescriptor) refl.Type {
	if typ == nil {
		panic("Type descriptor cannot be nil")
	}

	return c.GetType(typ.Name(), typ.Package())
}

// Creates an instance of an object type.
// Parameters:
// 			- type refl.Type
// 			an object type (factory function) to create.
// 			args ...interface{}
// 			arguments for the object constructor.
// Returns interface{}, error
// the created object instance and error.
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

// Creates an instance of an object type specified by its name and library where it is defined.
// Parameters:
// 				- name string
// 				an object type name.
// 				- pkg: string
// 				a package (module) where object type is defined.
// 				- args ...interface{}
// 				rguments for the object constructor.
// Returns any
// the created object instance.
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

// Creates an instance of an object type specified by type descriptor.
// Parameters:
// 			 - descriptor *TypeDescriptor
// 			a type descriptor that points to an object type
// 			- args ...interface{}
// 			arguments for the object constructor.

// Returns interface{}, error
// the created object instance and error.
func (c *TTypeReflector) CreateInstanceByDescriptor(typ *TypeDescriptor, args ...interface{}) (interface{}, error) {
	if typ == nil {
		err := errors.NewBadRequestError(
			"", "NO_TYPE_DESC", "Type descriptor cannot be nil",
		)
		return nil, err
	}

	return c.CreateInstance(typ.Name(), typ.Package(), args...)
}
