func create_subsets(i int, nums []int, arr []int, resp *[][]int) {
	if i==len(nums) {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*resp = append(*resp, tmp)
		return
	}

	arr_tmp := append(arr, nums[i])
	create_subsets(i+1, nums, arr_tmp, resp)

	create_subsets(i+1, nums, arr, resp)
}

func subsets(nums []int) [][]int {
	resp := [][]int{}
	create_subsets(0, nums, []int{}, &resp)
	return resp
}
