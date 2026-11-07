class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        resp = []
        queue = [([], -1)]
        N = len(nums)

        while len(queue)!=0:
            curr, startIdx = queue.pop()
            resp.append(curr)
            if startIdx>=N:
                continue

            for i in range(startIdx+1, N):
                tmp = curr + [nums[i]]
                queue.append((tmp, i))
        return resp