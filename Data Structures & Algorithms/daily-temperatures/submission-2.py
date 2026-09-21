class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack = [None] * len(temperatures)
        tos = -1
        ans = [0] * len(temperatures)

        for i, currentTemp in enumerate(temperatures):
            while tos!=-1 and currentTemp>stack[tos][1]:
                tosIdx, _ = stack[tos]
                ans[tosIdx] = i-tosIdx
                tos -= 1

            tos += 1
            stack[tos] = (i, currentTemp)

        return ans
        