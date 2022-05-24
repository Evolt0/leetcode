package main

import (
	"log"
	"math"
)

func reverse(x int) (rev int) {
	for x != 0 {
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
		if rev < math.MinInt32 || rev > math.MaxInt32 {
			return 0
		}
	}
	return
}

func main() {
	log.Println(1 << 31)
}