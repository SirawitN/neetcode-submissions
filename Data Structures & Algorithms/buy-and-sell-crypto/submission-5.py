class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        possibleProfit = 0
        buyPrice = prices[0]

        for price in prices[1:]:
            if price < buyPrice:
                buyPrice = price

            possibleProfit = max(possibleProfit, price-buyPrice)

        return possibleProfit
        