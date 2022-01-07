package fibonacci

func fib(n int) int {
	dp := make([]int, n+1)
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % mod
	}
	return dp[n]
}
