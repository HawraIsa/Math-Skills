package mathskills

import (

)

func CalculateAverage(data [] int) int {
	// calculate the average of a given data set.
	sum := 0

	for _, num := range data {
		sum += num
	}

	return sum / len(data)
}