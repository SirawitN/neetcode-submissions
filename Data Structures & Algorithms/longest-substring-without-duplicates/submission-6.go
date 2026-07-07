func lengthOfLongestSubstring(s string) int {
	longest := 0
	left := 0
	lastSeenBuff := make(map[rune]int)

	for i, r := range s {
		if seenIdx, ok := lastSeenBuff[r]; ok && seenIdx >= left {
			left = seenIdx + 1
		} 
		lastSeenBuff[r] = i
		
		// fmt.Println(i, left, lastSeenBuff[r])
		longest = max(longest, i-left+1)
	}

	return longest
}
