class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        ans = []
        stack = [([], -1)]
        N = len(nums)

        while len(stack)>0:
            tos_arr, tos_idx = stack.pop()
            tos_idx += 1
            if tos_idx==N:
                ans.append(tos_arr)
                continue
    
            stack.append((tos_arr+[nums[tos_idx]], tos_idx))
            stack.append((tos_arr, tos_idx))
    
        return ans

        
