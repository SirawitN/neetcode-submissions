func twoSum(nums []int, target int) []int {
    valueIdxBuff := make(map[int][]int, len(nums))
	for i, n := range nums {
		valueIdxBuff[n] = append(valueIdxBuff[n], i)
	}

	for i, n := range nums{
		t := target - n

		if idxs, ok := valueIdxBuff[t]; ok {
			for _, j := range idxs {
				if i!=j {
					return []int{i, j}
				}
			}
		}
	}

	return []int{-1, -1}
}
