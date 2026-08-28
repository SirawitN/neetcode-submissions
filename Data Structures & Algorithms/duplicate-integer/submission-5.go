func hasDuplicate(nums []int) bool {
    buff := make(map[int]struct{})

    for _, n := range nums {
        if _, ok := buff[n]; !ok {
            buff[n] = struct{}{}
        } else {
            return true
        }
    }

    return false
}
