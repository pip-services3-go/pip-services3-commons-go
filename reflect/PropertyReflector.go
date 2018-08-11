package reflect

import (
	refl "reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TPropertyReflector struct{}

var PropertyReflector = &TPropertyReflector{}

func (c *TPropertyReflector) toFieldType(obj interface{}) refl.Type {
	// Unwrap value
	wrap, ok := obj.(IValueWrapper)
	if ok {
		obj = wrap.InnerValue()
	}

	// Move from pointer to real struct
	typ := refl.TypeOf(obj)
	if typ.Kind() == refl.Ptr {
		typ = typ.Elem()
	}

	return typ
}

func (c *TPropertyReflector) toPropertyType(obj interface{}) refl.Type {
	return refl.TypeOf(obj)
}

func (c *TPropertyReflector) matchField(field refl.StructField, name string) bool {
	// Field must be public and match to name as case insensitive
	r, _ := utf8.DecodeRuneInString(field.Name)
	return unicode.IsUpper(r) &&
		strings.ToLower(field.Name) == strings.ToLower(name)
}

func (c *TPropertyReflector) matchPropertyGetter(property refl.Method, name string) bool {
	if property.Type.NumIn() != 1 || property.Type.NumOut() != 1 {
		return false
	}

	// Method must be public and match to name as case insensitive
	r, _ := utf8.DecodeRuneInString(property.Name)
	return unicode.IsUpper(r) &&
		strings.ToLower(property.Name) == strings.ToLower(name)
}

func (c *TPropertyReflector) matchPropertySetter(property refl.Method, name string) bool {
	if property.Type.NumIn() != 2 || property.Type.NumOut() != 0 {
		return false
	}

	// Method must be public and match to name as case insensitive
	r, _ := utf8.DecodeRuneInString(property.Name)
	name = "Set" + name
	return unicode.IsUpper(r) &&
		strings.ToLower(property.Name) == strings.ToLower(name)
}

func (c *TPropertyReflector) HasProperty(obj interface{}, name string) bool {
	if obj == nil {
		panic("Object cannot be nil")
	}
	if name == "" {
		panic("Property name cannot be empty")
	}

	fieldType := c.toFieldType(obj)
	if fieldType.Kind() == refl.Struct {
		for index := 0; index < fieldType.NumField(); index++ {
			field := fieldType.Field(index)
			if c.matchField(field, name) {
				return true
			}
		}
	}

	propType := c.toPropertyType(obj)
	for index := 0; index < propType.NumMethod(); index++ {
		method := propType.Method(index)
		if c.matchPropertyGetter(method, name) {
			return true
		}
	}

	return false
}

func (c *TPropertyReflector) GetProperty(obj interface{}, name string) interface{} {
	if obj == nil {
		panic("Object cannot be nil")
	}
	if name == "" {
		panic("Property name cannot be empty")
	}

	defer func() {
		// Do nothing and return nil
		recover()
	}()

	fieldType := c.toFieldType(obj)
	if fieldType.Kind() == refl.Struct {
		for index := 0; index < fieldType.NumField(); index++ {
			field := fieldType.Field(index)
			if c.matchField(field, name) {
				val := refl.ValueOf(obj)
				if val.Kind() == refl.Ptr {
					val = val.Elem()
				}
				return val.Field(index).Interface()
			}
		}
	}

	propType := c.toPropertyType(obj)
	for index := 0; index < propType.NumMethod(); index++ {
		method := propType.Method(index)
		if c.matchPropertyGetter(method, name) {
			val := refl.ValueOf(obj)
			return val.Method(index).Call([]refl.Value{})[0].Interface()
		}
	}

	return nil
}

func (c *TPropertyReflector) GetPropertyNames(obj interface{}) []string {
	if obj == nil {
		panic("Object cannot be nil")
	}

	properties := []string{}

	fieldType := c.toFieldType(obj)
	if fieldType.Kind() == refl.Struct {
		for index := 0; index < fieldType.NumField(); index++ {
			field := fieldType.Field(index)
			if c.matchField(field, field.Name) {
				properties = append(properties, field.Name)
			}
		}
	}

	propType := c.toPropertyType(obj)
	for index := 0; index < propType.NumMethod(); index++ {
		method := propType.Method(index)
		if c.matchPropertyGetter(method, method.Name) {
			properties = append(properties, method.Name)
		}
	}

	return properties
}

func (c *TPropertyReflector) GetProperties(obj interface{}) map[string]interface{} {
	if obj == nil {
		panic("Object cannot be nil")
	}

	defer func() {
		// Do nothing and return nil
		recover()
	}()

	properties := map[string]interface{}{}

	fieldType := c.toFieldType(obj)
	if fieldType.Kind() == refl.Struct {
		for index := 0; index < fieldType.NumField(); index++ {
			field := fieldType.Field(index)
			if c.matchField(field, field.Name) {
				val := refl.ValueOf(obj)
				if val.Kind() == refl.Ptr {
					val = val.Elem()
				}
				properties[field.Name] = val.Field(index).Interface()
			}
		}
	}

	propType := c.toPropertyType(obj)
	for index := 0; index < propType.NumMethod(); index++ {
		method := propType.Method(index)
		if c.matchPropertyGetter(method, method.Name) {
			val := refl.ValueOf(obj)
			properties[method.Name] = val.Method(index).Call([]refl.Value{})[0].Interface()
		}
	}

	return properties
}

func (c *TPropertyReflector) SetProperty(obj interface{}, name string, value interface{}) {
	if obj == nil {
		panic("Object cannot be nil")
	}
	if name == "" {
		panic("Property name cannot be empty")
	}

	defer func() {
		// Do nothing and return nil
		recover()
	}()

	fieldType := c.toFieldType(obj)
	if fieldType.Kind() == refl.Struct {
		for index := 0; index < fieldType.NumField(); index++ {
			field := fieldType.Field(index)
			if c.matchField(field, name) {
				val := refl.ValueOf(obj)
				if val.Kind() == refl.Ptr {
					val = val.Elem()
				}
				val.Field(index).Set(refl.ValueOf(value))
				return
			}
		}
	}

	propType := c.toPropertyType(obj)
	for index := 0; index < propType.NumMethod(); index++ {
		method := propType.Method(index)
		if c.matchPropertySetter(method, name) {
			val := refl.ValueOf(obj)
			val.Method(index).Call([]refl.Value{refl.ValueOf(value)})
		}
	}
}

func (c *TPropertyReflector) SetProperties(obj interface{}, values map[string]interface{}) {
	if values == nil || len(values) == 0 {
		return
	}

	for key, value := range values {
		c.SetProperty(obj, key, value)
	}
}
