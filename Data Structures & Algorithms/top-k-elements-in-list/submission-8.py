from collections import defaultdict

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:

        frequency = defaultdict(lambda: 0)
        for n in nums:
            frequency[n] += 1

        return sorted(frequency, key=frequency.get, reverse=True)[:k]
        