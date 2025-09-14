func maxProfit(prices []int) int {
    minIdx := 0
    maxIdx := 1
    highestValue := 0
            
    for i := 0; i < len(prices) - 1; i++ {
        if prices[minIdx] > prices[i + 1] {
            minIdx = i + 1
        } else if prices[minIdx] < prices[i + 1] {
            if prices[maxIdx] <= prices[i + 1] {
                maxIdx = i + 1
            } else if maxIdx < minIdx {
                maxIdx = i + 1
            }
        }

        if i + 1 <= maxIdx {
            if prices[maxIdx] - prices[minIdx] > highestValue {
                highestValue = prices[maxIdx] - prices[minIdx]
                fmt.Println("highestValue", minIdx, maxIdx)
            }
        }
    }
    return highestValue
}