func minEatingSpeed(piles []int, h int) int {
	var maxBananas int
	for _, bananas := range piles {
		maxBananas = max(maxBananas, bananas)
	}

	// edge case
	if len(piles)==h {
		return maxBananas
	}

	start := 1
	stop := maxBananas
	var candidateSpeed, testingSpeed int
	for start <= stop {
		var totalEatingHour float64
		testingSpeed = start + (stop-start)/2
		for _, bananas := range piles {
			totalEatingHour += math.Ceil(float64(bananas)/float64(testingSpeed))
		}

		if int(totalEatingHour) <= h {
			candidateSpeed = testingSpeed
			stop = testingSpeed-1
		} else {
			start = testingSpeed+1
		}
	}

	return candidateSpeed
}
