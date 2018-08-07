package random

import (
	"math"
	"math/rand"
)

type TRandomInteger struct{}

var RandomInteger *TRandomInteger = &TRandomInteger{}

func (c *TRandomInteger) NextInteger(min int, max int) int {
	if max-min <= 0 {
		return min
	}

	return min + rand.Intn(max-min)
}

func (c *TRandomInteger) UpdateInteger(value int, interval int) int {
	if interval <= 0 {
		interval = int(math.Trunc(0.1 * float64(value)))
	}
	minValue := value - interval
	maxValue := value + interval
	return c.NextInteger(minValue, maxValue)
}

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
