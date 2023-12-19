package mathskills

import (
	"sort"
)

func CalculateMedian(data [] int) int {
	// calculate the median of a given data set.
	sort.Ints(data)

	middle := len(data) / 2

	if len(data) % 2 == 0 {
		return (data[middle-1] + data[middle]) / 2
	}
	return data[middle]
}