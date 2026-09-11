class Solution:
	def twoSum(self, nums: List[int], target: int) -> List[int]:
		seen = dict()

		for i, n in enumerate(nums):
			complement = target - n

			if complement in seen:
				return [seen[complement], i]
			else:
				seen[n] = i

		return [-1, -1]
