func hasDuplicate(nums []int) bool {
    hashMap := make(map[int]struct{})

    for _, num := range nums{
        if _, ok := hashMap[num]; ok {
            return true
        } else {
            hashMap[num] = struct{}{}
        }
    }

    return false
}
