package reflect

import (
	refl "reflect"
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
)

/*
Helper class to perform property introspection and dynamic writing.

In contrast to PropertyReflector which only introspects regular objects, this ObjectWriter is also able to handle maps and arrays. For maps properties are key-pairs identified by string keys, For arrays properties are elements identified by integer index.

This class has symmetric implementation across all languages supported by Pip.Services toolkit and used to support dynamic data processing.

Because all languages have different casing and case sensitivity rules, this ObjectWriter treats all property names as case insensitive.

see
PropertyReflector

Example:
myObj := MyObject{}

ObjectWriter.SetProperty(myObj, "myProperty", 123)

myMap := { key1: 123, key2: "ABC" }
ObjectWriter.SetProperty(myMap, "key1", "XYZ")

myArray := [1, 2, 3]
ObjectWriter.SetProperty(myArray, "0", 123)
*/

type TObjectWriter struct{}

var ObjectWriter = &TObjectWriter{}

// Gets a real object value. If object is a wrapper, it unwraps the value behind it. Otherwise it returns the same object value.
// Parameters:
// 			 - obj interface{}
// 			an object to unwrap..
// Returns interface{}
// an actual (unwrapped) object value.
func (c *TObjectWriter) GetValue(obj interface{}) interface{} {
	wrap, ok := obj.(IValueWrapper)
	if ok {
		obj = wrap.InnerValue()
	}

	return obj
}

// Sets value of object property specified by its name.
// The object can be a user defined object, map or array. The property name correspondently must be object property, map key or array index.
// If the property does not exist or introspection fails this method doesn't do anything and doesn't any throw errors.
// Parameters:
// 			- obj interface{}
// 			an object to write property to.
// 			- name string
// 			a name of the property to set.
// 			- value interface{}
// 			a new value for the property to set.
func (c *TObjectWriter) SetProperty(obj interface{}, name string, value interface{}) {
	if obj == nil || name == "" {
		return
	}

	obj = c.GetValue(obj)
	val := refl.ValueOf(obj)

	if val.Kind() == refl.Map {
		name = strings.ToLower(name)
		for _, v := range val.MapKeys() {
			key := convert.StringConverter.ToString(v.Interface())
			key = strings.ToLower(key)
			if name == key {
				val.SetMapIndex(v, refl.ValueOf(value))
				return
			}
		}
		val.SetMapIndex(refl.ValueOf(name), refl.ValueOf(value))
		return
	}

	if val.Kind() == refl.Slice || val.Kind() == refl.Array {
		index := convert.IntegerConverter.ToIntegerWithDefault(name, -1)

		// Todo: Think how to resize slice

		// Set array element
		if index >= 0 && index < val.Len() {
			v := val.Index(index)
			v.Set(refl.ValueOf(value))
			return
		}
		return
	}

	PropertyReflector.SetProperty(obj, name, value)
}

// Sets values of some (all) object properties.
// The object can be a user defined object, map or array. Property values correspondently are object properties, map key-pairs or array elements with their indexes.
// If some properties do not exist or introspection fails they are just silently skipped and no errors thrown.
// see
// setProperty
// Parameters:
// 			 - obj interface{}
// 			an object to write properties to.
// 			- values map[string]interface{}
// 			a map, containing property names and their values.
func (c *TObjectWriter) SetProperties(obj interface{}, values map[string]interface{}) {
	if values == nil || len(values) == 0 {
		return
	}

	for key, value := range values {
		c.SetProperty(obj, key, value)
	}
}
