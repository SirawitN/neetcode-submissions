func productExceptSelf(nums []int) []int {
	allSum := 1
	var containZero int
	for _, num := range nums {
		if num != 0 {
			allSum *= num
		} else {
			containZero += 1
		}
	}

	output := make([]int, len(nums))
	for i, num := range nums {
		if containZero > 1 || (containZero>0 && num!=0) {
			output[i] = 0
		} else if num==0 {
			output[i] = allSum
		} else {
			output[i] = allSum / num
		}
	}

	return output
}
