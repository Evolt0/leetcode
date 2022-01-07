package max_profit

func maxProfit(prices []int) int {
	buy := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		buy = min(prices[i], buy)
		res = max(prices[i]-buy, res)
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
