func maxArea(heights []int) int {
    maxVol := 0
    i, j := 0, len(heights)-1

    for i < j {
        maxVol = max(maxVol, (j-i)*min(heights[i], heights[j]))

        if heights[i]<=heights[j] {
            i += 1
        } else {
            j -= 1
        }
    }

    return maxVol
}
