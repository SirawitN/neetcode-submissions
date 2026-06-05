func maxArea(heights []int) int {
    var i, j = 0, len(heights)-1
    var currentVol int
    highestVolSoFar := 0

    for {
        if i >= j {
            break
        }

        currentVol = (j-i) * min(heights[i], heights[j])
        highestVolSoFar = max(highestVolSoFar, currentVol)

        if heights[i] < heights[j] {
            i++
        } else if heights[i] > heights[j] {
            j--
        } else {
            i, j = i+1, j-1
        }
    }

    return highestVolSoFar
}
