package random

import (
	"math/rand"
)

//Random generator for boolean values.
//
//Example:
//
//  value1 := RandomBoolean.nextBoolean();    // Possible result: true
//  value2 := RandomBoolean.chance(1,3);      // Possible result: false
type TRandomBoolean struct{}

var RandomBoolean *TRandomBoolean = &TRandomBoolean{}

//Calculates "chance" out of "max chances". Example: 1 chance out of 3 chances (or 33.3%)
//
//  - chance: number  - a chance proportional to maxChances.
//  - maxChances: number - a maximum number of chances
//
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

//Generates a random boolean value.
//
func (c *TRandomBoolean) NextBoolean() bool {
	return rand.Float32()*100 < 50
}
