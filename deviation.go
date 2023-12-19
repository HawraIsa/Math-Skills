package mathskills

import (
	"math"
)

func CalculateStandardDeviation(variance int) int {
	// calculate the standard deviation for a certain variance
	return int(math.Sqrt(float64(variance)))
}