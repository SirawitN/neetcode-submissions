import operator

class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        arithmetic = {
            "+": operator.add,
            "-": operator.sub,
            "*": operator.mul,
            "/": lambda a, b: int(a / b),
        }
        stack = [None] * len(tokens)
        stackPtr = -1

        for t in tokens:
            if t in arithmetic:
                a, b = stack[stackPtr-1], stack[stackPtr]
                stackPtr -= 1
                stack[stackPtr] = arithmetic[t](a, b)
            else:
                stackPtr += 1
                stack[stackPtr] = int(t)
        
        return stack[stackPtr]