package count_negatives

func countNegatives(grid [][]int) int {
	M := len(grid)
	count := 0
	for i := range grid {
		N := len(grid[M-1])
		for j := range grid[i] {
			if grid[i][N-1] >= 0 {
				break
			} else {
				count++
				N--
			}
			j++
		}
	}
	return count
}
