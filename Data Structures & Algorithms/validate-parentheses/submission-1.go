func isValid(s string) bool {
    parentheses := make(map[rune]rune, 3)
	parentheses[')'] = '('
	parentheses['}'] = '{'
	parentheses[']'] = '['

	runes := []rune(s)
	stack := make([]rune, len(s))
	var stackIdx int
	for _, r := range runes {
		switch r {
			case '(', '{', '[':
				stack[stackIdx] = r
				stackIdx += 1
			default:
				if stackIdx==0 || stack[stackIdx-1] != parentheses[r] {
					return false
				} else {
					stackIdx -= 1
					continue
				}
		}
	}
	return stackIdx==0
}
