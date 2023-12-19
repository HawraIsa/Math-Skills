package mathskills

import (
	"math"
)

func CalculateVariance(data [] int, average int) int {
	// calculate the variance of the given data
	sum := 0.0

	for _, num := range data {
		sum += math.Pow(float64(num - average), 2)
	}
	return int(sum / float64(len(data)))
}