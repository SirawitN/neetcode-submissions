func maxProfit(prices []int) int {
	profit := 0
	minPriceSoFar := prices[0]

	for _, price := range prices {
		if price < minPriceSoFar {
			minPriceSoFar = price
		}
		profit = max(profit, price-minPriceSoFar)
	}

	return profit
}
