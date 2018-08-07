package random

import (
	"reflect"
)

type TRandomArray struct{}

var RandomArray *TRandomArray = &TRandomArray{}

func (c *TRandomArray) Pick(value interface{}) interface{} {
	if value == nil {
		return nil
	}

	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return nil
	}

	len := v.Len()
	if len == 0 {
		return nil
	}

	index := RandomInteger.NextInteger(0, len-1)

	v = v.Index(index)
	return v.Interface()
}
