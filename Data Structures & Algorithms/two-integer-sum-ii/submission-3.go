func twoSum(numbers []int, target int) []int {
	var sum int
	i, j := 0, len(numbers)-1

	for i < j{
		sum = numbers[i]+numbers[j]
		if (sum==target){
			break
		} 
		
		if (sum<target) {
			i += 1
		} else {
			j -= 1
		}
	}

	return []int{i+1, j+1}
}
