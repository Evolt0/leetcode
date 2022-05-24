package length_of_longest_substring

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	var last [128]int
	for i := 0; i < len(last); i++ {
		last[i] = -1
	}
	var left int = -1
	var ans int = 0
	for i := 0; i < len(s); i++ {
		var si = int(s[i])
		left = Max(left, last[si])
		last[si] = i
		ans = Max(ans, i-left)
	}
	return ans
}
