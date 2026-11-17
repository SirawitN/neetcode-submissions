class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        ans = []
        subset = []
        N = len(nums)

        def dfs(startIdx: int):
            ans.append(subset.copy())

            for i in range(startIdx, N):
                if i>startIdx and nums[i]==nums[i-1]:
                    continue

                subset.append(nums[i])
                dfs(i+1)
                subset.pop()
            
        nums.sort()
        dfs(0)

        return ans