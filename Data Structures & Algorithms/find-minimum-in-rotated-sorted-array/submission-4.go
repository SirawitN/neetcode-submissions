func findMin(nums []int) int {
	var left, mid, right int

	right = len(nums)-1
	candidate := nums[0]
	for  left <= right {
		mid = left + (right-left)/2

		if nums[mid]>nums[right] {
			left = mid+1
		} else {
			right = mid-1
			candidate = min(candidate, nums[mid])
		}
		// fmt.Println(left, right, nums[mid])
	}
	return candidate
}
