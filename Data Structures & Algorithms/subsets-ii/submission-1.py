class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        resp = []
        N = len(nums)

        def backtrack(start: int, subset: List[int], nums: List[int], N: int) -> None :
            resp.append(list(subset))

            for j in range(start, N):
                if j>start and nums[j]==nums[j-1]:
                    continue
                subset.append(nums[j])
                backtrack(j+1, subset, nums, N)

                subset.pop()
        
        nums.sort()
        backtrack(0, [], nums, N)
        return resp