package main

import "log"

func convert(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	t := r + r - 2
	c := (n + t - 1) / t * (r - 1)
	dp := make([][]byte, r)
	for i := range dp {
		dp[i] = make([]byte, c)
	}
	j, k := 0, 0
	for i, ch := range s {
		dp[j][k] = byte(ch)
		if i%t < r-1 {
			j++
		} else {
			j--
			k++
		}
	}
	ans := make([]byte, 0, n)
	for _, bytes := range dp {
		for _, b := range bytes {
			if b > 0 {
				ans = append(ans, b)
			}
		}

	}
	return string(ans)
}

func main() {
	log.Println(convert("PAYPALISHIRING", 3))
}
