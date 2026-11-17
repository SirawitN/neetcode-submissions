func subsets(nums []int) [][]int {
	ans := [][]int{}

	subset := []int{}
	var dfs func(int)
	dfs = func(i int) {
		if i==len(nums) {
			tmp := make([]int, len(subset))
			copy(tmp, subset)
			ans = append(ans, tmp)
			return
		}

		subset = append(subset, nums[i])
		dfs(i+1)

		subset = subset[:len(subset)-1]
		dfs(i+1)
	}
	dfs(0)
	return ans
}

