class Solution:
    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        ans = []
        subset = []
        N = len(nums)

        def dfs(i: int, cur_sum: int):
            if i==N or cur_sum>=target:
                if cur_sum==target:
                    ans.append(subset.copy())
                return

            subset.append(nums[i])
            dfs(i, cur_sum+nums[i])

            subset.pop()
            dfs(i+1, cur_sum)

        dfs(0, 0)
        return ans
                