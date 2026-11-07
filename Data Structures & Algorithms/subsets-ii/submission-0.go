import (
	"slices"
)

func subsetsWithDup(nums []int) [][]int {
	resp := [][]int{}
	memoization := make(map[string]bool)

	var createSubsets func(int, []int, []int, int)
	createSubsets = func(curIdx int, subset []int, nums []int, N int) {
		key := fmt.Sprint(subset, curIdx)
		if _, ok := memoization[key]; !ok && curIdx>=N {
			tmp := make([]int, len(subset))
			copy(tmp, subset)
			resp = append(resp, tmp)
			memoization[key] = true
			return 
		}

		if _, ok := memoization[key]; ok {
			return
		} else {
			memoization[key] = true
		}
		// fmt.Println(key, curIdx)

		added := append(subset, nums[curIdx])
		// memoization[fmt.Sprint(added)] = true
		createSubsets(curIdx+1, added, nums, N)

		// memoization[key] = true
		createSubsets(curIdx+1, subset, nums, N)
	}
	slices.Sort(nums)
	createSubsets(0, []int{}, nums, len(nums))
	return resp
}
