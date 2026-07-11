func searchMin(nums []int) int {
	var left, mid, candidateMinIndex int
	right := len(nums)-1
	candidateMin := nums[0]

	for left<=right {
		mid = left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid+1
		} else {
			right = mid-1
			if nums[mid] < candidateMin {
				candidateMin = nums[mid]
				candidateMinIndex = mid
			}
		}
	}

	return candidateMinIndex
}

func binSearch(nums []int, target int) int {
	var left, mid int
	right := len(nums)-1

	for left<=right {
		mid = left + (right-left)/2

		if nums[mid]==target {
			return mid
		}

		if nums[mid]>target {
			right = mid-1
		} else {
			left = mid+1
		}
	}

	return -1
}

func search(nums []int, target int) int {
	// find the index of the minimum value to seperate the rotated list into 2 sorted lists.
	minIndex := searchMin(nums)

	// fmt.Println(nums[:minIndex], nums[minIndex:])
	targetInFirstHalf := binSearch(nums[:minIndex], target)
	if targetInFirstHalf != -1 {
		return targetInFirstHalf
	}

	targetInSecondHalf := binSearch(nums[minIndex:], target)
	if targetInSecondHalf != -1 {
		return minIndex + targetInSecondHalf
	}

	return -1
}
