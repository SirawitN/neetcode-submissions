func twoSum(nums []int, target int) []int {
    seen := make(map[int]int, len(nums))

    for i, num := range nums {
        complement := target - num

        if j, exist := seen[complement]; exist {
            return []int{j, i}
        }
        seen[num] = i
    }

    return nil
}
