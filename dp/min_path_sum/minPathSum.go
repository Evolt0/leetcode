package min_path_sum

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch {
			case i == 0 && j == 0:
				dp[i][j] = grid[i][j]
			case i == 0:
				dp[i][j] = dp[i][j-1] + grid[i][j]
			case j == 0:
				dp[i][j] = dp[i-1][j] + grid[i][j]
			default:
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
