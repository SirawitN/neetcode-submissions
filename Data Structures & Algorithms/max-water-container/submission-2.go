func maxArea(heights []int) int {
	highestAreaSoFar := 0
	i, j := 0, len(heights)-1
	var currentArea int

	for i<j {
		currentArea = min(heights[i], heights[j])*(j-i)
		highestAreaSoFar = max(highestAreaSoFar, currentArea)

		if heights[i]<heights[j] {
			i++
		} else {
			j--
		}
	}

	return highestAreaSoFar
}
