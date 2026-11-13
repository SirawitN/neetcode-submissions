func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	var tmp_sum int

	for i<j {
		tmp_sum = numbers[i] + numbers[j]

		if tmp_sum==target {
			break
		} else if tmp_sum<target {
			i++
		} else {
			j--
		}
	}
	return []int{i+1, j+1}
}
