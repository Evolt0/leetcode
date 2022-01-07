package max_profit

import (
	"log"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	data := []int{7, 1, 5, 3, 6, 4}
	log.Println(maxProfit(data))
}
