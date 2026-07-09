func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix)-1, len(matrix[0])-1
	var left, right, middle int

	// search for a row that might contain the target
	right = row
	for left<=right {
		middle = left + (right-left)/2

		if matrix[middle][0] > target {
			right = middle-1
		} else if matrix[middle][0] <= target {
			left = middle+1
		}
	}
	potentialRow := max(0, right)

	// search for the target in the row
	left = 0
	right = col
	for left<=right {
		middle = left + (right-left)/2
		if matrix[potentialRow][middle] == target {
			return true
		}

		if matrix[potentialRow][middle]<target {
			left = middle+1
		} else {
			right = middle-1
		}
	}

	return false
}
