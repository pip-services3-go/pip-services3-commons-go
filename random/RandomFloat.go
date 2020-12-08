package random

import (
	"math/rand"
)

//
//Random generator for float values.
//
//Examples:
//
//  value1 := RandomFloat.nextFloat(5, 10);     // Possible result: 7.3
//  value2 := RandomFloat.nextFloat(10);        // Possible result: 3.7
//  value3 := RandomFloat.updateFloat(10, 3);   // Possible result: 9.2
type TRandomFloat struct{}

var RandomFloat *TRandomFloat = &TRandomFloat{}

//Generates a float in the range ['min', 'max']. If 'max' is omitted,
//then the range will be set to [0, 'min'].
//
//Parameters:
//			- min: float32 - minimum value of the float that will be generated.
//			If 'max' is omitted, then 'max' is set to 'min' and 'min' is set to 0.
//			- max: float32 - maximum value of the float that will be generated.
//			Defaults to 'min' if omitted.
//
//Returns generated random float32 value.
//
func (c *TRandomFloat) NextFloat(min float32, max float32) float32 {
	if max-min <= 0 {
		return min
	}

	return min + rand.Float32()*(max-min)
}

//Updates (drifts) a float value within specified range defined
//Parameters:
//
//			- value: float32 - value to drift.
//			- interval: float32 - a range. Default: 10% of the value
//
//Returns float32
//
func (c *TRandomFloat) UpdateFloat(value float32, interval float32) float32 {
	if interval <= 0 {
		interval = 0.1 * value
	}
	minValue := value - interval
	maxValue := value + interval
	return c.NextFloat(minValue, maxValue)
}
