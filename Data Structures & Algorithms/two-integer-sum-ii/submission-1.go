func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	var sum int

	for i < j {
		sum = numbers[i]+numbers[j]
		if sum==target {
			break
		}

		if sum>target {
			j-=1
		} else {
			i+=1
		}
	}

	return []int{i+1, j+1}
}
