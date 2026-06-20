func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)-1
	col := len(matrix[0])-1

	var start, stop, mid, targetRow int
	// search for a row that potentially contains  the target
	start = 0
	stop = row
	for start<= stop {
		mid = start + (stop-start)/2
		if matrix[mid][0] <= target {
			start = mid+1
			targetRow = mid
		} else {
			stop = mid-1
		}
	}

	// search the row
	start = 0
	stop = col
	for start <= stop {
		mid = start + (stop-start)/2
		if matrix[targetRow][mid]==target {
			return true
		}
		if matrix[targetRow][mid] < target {
			start = mid+1
		} else {
			stop = mid-1
		}
	}
	return false
}
