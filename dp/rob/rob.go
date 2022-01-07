package rob

func Rob(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = Max(dp[i-1], nums[i-1]+dp[i-2])
	}
	return dp[len(nums)]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
