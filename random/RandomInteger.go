package random

import (
	"math"
	"math/rand"
)

//
//Random generator for integer values.
//Example:
//
//  value1 := RandomInteger.nextInteger(5, 10);     // Possible result: 7
//  value2 := RandomInteger.nextInteger(10);        // Possible result: 3
//  value3 := RandomInteger.updateInteger(10, 3);   // Possible result: 9
type TRandomInteger struct{}

var RandomInteger *TRandomInteger = &TRandomInteger{}

//
//Generates a integer in the range ['min', 'max']. If 'max' is omitted, then the range
//will be set to [0, 'min'].
//
//Parameters:
//
//  - min: int - minimum value of the integer that will be generated. If 'max' is omitted,
//  then 'max' is set to 'min' and 'min' is set to 0.
//  - max: int - maximum value of the int that will be generated. Defaults to 'min' if omitted.
//
//Returns generated random integer value.
//

func (c *TRandomInteger) NextInteger(min int, max int) int {
	if max-min <= 0 {
		return min
	}

	return min + rand.Intn(max-min)
}

//
//Updates (drifts) a integer value within specified range defined
//
//Parameters:
//  - value: int - a integer value to drift.
//  - interval:int - a range. Default: 10% of the value
//
//Returns int
//
func (c *TRandomInteger) UpdateInteger(value int, interval int) int {
	if interval <= 0 {
		interval = int(math.Trunc(0.1 * float64(value)))
	}
	minValue := value - interval
	maxValue := value + interval
	return c.NextInteger(minValue, maxValue)
}

//
//Generates a random sequence of integers starting from 0 like: [0,1,2,3...??]
//
//Parameters:
//
//  - min: int - minimum value of the integer that will be generated. If 'max'
//  is omitted, then 'max' is set to 'min' and 'min' is set to 0.
//  - max: int - maximum value of the int that will be generated. Defaults to 'min' if omitted.
//
// Returns generated array of integers.
//
func (c *TRandomInteger) Sequence(min int, max int) []int {
	if min < 0 {
		min = 0
	}
	if max < min {
		max = min
	}

	count := c.NextInteger(min, max)

	result := make([]int, count, count)
	for i := range result {
		result[i] = i
	}

	return result
}
