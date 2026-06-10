func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, len(temperatures))
	result := make([]int, len(temperatures))

	// initialize the stack
	stack[0] = 0
	i := 1

	for j, temperature := range temperatures[1:] {
		// fmt.Println(j, temperature, temperatures[stack[i-1]])
		for {
			if i <= 0 || temperature <= temperatures[stack[i-1]] {
				stack[i] = j+1
				i += 1
				break
			}

			result[stack[i-1]] = j + 1 - stack[i-1]
			i -= 1
		}
	}
	return result
}
