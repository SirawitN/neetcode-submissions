from collections import defaultdict

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:

        if (len(s)!=len(t)):
            return False
		
        charSet = defaultdict(lambda: 0)

        for ch in s :
            charSet[ch] += 1

        for ch in t :
            if (ch not in charSet) or (charSet[ch] <= 0):
                return False
            else:
                charSet[ch] -= 1
		
        return True
			