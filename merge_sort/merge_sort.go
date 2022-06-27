package merge_sort

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

func mergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	middle := len(list) / 2
	left := list[:middle]
	right := list[middle:]
	return merge(mergeSort(left), mergeSort(right))
}
