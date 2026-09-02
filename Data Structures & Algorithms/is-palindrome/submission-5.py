class Solution:
    def isPalindrome(self, s: str) -> bool:
        i, j = 0, len(s)-1

        while i<len(s) and j>=0 and i<j:
            if not s[i].isalnum():
                i += 1
                continue
            if not s[j].isalnum():
                j -= 1
                continue

            if s[i].lower() != s[j].lower():
                return False
        
            i, j = i+1, j-1

        return True