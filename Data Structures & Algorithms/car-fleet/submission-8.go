import (
	"slices"
	"cmp"
)

type Car struct {
	Position int
	Speed int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	for i:=0; i<len(position); i++ {
		cars[i] = Car{
			Position: position[i],
			Speed: speed[i],
		}
	}

	slices.SortFunc(cars, func(a, b Car) int {
		return cmp.Compare(a.Position, b.Position)
	})

	remainTravelTime := make([]float64, len(cars))
	for i, c := range cars {
		remainTravelTime[i] = float64(target-c.Position) / float64(c.Speed)
	}
	// fmt.Println(remainTravelTime)

	carFleets := 1
	maxTravelTime := remainTravelTime[len(remainTravelTime)-1]
	for i:=len(remainTravelTime)-2; i>=0; i-- {
		if remainTravelTime[i] > maxTravelTime {
			carFleets++
			maxTravelTime = remainTravelTime[i]
		}
	}

	return carFleets
}
