class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = dict()
        i, j = -1, -1

        for k, n in enumerate(nums):
            complement = target-n

            if complement in seen:
                seen_idx = seen[complement]
                if seen_idx<k:
                    i, j = seen_idx, k
                else:
                    i, j = k, seen_idx
                break
            else:
                seen[n] = k if n not in seen else min(k, seen[n])

        return [i, j]