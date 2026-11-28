func search(nums []int, target int) int {
	var bsearch func([]int, int, int, int) int
	bsearch = func(nums []int, target int, start int, stop int) int {
		// terminating condition: we cant generate a subproblem any further
		if stop-start <= 1 {
			if nums[start]==target {
				return start
			} else {
				return -1
			}
		}

		mid := (start+stop)/2
		if nums[mid]>target {
			return bsearch(nums, target, start, mid)
		} else {
			return bsearch(nums, target, mid, stop)
		}
	}

	return bsearch(nums, target, 0, len(nums))
}
