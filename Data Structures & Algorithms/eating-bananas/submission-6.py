import math

class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        minEatSpeed = 1
        maxEatSpeed = max(piles)

        if len(piles)==h :
            return maxEatSpeed

        possibleSpeed = maxEatSpeed
        while minEatSpeed<=maxEatSpeed:
            currentSpeed = math.floor((minEatSpeed + maxEatSpeed) / 2)
            # print(minEatSpeed, maxEatSpeed, currentSpeed)

            eatHours = sum([math.ceil(pile/currentSpeed) for pile in piles])
            # print(eatHours, h)

            if eatHours <= h :
                possibleSpeed = min(possibleSpeed, currentSpeed)
                maxEatSpeed = currentSpeed - 1
            else:
                minEatSpeed = currentSpeed + 1
        
        return possibleSpeed