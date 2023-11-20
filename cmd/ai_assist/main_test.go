package main

import (
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestOLS(t *testing.T) {
	type args struct {
		data stats.Series
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OLS(tt.args.data)
			if got != tt.want {
				t.Errorf("OLS() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("OLS() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_createSeries(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name string
		args args
		want stats.Series
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSeries(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createSeries() = %v, want %v", got, tt.want)
			}
		})
	}
}
