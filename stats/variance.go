package mathskills

func Variance(nums []float64) float64 {
	mn := Average(nums)
	varc := (SumSqr(nums, mn)) / float64(len(nums))
	return varc
}

// func gets the deviation, squares them, and returns their sum
func SumSqr(nums []float64, mn float64) float64 {
	var sum float64
	for _, v := range nums {
		div := v - mn
		sum += (div * div)
	}
	return sum
}
