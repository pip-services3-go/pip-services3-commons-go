package random

import (
	"math"
	"math/rand"
)

type TRandomLong struct{}

var RandomLong *TRandomLong = &TRandomLong{}

func (c *TRandomLong) NextLong(min int64, max int64) int64 {
	if max-min <= 0 {
		return min
	}

	return min + rand.Int63n(max-min)
}

func (c *TRandomLong) UpdateLong(value int64, interval int64) int64 {
	if interval <= 0 {
		interval = int64(math.Trunc(0.1 * float64(value)))
	}
	minValue := value - interval
	maxValue := value + interval
	return c.NextLong(minValue, maxValue)
}

func (c *TRandomLong) Sequence(min int64, max int64) []int64 {
	if min < 0 {
		min = 0
	}
	if max < min {
		max = min
	}

	count := c.NextLong(min, max)

	result := make([]int64, count, count)
	for i := range result {
		result[i] = int64(i)
	}

	return result
}
