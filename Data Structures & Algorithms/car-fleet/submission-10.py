class Solution:
	def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
		cars = [(position[i], speed[i]) for i in range(len(position))]
		sortedByPosition = sorted(cars, key=lambda c: c[0])
		remainTravelTime = [(target-c[0])/c[1] for c in sortedByPosition]

		carFleet = 1
		maxTravelTimeSoFar = remainTravelTime[-1]
		for i in range(len(remainTravelTime)-2, -1, -1):
			if remainTravelTime[i]>maxTravelTimeSoFar:
				carFleet += 1
				maxTravelTimeSoFar = remainTravelTime[i]

		return carFleet
        