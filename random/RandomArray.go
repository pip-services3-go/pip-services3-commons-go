package random

import (
	"reflect"
)

//
//Random generator for array objects.
//
// Examples:
//
//  value1 := RandomArray.pick([1, 2, 3, 4]); // Possible result: 3
//
type TRandomArray struct{}

var RandomArray *TRandomArray = &TRandomArray{}

//Picks a random element from specified array.
//Parameters:
//
//  - values: an array of any interface
//
//Returns a randomly picked item.
//
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
