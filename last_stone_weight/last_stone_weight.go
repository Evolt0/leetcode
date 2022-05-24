package last_stone_weight

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	if len(stones) <= 1 {
		return stones[0]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(stones)))
	fmt.Println(stones)
	for {
		stones[0] = stones[0] - stones[1]
		stones[1] = 0
		sort.Sort(sort.Reverse(sort.IntSlice(stones)))
		if stones[1] == 0 {
			break
		}
	}
	return stones[0]
}
