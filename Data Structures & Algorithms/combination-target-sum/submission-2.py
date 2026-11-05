class Solution:
    def __init__(self):
        self.resp = []

    def backtrack(self, i: int, chosen: List[int], candidates: List[int], remaining: int):
        if remaining==0:
            self.resp.append(list(chosen))
            return

        if remaining < 0 or i==len(candidates):
            return

        chosen.append(candidates[i])
        self.backtrack(i, chosen, candidates, remaining-candidates[i])

        chosen.pop()
        self.backtrack(i+1, chosen, candidates, remaining)

    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        self.backtrack(0, [], nums, target)
        return self.resp
        