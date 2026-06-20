func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	
	start := 1
	stop := x
	var mid, ans int

	for start<=stop {
		mid = start + (stop-start)/2
		if mid == x/mid {
			return mid
		}

		if mid < x/mid {
			start = mid+1
			ans = mid
		} else {
			stop = mid-1
		}
	}

	return ans
}
