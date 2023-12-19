package mathskills

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseData(content string) []int {
	// parse the content of the file and return a slice of integers representing the data.
	lines := strings.Split(content, "\n")

	var data []int
	
	for _, line := range lines {
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error parsing data. Error : ", err)
			}
			data = append(data, num)
		}
	}
	return data
}
