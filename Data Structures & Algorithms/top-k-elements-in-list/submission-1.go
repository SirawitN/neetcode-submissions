func getFrequency(nums []int) map[int]int{
	frequency := make(map[int]int, len(nums))

    for _, num := range nums {
        frequency[num] += 1
    }

	return frequency
}

func topKFrequent(nums []int, k int) []int {
	freqList := make([][]int, len(nums)+1)
    freqs := getFrequency(nums)

    var freq int
    for num := range freqs {
        freq = freqs[num]
        freqList[freq] = append(freqList[freq], num)
    }

    var output, numsToConcat []int
    var numsToAppend int
    for freq := len(freqList)-1; freq >= 0; freq-- {
        numsToConcat = freqList[freq]
        numsToAppend = min(k-len(output), len(numsToConcat))
        output = append(output, numsToConcat[:numsToAppend]...)
    }
    return output
}
