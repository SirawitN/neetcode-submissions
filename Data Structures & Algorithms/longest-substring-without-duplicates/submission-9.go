func lengthOfLongestSubstring(s string) int {
    seen := make(map[rune]int, len(s))
    startIdx := 0
    longestSubstring := 0

    for i, r := range s {
        if seenIdx, ok := seen[r]; ok {
            startIdx = max(startIdx, seenIdx+1)
        }

        seen[r] = i
        longestSubstring = max(longestSubstring, i-startIdx+1)
    }

    return longestSubstring
}
