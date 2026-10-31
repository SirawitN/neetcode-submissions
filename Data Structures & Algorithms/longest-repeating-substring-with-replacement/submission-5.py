from collections import defaultdict

class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        charFreq = defaultdict(lambda : 0)
        highestFreq = 0
        startIdx = 0
        longestLength = 0

        for i, ch in enumerate(s):
            charFreq[ch] += 1
            
            # check the character's frequency
            highestFreq = max(highestFreq, charFreq[ch])

            # check if the startIdx should be moved or not
            if i-startIdx+1-highestFreq<=k:
                longestLength = max(longestLength, i-startIdx+1)
            else: 
                # move the startIdx and reduce the freq of the char by 1
                charFreq[s[startIdx]] -= 1
                startIdx += 1

        return longestLength