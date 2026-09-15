func maxProfit(prices []int) int {
    boughtPrice := prices[0]
    highestProfit := 0

    for _, currentPrice := range prices[1:] {
        if currentPrice < boughtPrice {
            boughtPrice = currentPrice
        } else if currentPrice > boughtPrice {
            highestProfit = max(highestProfit, currentPrice-boughtPrice)
        }
    }

    return highestProfit
}
