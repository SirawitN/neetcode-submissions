class Solution:
    def twoSum(self, nums: List[int], start: int, stop: int, target: int) -> List[List[int]]:
        seen = dict()
        partial_ans = []

        for i in range(start, stop):
            n = nums[i]
            complement = target-n
            
            if complement in seen and [-target, complement, n] not in partial_ans:
                partial_ans.append([-target, complement, n])
            else:
                seen[n] = i
        
        return partial_ans

    def threeSum(self, nums: List[int]) -> List[List[int]]:
        N = len(nums)
        nums.sort()
        ans = []

        for idx, num in enumerate(nums):
            if idx>0 and nums[idx-1]==num:
                continue

            sub_ans = self.twoSum(nums, idx+1, N, -num)
            if len(sub_ans)>0:
                ans.extend(sub_ans)
        return ans