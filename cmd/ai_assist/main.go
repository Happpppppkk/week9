package main

import (
	"fmt"
	"time"

	"github.com/montanaflynn/stats"
)

// input data
var x1 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
var x2 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
var x3 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
var x4 = []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
var y1 = []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
var y2 = []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74}
var y3 = []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}
var y4 = []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}

// create main function
func main() {
	data := []stats.Series{
		createSeries(x1, y1),
		createSeries(x1, y1),
		createSeries(x2, y2),
		createSeries(x3, y3),
		createSeries(x4, y4),
	}

	iterations := 500

	totalPT := time.Duration(0)

	for i := 0; i < iterations; i++ {
		startTime := time.Now()
		for _, d := range data {
			OLS(d)
		}
		totalPT += time.Since(startTime)
	}
	averagePT := totalPT / time.Duration(iterations)
	fmt.Println("Average processing time: ", averagePT)
}
func OLS(data stats.Series) (float64, float64) {
	regression, _ := stats.LinearRegression(data)

	maxX := regression[0].X
	maxY := regression[0].Y
	minX := regression[0].X
	minY := regression[0].Y

	for _, point := range regression {
		if point.X > maxX {
			maxX = point.X
		}
		if point.X < minX {
			minX = point.X
		}
		if point.Y > maxY {
			maxY = point.Y
		}
		if point.Y < minY {
			minY = point.Y
		}
	}

	slope := (maxY - minY) / (maxX - minX)
	intercept := minY - slope*minX
	//R, _ := stats.Correlation(x, y)
	fmt.Printf("Slope: %.4f\n", slope)
	fmt.Printf("Intercept: %.4f\n", intercept)
	//fmt.Println(regression)
	return slope, intercept
}

func createSeries(x, y []float64) stats.Series {
	var data stats.Series
	for i := 0; i < len(x); i++ {
		data = append(data, stats.Coordinate{X: x[i], Y: y[i]})
	}
	if len(x) != len(y) {
		panic("check yer data")
	}

	return data
}
