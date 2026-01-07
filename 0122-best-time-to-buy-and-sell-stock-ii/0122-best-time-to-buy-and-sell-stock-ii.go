func maxProfit(prices []int) int {
    l := len(prices)
	diff := make([]int, l)

	sum := 0
	diff[0] = 0
	for i := 1; i < l; i++ {
		diff[i] = prices[i] - prices[i-1]
		if diff[i] < 0 {
			diff[i] = 0
		}

		sum += diff[i]
	}

	return sum
}