class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        ans = []
        subset = []

        def dfs(i: int):
            if i==len(nums):
                ans.append(subset.copy())
                return
            
            # include the number at i
            subset.append(nums[i])
            dfs(i+1)

            # exclude the number at i
            subset.pop()
            dfs(i+1)

        dfs(0)
        return ans

        
