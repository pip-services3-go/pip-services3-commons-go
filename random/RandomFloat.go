package random

import (
	"math/rand"
)

type TRandomFloat struct{}

var RandomFloat *TRandomFloat = &TRandomFloat{}

func (c *TRandomFloat) NextFloat(min float32, max float32) float32 {
	if max-min <= 0 {
		return min
	}

	return min + rand.Float32()*(max-min)
}

func (c *TRandomFloat) UpdateFloat(value float32, interval float32) float32 {
	if interval <= 0 {
		interval = 0.1 * value
	}
	minValue := value - interval
	maxValue := value + interval
	return c.NextFloat(minValue, maxValue)
}
