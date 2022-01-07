package rob

import (
	"log"
	"testing"
)

func TestRob(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	rob := Rob(nums)
	log.Println(rob)
}
