from collections import defaultdict

class Solution:
	def twoSum(self, nums: List[int], target: int) -> List[int]:
		valueIdxBuff = defaultdict(list)

		for i, n in enumerate(nums):
			valueIdxBuff[n] += [i]

		for i, n in enumerate(nums):
			t = target - n
			if t in valueIdxBuff:
				for j in valueIdxBuff[t]:
					if i!=j:
						return [i, j]
		return [-1, -1]

        