func binSearch(target, start, stop int) int {
	if start >= stop {
		return stop
	}
	if stop-start==1 {
		if stop*stop <= target {
			return stop
		}
		return start
	}

	mid := start + (stop-start)/2
	if mid*mid==target {
		return mid
	}

	if mid*mid > target {
		return binSearch(target, start, mid-1)
	}
	return binSearch(target, mid, stop)
}

func mySqrt(x int) int {
	return binSearch(x, 1, x)
}
