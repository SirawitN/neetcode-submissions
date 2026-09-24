from collections import defaultdict

class Solution:
	def characterReplacement(self, s: str, k: int) -> int:
		charFreq = defaultdict(lambda: 0)
		start = 0
		maxLen, maxFreq = 0, 0

		for stop, ch in enumerate(s):
			charFreq[ch] += 1
			maxFreq = max(maxFreq, charFreq[ch])

			if ((stop-start+1)-maxFreq > k):
				charFreq[s[start]] -= 1
				start += 1

			maxLen = max(maxLen, stop-start+1)

		return maxLen