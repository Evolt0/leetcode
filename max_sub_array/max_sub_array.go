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
	test := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	println(maxSubArrayDP(test))
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

func maxSubArrayDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	pre1 := nums[0]
	ans, pre2 := pre1, pre1
	for i := 1; i < len(nums); i++ {
		pre1 = Max(pre2+nums[i], nums[i])
		pre2 = pre1
		ans = Max(pre1, ans)
	}
	return ans
}
