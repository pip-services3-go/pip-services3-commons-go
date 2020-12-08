package random

import (
	"math/rand"
)

//Random generator for double values.
//
//Example:
//
//  value1 := RandomDouble.nextDouble(5, 10);     // Possible result: 7.3
//  value2 := RandomDouble.nextDouble(10);        // Possible result: 3.7
//  value3 := RandomDouble.updateDouble(10, 3);   // Possible result: 9.2
type TRandomDouble struct{}

var RandomDouble *TRandomDouble = &TRandomDouble{}

//Generates a random double value in the range ['min', 'max'].
//Parameters:
//
//  - min: float64 - minimum range value
//  - max: float64 - max range value
//
//Returns: float64 - a random value.
//
func (c *TRandomDouble) NextDouble(min float64, max float64) float64 {
	if max-min <= 0 {
		return min
	}

	return min + rand.Float64()*(max-min)
}

//Updates (drifts) a double value within specified range defined
//Parameters:
//
//  - value: float64 - value to drift.
//  - interval: float64 - a range to drift. Default: 10% of the value
//
//Returns float64
//
func (c *TRandomDouble) UpdateDouble(value float64, interval float64) float64 {
	if interval <= 0 {
		interval = 0.1 * value
	}
	minValue := value - interval
	maxValue := value + interval
	return c.NextDouble(minValue, maxValue)
}
