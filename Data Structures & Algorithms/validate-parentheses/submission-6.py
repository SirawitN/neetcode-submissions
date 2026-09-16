class Solution:
    def isValid(self, s: str) -> bool:
        closeBrackets = {
            ")": "(",
            "}": "{",
            "]": "["
        }

        stack = [None] * len(s)
        stackPtr = -1

        for ch in s:
            # check if the current ch is a close bracket
            if ch in closeBrackets:
                if stack[stackPtr] == closeBrackets[ch]:
                    stackPtr -= 1
                    continue
                else:
                    return False
            else:
                stackPtr += 1
                stack[stackPtr] = ch
        
        return stackPtr == -1