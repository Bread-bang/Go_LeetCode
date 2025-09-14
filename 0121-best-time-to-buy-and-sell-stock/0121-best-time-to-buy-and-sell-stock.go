func maxProfit(prices []int) int {
    minPrice := prices[0]
    maxProfit := 0

    for i := 1; i < len(prices); i++ {
        profit := prices[i] - minPrice
        if profit > maxProfit {
            maxProfit = profit
        }

        if minPrice > prices[i] {
            minPrice = prices[i]
        }
    }

    return maxProfit
}