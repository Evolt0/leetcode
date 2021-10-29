package main

func maxSubArray(nums []int) int {
	//dp := make([][]int, len(nums))

	max := nums[0]

	for i := 0; i < len(nums); i++ {
		mid := nums[i]
		if mid > max {
			max = mid
		}
		for j := i + 1; j < len(nums); j++ {
			mid += nums[j]
			if mid > max {
				max = mid
			}
		}
	}
	return max
}

func main() {
	test := []int{-2, 1}
	println(maxSubArrayV2(test))
}

func maxSubArrayV2(nums []int) int {
	max := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		max = Max(sum, max)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
