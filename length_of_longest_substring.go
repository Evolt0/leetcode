package main

import (
	"fmt"
	"sort"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var ans []int
	tag := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				tag = tag - i
				break
			} else {
				tag = j
			}
		}
		if tag != 0 {
			ans = append(ans, tag)
		}
	}
	fmt.Println(ans)
	sort.Sort(sort.Reverse(sort.IntSlice(ans)))
	return ans[0]
}
func TestLengthOfLongestSubString(t *testing.T) {
	t.Log(lengthOfLongestSubstring(" "))
}
