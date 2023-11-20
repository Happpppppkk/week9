package main

import (
	"fmt"
	"math"
	"time"

	"github.com/montanaflynn/stats"
)

// Define a struct to store the Anscombe data
type AnscombeData struct {
	x1, x2, x3, x4 []float64
	y1, y2, y3, y4 []float64
}

func main() {
	// Define the Anscombe data
	anscombe := AnscombeData{
		x1: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		x2: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		x3: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		x4: []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
		y1: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		y2: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
		y3: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		y4: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
	}

	// Define the number of iterations
	iterations := 500

	// Measure the average execution time
	start := time.Now()
	for i := 0; i < iterations; i++ {
		runAnalysis(anscombe)
	}
	end := time.Now()

	// Calculate the average execution time
	averageTime := float64(end.Sub(start).Nanoseconds()) / float64(iterations) / math.Pow(10, 6) // in milliseconds

	// Print the average execution time
	fmt.Printf("Average execution time over %d iterations: %.6f milliseconds\n", iterations, averageTime)
}

// Function to run the analysis
func runAnalysis(data AnscombeData) {
	// Perform linear regression for each dataset
	slope1, intercept1 := linearRegression(data.x1, data.y1)
	slope2, intercept2 := linearRegression(data.x2, data.y2)
	slope3, intercept3 := linearRegression(data.x3, data.y3)
	slope4, intercept4 := linearRegression(data.x4, data.y4)

	// You can use the calculated slopes and intercepts as needed
	fmt.Printf("Slope and Intercept for Dataset 1: %.6f, %.6f\n", slope1, intercept1)
	fmt.Printf("Slope and Intercept for Dataset 2: %.6f, %.6f\n", slope2, intercept2)
	fmt.Printf("Slope and Intercept for Dataset 3: %.6f, %.6f\n", slope3, intercept3)
	fmt.Printf("Slope and Intercept for Dataset 4: %.6f, %.6f\n", slope4, intercept4)
}

// Function to perform linear regression using github.com/montanaflynn/stats package
// Function to perform linear regression using github.com/montanaflynn/stats package
// Function to perform linear regression using github.com/montanaflynn/stats package
func linearRegression(x, y []float64) (slope, intercept float64) {
	// Perform linear regression using stats package
	series := make(stats.Series, len(x))
	for i := 0; i < len(x); i++ {
		series[i] = stats.Coordinate{X: x[i], Y: y[i]}
	}

	regressions, err := stats.LinearRegression(series)
	if err != nil {
		// Handle the error here
		fmt.Println("Error:", err)
		return 0, 0
	}

	// Extract slope and intercept from the regression result
	slope = (regressions[0].Y - regressions[1].Y) / (regressions[0].X - regressions[1].X)
	intercept = regressions[0].Y - slope*regressions[0].X

	return slope, intercept
}
