func lengthOfLongestSubstring(s string) int {
	chars := make(map[rune]int)
	var longestLengthSoFar, left int

	for currentIndex, ch := range s {
		if lastSeenIndex, ok := chars[ch]; ok && lastSeenIndex >= left {
			left = lastSeenIndex + 1
		}
		
		chars[ch] = currentIndex
		longestLengthSoFar = max(longestLengthSoFar, currentIndex-left+1)
		// fmt.Println(currentIndex, left)
	}

	return longestLengthSoFar
}
