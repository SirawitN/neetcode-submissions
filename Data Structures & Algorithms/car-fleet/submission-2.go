import (
	"slices"
	"cmp"
)

type Car struct{
	position int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	N := len(position)
	cars := make([]Car, N)
	for i := range N {
		cars[i] = Car{
			position: position[i],
			speed: speed[i],
		}
	}


	slices.SortFunc(cars, func(a, b Car) int {
		return cmp.Compare(a.position, b.position)
	})

	stack := make([]float64, N)
	for i, car := range cars {
		timeToTarget := float64((target - car.position))/float64(car.speed)
		stack[i] = timeToTarget
	}

	var carFleet int
	var lastArriveTime float64
	for i:=N-1; i>=0; i-=1{
		if lastArriveTime==0.0 || stack[i] > lastArriveTime{
			lastArriveTime = stack[i]
			carFleet += 1
		} else {
			continue
		}
	}

	return carFleet
}
