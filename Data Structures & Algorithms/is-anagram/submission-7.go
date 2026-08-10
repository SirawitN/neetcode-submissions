func isAnagram(s string, t string) bool {
	if len(s)!=len(t) {
		return false
	}

	charSet1 := make(map[rune]int)

	for _, r := range s {
		charSet1[r] += 1
	}

	for _, r := range t {
		count, ok := charSet1[r]

		if !ok || count<=0 {
			return false
		} else {
			charSet1[r] -= 1
		}
	}

	return true
}
