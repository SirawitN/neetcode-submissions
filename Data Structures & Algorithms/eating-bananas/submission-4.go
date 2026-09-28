import (
    "slices"
)

func minEatingSpeed(piles []int, h int) int {
    minEatSpeed := 1
    maxEatSpeed := slices.Max(piles)

    if len(piles)==h {
        return maxEatSpeed
    }

    var currentSpeed int
    var eatHours int
    possibleSpeed := maxEatSpeed

    for minEatSpeed<=maxEatSpeed {
        // fmt.Println(minEatSpeed, maxEatSpeed)
        eatHours = 0
        currentSpeed = (minEatSpeed + maxEatSpeed) / 2

        for _, pile := range piles {
            eatHours += (pile + currentSpeed - 1) / currentSpeed
        }

        if eatHours <= h {
            possibleSpeed = min(possibleSpeed, currentSpeed)
            maxEatSpeed = currentSpeed - 1
        } else {
            minEatSpeed = currentSpeed + 1
        }
    }

    return possibleSpeed
}
