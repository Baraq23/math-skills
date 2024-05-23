package mathskills

import (
	"testing"
)

func TestAverage(t *testing.T) {
	type test struct {
		Name     string
		Data     []float64
		Expected int
	}
	tests := []test{
		{"Testing positive values", []float64{8, 22, 61}, 30},
		{"Testing negative values", []float64{-8, -22, -61}, -30},
		{"Testing mixed values", []float64{8.2, -22.1, 60}, 15},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			if got := int(Average(tc.Data)); got != tc.Expected {
				t.Errorf("%v : want: %v got: %v", tc.Name, tc.Expected, got)
			}
		})
	}
}
