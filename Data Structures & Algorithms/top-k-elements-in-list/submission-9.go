import (
	"cmp"
	"slices"
	"maps"
)

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)

	for _, n := range nums {
		frequency[n] += 1
	}

	keys := slices.Collect(maps.Keys(frequency))
	slices.SortFunc(keys, func(a int, b int) int {
		return cmp.Compare(frequency[b], frequency[a])
	})
	return keys[:k]
}
