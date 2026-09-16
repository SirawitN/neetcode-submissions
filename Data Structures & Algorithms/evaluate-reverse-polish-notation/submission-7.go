func evalRPN(tokens []string) int {
    type MathOp func(int, int) int

    arithmetic := map[string]MathOp {
        "+": func(a, b int) int {return a+b},
        "-": func(a, b int) int {return a-b},
        "*": func(a, b int) int {return a*b},
        "/": func(a, b int) int {return a/b},
    }

    stack := make([]int, len(tokens))
    stackPtr := -1

    for _, t := range tokens {
        if operation, ok := arithmetic[t]; ok {
            a, b := stack[stackPtr-1], stack[stackPtr]
            stackPtr -= 1
            stack[stackPtr] = operation(a, b)
        } else {
            stackPtr += 1

            stack[stackPtr], _ = strconv.Atoi(t)
        }
    }

    return stack[stackPtr]
}
