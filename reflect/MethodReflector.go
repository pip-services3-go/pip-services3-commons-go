package reflect

import (
	refl "reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TMethodReflector struct{}

var MethodReflector = &TMethodReflector{}

func (c *TMethodReflector) matchMethod(method refl.Method, name string) bool {
	// Method must be public and match to name as case insensitive
	r, _ := utf8.DecodeRuneInString(method.Name)
	return unicode.IsUpper(r) &&
		strings.ToLower(method.Name) == strings.ToLower(name)
}

func (c *TMethodReflector) HasMethod(obj interface{}, name string) bool {
	if obj == nil {
		panic("Object cannot be nil")
	}
	if name == "" {
		panic("Method name cannot be empty")
	}

	objType := refl.TypeOf(obj)

	for index := 0; index < objType.NumMethod(); index++ {
		method := objType.Method(index)
		if c.matchMethod(method, name) {
			return true
		}
	}

	return false
}

func (c *TMethodReflector) InvokeMethod(obj interface{}, name string, args ...interface{}) interface{} {
	if obj == nil {
		panic("Object cannot be nil")
	}
	if name == "" {
		panic("Method name cannot be empty")
	}

	defer func() {
		// Do nothing and return nil
		recover()
	}()

	objType := refl.TypeOf(obj)

	// Convert arguments
	inputs := make([]refl.Value, len(args))
	for index := range args {
		inputs[index] = refl.ValueOf(args[index])
	}

	for index := 0; index < objType.NumMethod(); index++ {
		method := objType.Method(index)
		if c.matchMethod(method, name) {
			results := refl.ValueOf(obj).Method(index).Call(inputs)
			if len(results) == 0 {
				return nil
			}
			return results[0].Interface()
		}
	}

	return nil
}

func (c *TMethodReflector) GetMethodNames(obj interface{}) []string {
	methods := []string{}

	objType := refl.TypeOf(obj)

	for index := 0; index < objType.NumMethod(); index++ {
		method := objType.Method(index)
		if c.matchMethod(method, method.Name) {
			methods = append(methods, method.Name)
		}
	}

	return methods
}
