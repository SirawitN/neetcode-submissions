class Solution:
    def isPermutation(self, charFreq1: list, charFreq2: list) -> bool:
        for i in range(26):
            if charFreq1[i]!=charFreq2[i]:
                return False
        return True

    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False
    
        charFreqS1 = [0] * 26
        charFreqS2 = [0] * 26
        order_a = ord('a')

        for i in range(len(s1)):
            charFreqS1[ord(s1[i]) - order_a] += 1
            
        startWind = 0
        stopWind = min(startWind + len(s1) - 1, len(s2)-1)
        for i in range(startWind, stopWind):
            charFreqS2[ord(s2[i]) - order_a] += 1

        while stopWind<len(s2):
            charFreqS2[ord(s2[stopWind]) - order_a] += 1
            if self.isPermutation(charFreqS1, charFreqS2):
                return True

            charFreqS2[ord(s2[startWind]) - order_a] -= 1
            startWind += 1
            stopWind += 1
        
        return False