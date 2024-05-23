package mathskills

func Median(nums []float64) float64 {
	var med float64
	// we rearrange the values in ascending order
	nums = SortVal(nums)
	//
	if len(nums)%2 == 0 {
		indx := (len(nums) / 2) - 1
		med = (nums[indx] + nums[indx+1]) / 2
	} else if len(nums)%2 != 0 {
		indx := ((len(nums) / 2) + 1) - 1
		med = nums[indx]
	}
	return med
}

// Function to sort the values using bubble sort method
func SortVal(nums []float64) []float64 {
	for i := 0; i < len(nums); i++ {
		for i := range nums {
			if i != len(nums)-1 {
				if nums[i] > nums[i+1] {
					temp := nums[i]
					nums[i] = nums[i+1]
					nums[i+1] = temp
				}
			}
		}
	}
	return nums
}
