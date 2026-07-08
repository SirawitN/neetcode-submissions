func characterReplacement(s string, k int) int {
	left := 0
	maxLength, maxCharFreq := 0, 0
	charFreq := make(map[byte]int)

	for i := range len(s) {
		charFreq[s[i]] += 1
		maxCharFreq = max(maxCharFreq, charFreq[s[i]])

		if i-left+1-maxCharFreq > k {
			charFreq[s[left]] -= 1
			left += 1
		}
		maxLength = max(maxLength, i-left+1)
	}
	return maxLength
}
