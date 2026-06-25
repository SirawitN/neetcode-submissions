func findMin(nums []int) int {
	start := 0
	stop := len(nums)-1
	ans := nums[0]
	var mid int

	for start <= stop {
		mid = start + (stop-start)/2
		// fmt.Println(mid, ans)

		if nums[mid] > nums[stop] {
			start = mid+1
		} else {
			if nums[mid]<ans {
				ans = nums[mid]
			}
			stop = mid-1
		}
	}

	return ans
}
