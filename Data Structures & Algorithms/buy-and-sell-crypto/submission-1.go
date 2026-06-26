func maxProfit(prices []int) int {
	minPriceSoFar := prices[0]
	maxProfit := 0

	for _, price := range prices[1:] {
		if price > minPriceSoFar {
			maxProfit = max(maxProfit, price-minPriceSoFar)
		} else if price < minPriceSoFar {
			minPriceSoFar = price
		}
	}
	return maxProfit
}
