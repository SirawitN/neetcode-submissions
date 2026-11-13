import (
	"slices"
)

func twoSum(nums []int, start int, stop int, target int) [][]int {
	partial_ans := [][]int{}
	seen := make(map[int]struct{})
	duplicated_ans := make(map[[3]int]struct{})
	var n, complement int

	for i:=start; i<stop; i++ {
		n = nums[i]
		complement = target - n

		if _, ok := seen[complement]; ok {
			ans := [3]int{-1*target, complement, n}
			if _, dup := duplicated_ans[ans]; !dup {
				duplicated_ans[ans] = struct{}{}
				partial_ans = append(partial_ans, ans[:])
			}
		} else {
			seen[n] = struct{}{}
		}
	}

	return partial_ans
}

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	slices.Sort(nums)
	N := len(nums)

	for i, n := range nums {
		if i>0 && nums[i]==nums[i-1] {
			continue
		}

		sub_ans := twoSum(nums, i+1, N, -1*n)
		ans = append(ans, sub_ans...)
	}
	return ans
}
