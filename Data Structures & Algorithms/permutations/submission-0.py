class Solution:
    def __init__(self):
        self.resp = []

    def genPermute(self, N: int, chosen: List[bool], candidates: List[int], nums: List[int]):
        if len(candidates)==N:
            self.resp.append(list(candidates))
            return
        for i in range(N):
            if chosen[i]:
                continue
            candidates.append(nums[i])
            chosen[i] = True
            self.genPermute(N, chosen, candidates, nums)

            candidates.pop()
            chosen[i] = False

    def permute(self, nums: List[int]) -> List[List[int]]:
        N = len(nums)
        if N==1:
            return [nums]

        self.genPermute(N, [False]*N, [], nums)
        return self.resp

        