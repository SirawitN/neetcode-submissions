from collections import defaultdict

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freqMap = defaultdict(lambda: 0)
        for n in nums:
            freqMap[n] += 1

        buffer = [[] for _ in range(len(nums)+1)]
        for numb, count in freqMap.items():
            buffer[count].append(numb)

        result = []
        for i in range(len(nums), 0, -1):
            if len(result)>=k:
                break
            for n in buffer[i]:
                if len(result)>=k:
                    break
                result.append(n)
        
        return result
        

        