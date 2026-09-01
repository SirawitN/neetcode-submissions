func topKFrequent(nums []int, k int) []int {
    elementFreq := make(map[int]int)

    for _, n := range nums {
        elementFreq[n] += 1
    }

    
    buckets := make([][]int, len(nums)+1)
    for n, count := range elementFreq{
        buckets[count] = append(buckets[count], n)
    }

    result := make([]int, 0, k)
    OUTER:
    for i:=len(buckets)-1; i>=0&&len(result)<k; i-- {
        for _, n := range buckets[i] {
            if len(result)>=k {
                break OUTER
            }
            result = append(result, n)
        }
    }
    return result
}
