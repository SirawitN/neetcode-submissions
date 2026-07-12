func isValid(s string) bool {
    stack := make([]rune, len(s))
	stackIdx := -1

	parentheses := map[rune]rune {
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range s {
		if p, ok := parentheses[c]; ok {
			if stackIdx>=0 && stack[stackIdx]==p {
				stackIdx -= 1
			} else {
				// fmt.Println(p, stack[stackIdx], stack)
				return false
			}
		} else {
			stackIdx += 1
			stack[stackIdx] = c
		}
	}

	// fmt.Println(stackIdx)
	return stackIdx == -1
}
