func hasDuplicate(nums []int) bool {
    seen := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		} else {
			seen[n] = struct{}{}
		}
	}

	return false
}
