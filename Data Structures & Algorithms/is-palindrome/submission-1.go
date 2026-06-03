func isPalindrome(s string) bool {
	var i, j = 0, len(s)-1

	for {
		if i>=j {
			break
		}

		if !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i += 1
			continue
		}

		if !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j -= 1
			continue
		}

		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		} else {
			i += 1
			j -= 1
		}
	}
	return true
}
