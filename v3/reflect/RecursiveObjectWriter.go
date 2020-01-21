package reflect

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/convert"
)

/*
Helper class to perform property introspection and dynamic writing.

It is similar to ObjectWriter but writes properties recursively through the entire object graph. Nested property names are defined using dot notation as "object.subobject.property"
*/
type TRecursiveObjectWriter struct{}

var RecursiveObjectWriter = &TRecursiveObjectWriter{}

func (c *TRecursiveObjectWriter) createProperty(obj interface{}, names []string, nameIndex int) interface{} {
	// Todo: Complete implementation
	// If next field is index then create an array
	subField := ""
	if len(names) > nameIndex+1 {
		subField = names[nameIndex+1]
	}
	subFieldIndex := convert.IntegerConverter.ToNullableInteger(subField)
	if subFieldIndex != nil {
		return []interface{}{}
	}

	// Else create a dictionary
	return map[string]interface{}{}
}

func (c *TRecursiveObjectWriter) performSetProperty(obj interface{}, names []string, nameIndex int, value interface{}) {
	if nameIndex < len(names)-1 {
		subObj := ObjectReader.GetProperty(obj, names[nameIndex])
		if subObj != nil {
			c.performSetProperty(subObj, names, nameIndex+1, value)
		} else {
			subObj = c.createProperty(obj, names, nameIndex)
			if subObj != nil {
				c.performSetProperty(subObj, names, nameIndex+1, value)
				ObjectWriter.SetProperty(obj, names[nameIndex], subObj)
			}
		}
	} else {
		ObjectWriter.SetProperty(obj, names[nameIndex], value)
	}
}

// Recursively sets value of object and its subobjects property specified by its name.
// The object can be a user defined object, map or array. The property name correspondently must be object property, map key or array index.
// If the property does not exist or introspection fails this method doesn't do anything and doesn't any throw errors.
// Parameters:
// 				- obj interface{}
// 				an object to write property to.
// 				- name string
// 				a name of the property to set.
// 				- value interface{}
// 				a new value for the property to set.

func (c *TRecursiveObjectWriter) SetProperty(obj interface{}, name string, value interface{}) {
	if obj == nil || name == "" {
		return
	}

	names := strings.Split(name, ".")
	if len(names) == 0 {
		return
	}

	c.performSetProperty(obj, names, 0, value)
}

// Recursively sets values of some (all) object and its subobjects properties.
// The object can be a user defined object, map or array. Property values correspondently are object properties, map key-pairs or array elements with their indexes.
// If some properties do not exist or introspection fails they are just silently skipped and no errors thrown.
// see
// SetProperty
// Parameters:
// 			- obj interface{}
// 			an object to write properties to.
// 			- values map[atring]inteerface{}
// 			a map, containing property names and their values.
func (c *TRecursiveObjectWriter) SetProperties(obj interface{}, values map[string]interface{}) {
	if values == nil || len(values) == 0 {
		return
	}

	for key, value := range values {
		c.SetProperty(obj, key, value)
	}
}

// Copies content of one object to another object by recursively reading all properties from source object and then recursively writing them to destination object.
// Parameters:
// 			- dest interface{}
// 			a destination object to write properties to.
// 			- src interface{}
// 			a source object to read properties from
func (c *TRecursiveObjectWriter) CopyProperties(dest interface{}, src interface{}) {
	if dest == nil || src == nil {
		return
	}

	values := RecursiveObjectReader.GetProperties(src)
	c.SetProperties(dest, values)
}
