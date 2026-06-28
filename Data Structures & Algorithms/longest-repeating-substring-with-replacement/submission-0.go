func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	var left, max_len, max_freq int

	for right := 0; right<len(s); right+=1 {
		count[s[right]] += 1
		max_freq = max(max_freq, count[s[right]])

		for right-left+1-max_freq > k {
			count[s[left]] -= 1
			left += 1
		}

		max_len = max(max_len, right-left+1)
	}
	return max_len
}
