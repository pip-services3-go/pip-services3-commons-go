package reflect

import (
	"strings"

	"github.com/pip-services-go/pip-services-commons-go/convert"
)

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

func (c *TRecursiveObjectWriter) SetProperties(obj interface{}, values map[string]interface{}) {
	if values == nil || len(values) == 0 {
		return
	}

	for key, value := range values {
		c.SetProperty(obj, key, value)
	}
}

func (c *TRecursiveObjectWriter) CopyProperties(dest interface{}, src interface{}) {
	if dest == nil || src == nil {
		return
	}

	values := RecursiveObjectReader.GetProperties(src)
	c.SetProperties(dest, values)
}
