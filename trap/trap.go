package trap

func trap(height []int) int {
	ans := 0
	for i := 1; i < len(height)-1; i++ {
		maxLeft, maxRight := 0, 0
		for j := i; j >= 0; j-- {
			maxLeft = Max(maxLeft, height[j])
		}
		for j := i; j < len(height); j++ {
			maxRight = Max(maxRight, height[j])
		}
		ans += Min(maxLeft, maxRight) - height[i]
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
