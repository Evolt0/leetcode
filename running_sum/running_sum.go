package running_sum

func runningSum(nums []int) []int {
	var buf int
	for i := range nums {
		if i == 0 {
			buf = nums[i]
		} else {
			buf = buf + nums[i]
		}
		nums[i] = buf
	}
	return nums
}
