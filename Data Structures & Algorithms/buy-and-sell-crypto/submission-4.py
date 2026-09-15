class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        highestProfit = 0
        boughtPrice = prices[0]

        for currentPrice in prices:
            boughtPrice = currentPrice if currentPrice < boughtPrice else boughtPrice

            if currentPrice > boughtPrice:
                highestProfit = max(highestProfit, currentPrice-boughtPrice)

        return highestProfit