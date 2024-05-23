package mathskills

import "math"

func SD(nums []float64) float64 {
	variance := Variance(nums)
	sd := math.Sqrt(variance)
	return sd
}
