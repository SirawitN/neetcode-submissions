from collections import defaultdict

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        frequencyBuff = defaultdict(lambda: 0)
        for n in nums:
            frequencyBuff[n] += 1

        sortedNums = sorted(frequencyBuff, key=lambda k: frequencyBuff[k], reverse=True)

        return sortedNums[:k]
        