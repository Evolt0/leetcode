package min_path_sum

import (
	"log"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	grid := [][]int{{1,2,3},{4,5,6}}
	log.Println(minPathSum(grid))
}
