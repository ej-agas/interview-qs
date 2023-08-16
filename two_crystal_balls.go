package interview_qs

import (
	"fmt"
	"math"
)

// TwoCrystalBalls
// Given two crystal balls that will break if dropped from high enough distance,
// determine the exact spot in which it will break in the most optimized way.
// The time complexity of the function is O(sqrt(N))
func TwoCrystalBalls(breaks []bool) string {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount

	// Jump by the square root of N
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	// Jump back by the square root of N
	i -= jumpAmount

	// Walk forward by the square root of N
	// i <= len(breaks) prevents jumping out of the slice
	for j := 0; j <= jumpAmount && i <= len(breaks)-1; j, i = j+1, i+1 {
		if breaks[i] {
			return fmt.Sprintf("Crystal Balls break at index %d.", i)
		}
	}

	return fmt.Sprintf("Crystal balls did not break at any point.")
}
