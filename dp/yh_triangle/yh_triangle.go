package yh_triangle

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	dp[0] = make([]int, 1)
	dp[0][0] = 1
	for i := 1; i < numRows; i++ {
		dp[i] = make([]int, i+1)
		for j := 1; j <= i; j++ {
			dp[i][0] = 1
			if i == j {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			}
		}
	}
	return dp
}

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}
