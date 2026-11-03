class Solution:
    def __init__(self):
        self.resp = list()

    def createCombSum(self, i: int, nums: List[int], arr: List[int], curSum: int, target: int):
        if curSum>=target or i==len(nums):
            if curSum==target:
                self.resp.append(list(arr))
            return

        arr.append(nums[i])
        self.createCombSum(i, nums, arr, curSum+nums[i], target)

        arr.pop()
        self.createCombSum(i+1, nums, arr, curSum, target)

    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        self.createCombSum(0, nums, [], 0, target)
        return self.resp
        