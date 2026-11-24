func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	i, j := -1, -1

	for k, n := range nums {
		complement := target-n

		if seen_idx, ok := seen[complement]; ok {
			if seen_idx<k {
				i, j = seen_idx, k
			} else {
				i, j = k, seen_idx
			}
			break
		} else {
			if prev_seen_idx, ok := seen[n]; ok {
				seen[n] = min(k, prev_seen_idx)
			} else {
				seen[n] = k
			}
		}
	}

	return []int{i, j}
}
