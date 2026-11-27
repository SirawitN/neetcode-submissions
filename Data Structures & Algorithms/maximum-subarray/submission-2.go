func findCrossMaxSubArray(nums []int, start int, mid int, stop int) int {
	maxLeft := nums[mid-1]
	sum := nums[mid-1]
	for i:=mid-2; i>=start; i-- {
		sum += nums[i]
		maxLeft = max(maxLeft, sum)
	}

	maxRight := nums[mid]
	sum = nums[mid]
	for j:=mid+1; j<stop; j++ {
		sum += nums[j]
		maxRight = max(maxRight, sum)
	}
	return maxLeft+maxRight
}

func findMaxSubArray(nums []int, start int, stop int) int {
	// fmt.Println(start, stop)
	if stop-start<=1 {
		return nums[start]
	}

	mid := (start+stop)/2
	maxLeft := findMaxSubArray(nums, start, mid)
	maxRight := findMaxSubArray(nums, mid, stop)
	maxCross :=  findCrossMaxSubArray(nums, start, mid, stop)

	if maxLeft >= maxRight && maxLeft >= maxCross {
		return maxLeft
	} else if maxRight >= maxLeft && maxRight >= maxCross {
		return maxRight
	}

	return maxCross
}

func maxSubArray(nums []int) int {
    return findMaxSubArray(nums, 0, len(nums))
}
