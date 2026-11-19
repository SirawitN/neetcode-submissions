class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack = []
        stack_idx = -1
        N = len(temperatures)
        ans = [0 for _ in range(N)]

        for i, temp in enumerate(temperatures):
            while len(stack)>0 and temp>stack[stack_idx][1]:
                j, tos_temp = stack[stack_idx]
                ans[j] = i-j
                stack.pop()
                stack_idx -= 1
            stack.append((i, temp))
            stack_idx += 1

        return ans
                
