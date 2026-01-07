func maxProfit(prices []int) int {
    l := len(prices)

	sum := 0
	for i := 1; i < l; i++ {
        if prices[i] > prices[i-1] {
            sum += prices[i] - prices[i-1]
        }
	}

	return sum
}