package candy

func candy(ratings []int) int {
	if len(ratings) < 2 {
		return len(ratings)
	}
	result := make([]int, len(ratings))
	for i := 0; i < len(result); i++ {
		result[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			result[i] = result[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			result[i-1] = max(result[i-1], result[i]+1)
		}
	}
	sum := 0
	for _, value := range result {
		sum += value
	}
	return sum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
