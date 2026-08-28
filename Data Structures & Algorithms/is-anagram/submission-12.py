from collections import defaultdict

class Solution:
	def isAnagram(self, s: str, t: str) -> bool:
		if len(s)!=len(t):
			return False
		
		buff = defaultdict(lambda: 0)
		for r1 in s:
			buff[r1] += 1
		
		for r2 in t:
			buff[r2] -= 1

			if buff[r2]<0:
				return False
		
		return True

        