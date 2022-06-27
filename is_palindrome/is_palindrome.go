package is_palindrome

import (
	"math"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 || x >= math.MaxInt32 || x <= math.MinInt32 {
		return false
	}
	data, re := x, 0
	for data != 0 {
		re = re*10 + data%10
		data = data / 10
	}
	if re == x {
		return true
	}
	return false
}

func isPalindromeV2(x int) bool {
	if x < 0 || x >= math.MaxInt32 || x <= math.MinInt32 {
		return false
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
