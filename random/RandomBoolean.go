package random

import (
	"math/rand"
)

type TRandomBoolean struct{}

var RandomBoolean *TRandomBoolean = &TRandomBoolean{}

func (c *TRandomBoolean) Chance(chances int, maxChances int) bool {
	if chances < 0 {
		chances = 0
	}
	if maxChances < 0 {
		maxChances = 0
	}
	if chances == 0 && maxChances == 0 {
		return false
	}
	if maxChances < chances {
		maxChances = chances
	}
	start := (maxChances - chances) / 2
	end := start + chances
	hit := rand.Intn(maxChances + 1)
	return hit >= start && hit <= end
}

func (c *TRandomBoolean) NextBoolean() bool {
	return rand.Float32()*100 < 50
}
