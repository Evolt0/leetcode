package main

import (
	"testing"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}
func TestFindMedianSortedArrays(t *testing.T) {
	num1 := []int{1}
	num2 := []int{}
	println(findMedianSortedArrays(num1, num2))
}
