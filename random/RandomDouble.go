package random

import (
	"math/rand"
)

type TRandomDouble struct{}

var RandomDouble *TRandomDouble = &TRandomDouble{}

func (c *TRandomDouble) NextDouble(min float64, max float64) float64 {
	if max-min <= 0 {
		return min
	}

	return min + rand.Float64()*(max-min)
}

func (c *TRandomDouble) UpdateDouble(value float64, interval float64) float64 {
	if interval <= 0 {
		interval = 0.1 * value
	}
	minValue := value - interval
	maxValue := value + interval
	return c.NextDouble(minValue, maxValue)
}
