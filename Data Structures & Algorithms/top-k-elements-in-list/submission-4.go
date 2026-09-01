import (
    "maps"
    "slices"
)

func topKFrequent(nums []int, k int) []int {
    elementFreq := make(map[int]int, len(nums))

    for _, n := range nums {
        elementFreq[n] += 1
    }

    keys := slices.Collect(maps.Keys(elementFreq))
    slices.SortFunc(keys, func(a, b int) int {
        aFreq, bFreq := elementFreq[a], elementFreq[b]
        if aFreq != bFreq {
            if aFreq > bFreq {
                return -1
            }
            return 1
        }

        if a < b{
            return -1
        } else if a > b{
            return 1
        }
        return 0
    })

    return keys[:k]
}
