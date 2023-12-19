package main

import (
	"fmt"
	"mathskills"
	"os"
)

func main() {

	// check if file path is provided in arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path.")
		return
	}

	// read file
	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file. Error : ", err)
	}

	// parse data
	data := mathskills.ParseData(string(content))

	// calculate statistics
	average := mathskills.CalculateAverage(data)
	median := mathskills.CalculateMedian(data)
	variance := mathskills.CalculateVariance(data, average)
	standardDeviation := mathskills.CalculateStandardDeviation(variance)

	// print results
	fmt.Printf("Average: %d\n", average)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", variance)
	fmt.Printf("Standard Deviation: %d\n", standardDeviation)
}
