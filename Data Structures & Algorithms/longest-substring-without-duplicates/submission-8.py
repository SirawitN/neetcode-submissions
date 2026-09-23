class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        seen = dict()
        start = 0
        longestSubstring = 0

        for i, ch in enumerate(s):
            if ch in seen:
                start = max(seen[ch] + 1, start)
                        
            seen[ch] = i
            longestSubstring = max(longestSubstring, i-start+1)

        return longestSubstring