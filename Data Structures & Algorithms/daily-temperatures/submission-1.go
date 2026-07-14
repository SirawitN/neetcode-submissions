func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, len(temperatures)+1)
	stackIdx := 0

	for i, t := range temperatures {
		for i>0 && stackIdx>0 && t > temperatures[stack[stackIdx]] {
			result[stack[stackIdx]] = i-stack[stackIdx]
			stackIdx -= 1
			// fmt.Println(i, stackIdx, result)
		}

		stackIdx += 1
		stack[stackIdx] = i
	}

	return result
}
