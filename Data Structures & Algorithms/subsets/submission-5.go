func subsets(nums []int) [][]int {
	resp := [][]int{}
	N := len(nums)

	var createSubset func(int, []int, []int, int)
	createSubset = func(curIdx int, subset []int, nums []int, N int) {
		if curIdx>=N {
			tmp := make([]int, len(subset))
			copy(tmp, subset)
			resp = append(resp, tmp)
			return
		}

		added := append(subset, nums[curIdx])
		createSubset(curIdx+1, added, nums, N)

		createSubset(curIdx+1, subset, nums, N)
	}

	createSubset(0, []int{}, nums, N)
	return resp
}
