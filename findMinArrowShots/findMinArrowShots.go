package main

import (
	"sort"
)

func main() {
	i := [][]int{{8, 7}, {1, 4}}
	println(findMinArrowShots(i))
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1
	max := points[0][1]
	for _, value := range points {
		if value[0] > max {
			count++
			max = value[1]
		}
	}
	return count
}
