import (
	"cmp"
	"slices"
)

func getFrequency(nums []int) map[int][]int{
	var freq int
	curNum := nums[0]
	count := make(map[int][]int, len(nums))
	for curNumIdx, curIdx := 0, 0; curIdx<=len(nums); curIdx++ {
		if curIdx==len(nums) {
			freq = curIdx - curNumIdx
			count[freq] = append(count[freq], curNum)
			break
		}

		if nums[curIdx]!=curNum {
			freq = curIdx-curNumIdx
			count[freq] = append(count[freq], curNum)
			curNum = nums[curIdx]
			curNumIdx = curIdx
		} else {
			continue
		}
	}

	return count
}

func topKFrequent(nums []int, k int) []int {
	sort.Ints(nums)
	frequency := getFrequency(nums)

	allFrequencies := make([]int, 0, len(frequency))
	for freq := range frequency {
		allFrequencies = append(allFrequencies, freq)
	}
	slices.SortFunc(allFrequencies, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	output := make([]int, 0, len(allFrequencies))
	for _, freq := range allFrequencies {
		if len(output) >= k {
			break
		}
		minEls := min(k-len(output), len(frequency[freq]))
		output = slices.Concat(output, frequency[freq][:minEls])
	}
	return output
}
