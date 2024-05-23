package mathskills

func Average(nums []float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	mean := sum / float64(len(nums))
	return mean
}
