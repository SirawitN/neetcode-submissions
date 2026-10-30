class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        lastSeen = dict()
        longestLength = 0
        startIdx = 0
        i = 0

        for ch in s:
            if ch in lastSeen:
                startIdx = max(startIdx, lastSeen[ch]+1)
            longestLength = max(longestLength, i-startIdx+1)
            lastSeen[ch] = i
            i += 1

        return longestLength