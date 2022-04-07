package main

func dichotomize(src []int, target int) int {
	l, r := 0, len(src)
	for l < r {
		m := l + (r-l)/2
		if src[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
