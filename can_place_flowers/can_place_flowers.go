package main

import "fmt"

func main() {
	s := []int{0,0,1}
	println(canPlaceFlowersV2(s, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	if length < n {
		return false
	}
	if length == 1 {
		if (flowerbed[0] == 0 && n <= 1) || (flowerbed[0] == 1 && n < 1) {
			return true
		}
		return false
	}

	count := 0
	for i := 0; i < length; i++ {
		if i == 0 && flowerbed[i] == flowerbed[i+1] {
			count++
			flowerbed[i] = 1
		} else if i == length-1 {
			if flowerbed[i] == flowerbed[i-1] {
				count++
				flowerbed[i] = 1
			}
			break
		} else {
			if flowerbed[i] == flowerbed[i+1] && flowerbed[i] == flowerbed[i-1] {
				count++
				flowerbed[i] = 1
			}
		}
	}
	return count >= n
}

func canPlaceFlowersV2(flowerbed []int, n int) bool {
	length := len(flowerbed)
	count := 0
	fmt.Println(length )
	for i := 0; i < length; {
		if flowerbed[i] == 1 {
			i += 2
		} else if i == length-1 || flowerbed[i+1] == 0 {
			count++
			i += 2
		} else {
			i += 3
		}
	}
	fmt.Println(count )
	return count >= n
}
