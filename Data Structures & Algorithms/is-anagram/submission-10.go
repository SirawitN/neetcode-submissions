func isAnagram(s string, t string) bool {
	if len(s)!=len(t) {
		return false
	}

	buff := make(map[rune]int)
	for _, r1 := range(s) {
		buff[r1] += 1
	}

	for _, r2 := range(t) {
		buff[r2] -= 1
		if buff[r2] < 0 {
			return false
		}
	}
	return true
}
