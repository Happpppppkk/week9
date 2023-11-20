package main

import (
	"testing"
)

func TestLinearRegression(t *testing.T) {
	// Test cases with known results
	testCases := []struct {
		x         []float64
		y         []float64
		slope     float64
		intercept float64
	}{
		{
			x:         []float64{1, 2, 3, 4, 5},
			y:         []float64{2, 3, 4, 5, 6},
			slope:     1.0,
			intercept: 1.0,
		},
		{
			x:         []float64{1, 2, 3, 4, 5},
			y:         []float64{1, 2, 3, 4, 5},
			slope:     1.0,
			intercept: 0.0,
		},
		// Add more test cases as needed
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			slope, intercept := linearRegression(testCase.x, testCase.y)
			if slope != testCase.slope || intercept != testCase.intercept {
				t.Errorf("LinearRegression(%v, %v) = (%v, %v), expected (%v, %v)",
					testCase.x, testCase.y, slope, intercept, testCase.slope, testCase.intercept)
			}
		})
	}
}
