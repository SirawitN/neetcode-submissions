func isAnagram(s string, t string) bool {
	if len(s)!=len(t) {
		return false
	}
	
	char_count := make(map[rune]int)

	for _, ch := range s {
		char_count[ch] += 1
	}

	for _, ch := range t {
		char_count[ch] -= 1

		if char_count[ch]<0 {
			return false
		}
	}

	return true
}
