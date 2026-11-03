class Solution:
    def __init__(self):
        self.resp = []

    def create_subsets(self, i: int, nums: List[int], arr: List[int]):
        if i==len(nums):
            self.resp.append(list(arr))
            return

        arr.append(nums[i])
        self.create_subsets(i+1, nums, arr)

        arr.pop()
        self.create_subsets(i+1, nums, arr)

    def subsets(self, nums: List[int]) -> List[List[int]]:
        self.create_subsets(0, nums, [])

        return self.resp