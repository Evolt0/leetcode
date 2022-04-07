package main

import "fmt"

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	ans := ""
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		for k := 0; k <= i; k++ {
			dp[i][k] = s[i] == s[k] && (i-1 < k+1 || dp[i-1][k+1])
			if dp[i][k] && i-k+1 > len(ans) {
				ans = s[k : i+1]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestPalindrome2("adacad"))
}

func longestPalindrome2(s string) string {
	dp := make([][]bool, len(s))
	ans := ""
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		for k := 0; k <= i; k++ {
			if s[i] == s[k] && (i-1 < k+1 || dp[i-1][k+1]) {
				dp[i][k] = true
			} else {
				dp[i][k] = false
			}
			if dp[i][k] && i+1-k > len(ans) {
				ans = s[k : i+1]
			}
		}
	}
	return ans
}
