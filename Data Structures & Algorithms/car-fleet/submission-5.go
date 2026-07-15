type Car struct {
	Position int
	Speed int
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([]Car, n)
	for i := range n {
		cars[i] = Car{
			Position: position[i],
			Speed: speed[i],
		}
	}

	// sort the slice in a descending order
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Position < cars[j].Position
	})

	// Calculate precise fractional time to reach destination
	timeToReachDest := make([]float64, n)
	for i := 0; i < n; i++ {
		timeToReachDest[i] = float64(target-cars[i].Position) / float64(cars[i].Speed)
	}

	var fleetCount int
	var maxTimeAhead float64 // Tracks the time of the lead car of the current fleet

	// Iterate backwards (from the car closest to the target down to the furthest)
	for i := n - 1; i >= 0; i-- {
		// If this car takes LONGER than the lead car of the fleet ahead of it,
		// it can never catch up. It forms a brand new fleet.
		if timeToReachDest[i] > maxTimeAhead {
			fleetCount++
			maxTimeAhead = timeToReachDest[i] // This car becomes the new fleet leader
		}
		// If its time is <= maxTimeAhead, it catches up and becomes part of that fleet.
	}

	return fleetCount
}
