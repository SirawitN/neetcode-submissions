class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = dict()

        for k, n in enumerate(nums):
            complement = target-n

            if complement in seen:
                complement_idx = seen[complement]
                if complement_idx <= k:
                    i, j = complement_idx, k
                else:
                    i, j = k, complement_idx

                break
            else:
                seen[n] = k
        
        return [i, j]