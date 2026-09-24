func characterReplacement(s string, k int) int {
	charFreq := make(map[rune]int, 26)
	start := 0
	maxLength, maxFreq := 0, 0

	for stop, r := range s {
		charFreq[r] += 1
		maxFreq = max(maxFreq, charFreq[r])

		if (stop-start+1)-maxFreq > k {
			charFreq[rune(s[start])] -= 1
			start += 1
		}

		maxLength = max(maxLength, stop-start+1)
	}

	return maxLength
}
