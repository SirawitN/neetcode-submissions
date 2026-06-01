func twoSum(nums []int, target int) []int {
    indexes := make(map[int][]int, len(nums))
    diff := make(map[int]int, len(nums))
    for i, num := range nums {
        indexes[num] = append(indexes[num], i)
        diff[num] = target - num
    }

    for i, num := range nums {
        targetPair := diff[num]
        if index, ok := indexes[targetPair]; ok {
            for _, j := range index {
                if i != j {
                    return []int{i, j}
                }
            }
        }
    }
    return []int{-1, -1}
}
