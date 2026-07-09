func sumHour(piles []int, k int) int {
	sum := 0
	for _, bananas := range piles {
		// Ceiling integer division formula: (a + b - 1) / b
		sum += (bananas+k-1)/k
	}
	return sum
}

func minEatingSpeed(piles []int, h int) int {
	minK, maxK := 1, piles[0]
	for _, bananas := range(piles[1:]) {
		maxK = max(maxK, bananas)
	}

	if len(piles)==h {
		return maxK
	}

	var midK, candidateK, hours int
	for minK<=maxK {
		midK = minK + (maxK-minK)/2
		hours = sumHour(piles, midK)

		if hours>h {
			minK=midK+1
		} else {
			candidateK = midK
			maxK = midK-1
		}
	}
	return candidateK
}
