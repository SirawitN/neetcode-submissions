func Merge(arr []int, start int, mid int, stop int) {
	merged := make([]int, stop-start+1)
	i, j := start, mid

	for k:=0; k<stop-start+1; k++ {
		if j>=stop {
			merged[k] = arr[i]
			i++
		} else if i>=mid {
			merged[k] = arr[j]
			j++
		} else {
			if arr[i]<=arr[j] {
				merged[k] = arr[i]
				i++
			} else {
				merged[k] = arr[j]
				j++
			}
		}
	}
	
	copy(arr[start:stop], merged)
	return
}

func DivideSort(nums []int, start int, stop int) {
	if stop-start <= 1 {
		return
	}

	mid := (stop+start)/2
	DivideSort(nums, start, mid)
	DivideSort(nums, mid, stop)
	Merge(nums, start, mid, stop)
	return
}

func sortArray(nums []int) []int {
    DivideSort(nums, 0, len(nums))
	return nums
}
