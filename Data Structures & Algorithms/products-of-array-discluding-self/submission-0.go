func productExceptSelf(nums []int) []int {
	allSum := 1
	var containZero, containMoreOneZero bool
	for _, num := range nums {
		if num != 0 {
			allSum *= num
		} else {
			if !containZero {
				containZero = true
			} else {
				containMoreOneZero = true
			}
		}
	}

	output := make([]int, len(nums))
	for i, num := range nums {
		if containMoreOneZero || (containZero && num!=0) {
			output[i] = 0
		} else if num==0 {
			output[i] = allSum
		} else {
			output[i] = allSum / num
		}
	}

	return output
}
