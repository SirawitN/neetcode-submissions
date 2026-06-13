func binSearch(nums []int, target int, start, stop int) int {
	if start > stop {
		return start
	}

	mid := start + (stop-start)/2
	if nums[mid]==target {
		return mid
	}

	if nums[mid] > target {
		return binSearch(nums, target, start, mid-1)
	}
	return binSearch(nums, target, mid+1, stop)
}

func searchInsert(nums []int, target int) int {
	return binSearch(nums, target, 0, len(nums)-1)
}
