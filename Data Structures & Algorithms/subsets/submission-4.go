func subsets(nums []int) [][]int {
	resp := [][]int{}

	var createSubsets func(int, []int, []int, int, *[][]int)
	createSubsets = func(curIdx int, currSubset []int, nums []int, N int, resp *[][]int) {
		if curIdx>=N {
			tmp := make([]int, len(currSubset))
			copy(tmp, currSubset)
			*resp = append(*resp, tmp)
			return
		}
		
		tmp := append(currSubset, nums[curIdx])
		createSubsets(curIdx+1, tmp, nums, N, resp)

		createSubsets(curIdx+1, currSubset, nums, N, resp)
	}

	createSubsets(0, []int{}, nums, len(nums), &resp)
	return resp
}
