func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var middle int

	for left <= right {
		middle = left + (right-left)/2
		if nums[middle]==target {
			return middle
		}

		if nums[middle] < target {
			left = middle+1
		} else {
			right = middle-1
		}
	}
	return -1
}