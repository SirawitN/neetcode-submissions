func isValid(s string) bool {
    closeBrackets := make(map[rune]rune)
    closeBrackets[')'] = '('
    closeBrackets['}'] = '{'
    closeBrackets[']'] = '['

    stack := make([]rune, len(s))
    stackPtr := -1

    for _, r := range s {
        if openBracket, ok := closeBrackets[r]; ok {
            if stackPtr > -1 && stack[stackPtr] == openBracket {
                stackPtr--
                continue
            } else {
                return false
            }
        } else {
            stackPtr++
            stack[stackPtr] = r
        }
    }

    return stackPtr == -1
}
