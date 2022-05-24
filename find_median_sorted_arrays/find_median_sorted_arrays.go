package find_median_sorted_arrays

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	switch {
	case len(nums1) == 0:
		return 0
	case len(nums1)%2 == 1:
		return float64(nums1[len(nums1)/2])
	case len(nums1)%2 == 0:
		return float64(nums1[len(nums1)/2]+nums1[(len(nums1)/2)-1]) / 2
	default:
		return 0
	}
}
