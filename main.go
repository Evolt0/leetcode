package main

import (
	"fmt"
	"github.com/Evolt0/leetcode/quick_sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := []int{5, 3, 2, 1, 6, 4, 8, 7}
	sort := quick_sort.QuickSort(s)
	fmt.Println(sort)
}

func reciprocalTopN(list *ListNode, topN int) int {
	top, bottom := list, list

	for i := 0; i < topN; i++ {
		if bottom == nil {
			return 0
		}
		bottom = bottom.Next
	}
	for bottom != nil {
		fmt.Println(bottom.Val)
		top = top.Next
		bottom = bottom.Next
	}
	return top.Val
}

func quickSort(list []int) []int {
	return _quickSort(list, 0, len(list))
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partition := Partition(arr, left, right)
		_quickSort(arr, left, partition-1)
		_quickSort(arr, partition+1, right)
	}
	return arr
}

func Partition(arr []int, left, right int) int {
	pivot, index := left, left+1
	for i := index; i < right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index++
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
