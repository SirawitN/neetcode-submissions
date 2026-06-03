func productExceptSelf(nums []int) []int {
	prefixMul := make([]int, len(nums))
	suffixMul := make([]int, len(nums))

	for i := 0; i<len(nums); i++ {
		if i==0 {
			prefixMul[i] = 1
		} else {
			prefixMul[i] = prefixMul[i-1]*nums[i-1]
		}
	}
	for j := len(nums)-1; j>=0; j-- {
		if j==len(nums)-1 {
			suffixMul[j] = 1
		} else {
			suffixMul[j] = suffixMul[j+1]*nums[j+1]
		}
	}

	fmt.Println(prefixMul)
	fmt.Println(suffixMul)
	output := make([]int, len(nums))
	for i := range nums {
		output[i] = prefixMul[i] * suffixMul[i]
	}
	return output
}
