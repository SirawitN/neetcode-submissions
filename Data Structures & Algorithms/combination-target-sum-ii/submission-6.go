import (
	"slices"
)

func combinationSum2(candidates []int, target int) [][]int {
	resp := [][]int{}
	slices.Sort(candidates)

	var backtrack func(int, []int, []int, int)
	backtrack = func(start int, chosen []int, candidates []int, remaining int) {
		if remaining==0 {
			tmp := make([]int, len(chosen))
			copy(tmp, chosen)
			resp = append(resp, tmp)
			return
		}

		for i:=start; i<len(candidates); i++ {
			if candidates[i]>remaining {
				break
			}

			if i>start && candidates[i]==candidates[i-1] {
				continue
			}

			chosen = append(chosen, candidates[i])
			backtrack(i+1, chosen, candidates, remaining-candidates[i])

			chosen = chosen[:len(chosen)-1]
		}
	}

	backtrack(0, []int{}, candidates, target)
	return resp
}
