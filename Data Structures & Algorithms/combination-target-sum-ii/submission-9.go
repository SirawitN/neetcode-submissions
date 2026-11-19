import (
	"slices"
)

func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	N := len(candidates)

	var dfs func(int, int, []int) 
	dfs = func(start_idx int, cur_sum int, cur_candidate []int) {
		if start_idx==N || cur_sum>=target {
			if cur_sum==target {
				tmp := make([]int, len(cur_candidate))
				copy(tmp, cur_candidate)
				ans = append(ans, tmp)
			}
			return
		}

		for j:=start_idx; j<N; j++ {
			if j>start_idx && candidates[j]==candidates[j-1] {
				continue
			}

			cur_candidate = append(cur_candidate, candidates[j])
			dfs(j+1, cur_sum+candidates[j], cur_candidate)
			cur_candidate = cur_candidate[:len(cur_candidate)-1]
		}
	}

	slices.Sort(candidates)
	dfs(0, 0, []int{})
	return ans
}
