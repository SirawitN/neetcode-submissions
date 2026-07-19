func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	var a, b rune

	for i<=j {
		a = unicode.ToLower(rune(s[i]))
		b = unicode.ToLower(rune(s[j]))

		if !unicode.IsDigit(a) && !unicode.IsLetter(a) {
			// fmt.Println('a', a)
			i += 1
			continue
		} 
		if !unicode.IsDigit(b) && !unicode.IsLetter(b) {
			// fmt.Println('b', b)
			j -= 1
			continue
		}

		if a != b {
			return false
		}
		i += 1
		j -= 1
	}
	return true
}
