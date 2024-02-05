package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Check if a file path is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your-program.go data.txt")
		return
	}

	// Extract file path from command-line arguments
	filePath := os.Args[1]

	// Read data from the file
	data, err := readDataFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate and print statistics
	printStatistics(data)
}
/*------------------------------------------------------------------------- Read data from file -----------------------------------------------------------*/

func readDataFromFile(filePath string) ([]float64, error) {
	// open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	} // ensure the file is closed after the program finishes
	defer file.Close()
	
	// create an empty slice to store float64 values
	var data []float64
	
	// create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// convert the text from the file to a float64 value
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		} // append the value to the slice
		data = append(data, value)
	}
	// check if any errors occured through scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	// return the slice
	return data, nil
}

/*------------------------------------------------------------------------- Print statistics -----------------------------------------------------------*/
func printStatistics(data []float64) {
	// Calculate average
	average := calculateAverage(data)

	// Calculate median
	median := calculateMedian(data)

	// Calculate variance
	variance := calculateVariance(data, average)

	// Calculate standard deviation
	stdDev := calculateStandardDeviation(variance)

	// Print the results with rounding
	fmt.Printf("Average: %d\n", roundToInt(average))
	fmt.Printf("Median: %d\n", roundToInt(median))
	fmt.Printf("Variance: %d\n", roundToInt(variance))
	fmt.Printf("Standard Deviation: %d\n", roundToInt(stdDev))
}
/*------------------------------------------------------------------------- Calculate average -----------------------------------------------------------*/

func calculateAverage(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value // sum values up
	} // divide sum by total num of elements
	return sum / float64(len(data))
}
/*------------------------------------------------------------------------- Calculate median -----------------------------------------------------------*/

func calculateMedian(data []float64) float64 {
	//  sort the data slice in ascending order. 
	sort.Float64s(data)
	n := len(data)
	// if the length is even
	if n%2 == 0 { //  return the average of the two middle values.
		return (data[n/2-1] + data[n/2]) / 2.0
	}
	// if number is odd, return the middle value
	return data[n/2]
}
/*------------------------------------------------------------------------- Calculate variance -----------------------------------------------------------*/

func calculateVariance(data []float64, average float64) float64 {
	sum := 0.0
	// iterate through slice elements
	for _, value := range data {
		// subtract the average from each element, square the result then sums up the squared differences
		sum += math.Pow(value-average, 2)
	} // divides the sum by the total number of elements
	return sum / float64(len(data))
}
/*------------------------------------------------------------------------- Calculate standard deviation -----------------------------------------------------------*/

func calculateStandardDeviation(variance float64) float64 {
	// return the square root of the variance
	return math.Sqrt(variance)
}
/*------------------------------------------------------------------------- Round to integer -----------------------------------------------------------*/

func roundToInt(value float64) int {
	// round the float64 value to the nearest integer
	return int(math.Round(value))
}
