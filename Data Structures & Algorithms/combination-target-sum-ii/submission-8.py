class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        ans = []
        N = len(candidates)

        def dfs(startIdx: int, cur_sum: int, cur_candidate: List[int]):
            if startIdx==N or cur_sum >= target:
                if cur_sum==target:
                    ans.append(cur_candidate.copy())
                return

            for j in range(startIdx, N):
                if j>startIdx and candidates[j]==candidates[j-1]:
                    continue

                cur_candidate.append(candidates[j])
                dfs(j+1, cur_sum+candidates[j], cur_candidate)
                cur_candidate.pop()

        candidates.sort()
        dfs(0, 0, [])
        return ans