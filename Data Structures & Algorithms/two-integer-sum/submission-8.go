func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	var complement, complement_idx int
	var i, j int
	var ok bool

	for k, n := range nums {
		complement = target-n

		if complement_idx, ok = seen[complement]; ok {
			if complement_idx <= k {
				i, j = complement_idx, k
			} else {
				i, j = k, complement_idx
			}
			break
		} else {
			seen[n] = k
		}
	}

	return []int{i, j}
}
