class Solution:
    def __init__(self):
        self.resp = list()

    def backtrack(self, start: int, chosen: List[int], candidates: List[int], remaining: int):
        if remaining==0:
            self.resp.append(list(chosen))
            return

        for i in range(start, len(candidates)):
            if candidates[i]>remaining:
                break
            if i>start and candidates[i]==candidates[i-1]:
                continue

            chosen.append(candidates[i])
            self.backtrack(i+1, chosen, candidates, remaining-candidates[i])
            chosen.pop()

    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        self.backtrack(0, [], candidates, target)

        return self.resp