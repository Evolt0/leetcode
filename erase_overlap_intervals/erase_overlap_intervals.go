package main

import "sort"

func main() {

}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 1 {
		return 0
	}
	// 数组排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	prev, result := intervals[0][1], 0
	for i := 1; i < len(intervals); i++ {
		if prev > intervals[i][0] {
			result++
		} else {
			prev = intervals[i][1]
		}
	}
	return result
}
