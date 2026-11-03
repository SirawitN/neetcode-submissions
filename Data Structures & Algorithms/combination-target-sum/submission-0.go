func createCombSum(i int, nums []int, arr []int, resp *[][]int, sum int, target int) {
	if i==len(nums) || sum>=target {
		if sum==target {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			*resp = append(*resp, tmp)
		}
		return
	}

	arr_tmp := append(arr, nums[i])
	createCombSum(i, nums, arr_tmp, resp, sum+nums[i], target)

	createCombSum(i+1, nums, arr, resp, sum, target)	
}

func combinationSum(nums []int, target int) [][]int {
    resp := [][]int{}
	createCombSum(0, nums, []int{}, &resp, 0, target)
	return resp
}
