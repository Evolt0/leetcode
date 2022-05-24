package climb_stairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for j := 2; j <= n; j++ {
		dp[j] = dp[j-1] + dp[j-2]
	}
	return dp[n]
}

func climbStairsV2(n int) int {
	if n <= 2 {
		return n
	}
	n1, n2, re := 1, 1, 0
	for i := 2; i <= n; i++ {
		re = n1 + n2
		n1 = n2
		n2 = re
	}
	return re
}
