import math

class Solution:
    def findDeflection(self, nums: List[int]) -> int:
        left, right = 0, len(nums)-1
        possMin, possDft = nums[left], left

        while left<=right:
            mid = (left+right)//2
            # print(mid)

            if (nums[left]<nums[mid]) and (nums[left]<nums[right]):
                right = mid-1
                if nums[left] < possMin:
                    possMin = nums[left]
                    possDft = left
            elif (nums[right]<nums[left]) and (nums[right]<nums[mid]):
                left = mid+1
                if nums[right] < possMin:
                    possMin = nums[right]
                    possDft = right
            else:
                right = mid-1
                if nums[mid] < possMin:
                    possMin = nums[mid]
                    possDft = mid
        
        return possDft

    def binSearch(self, nums: List[int], left: int, right: int, target: int) -> int:
        while left <= right:
            mid = (left+right)//2

            if nums[mid]==target:
                return mid

            if nums[mid]>target:
                right = mid-1
            else:
                left = mid+1

        return -1

    def search(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums)-1
        deflectionPt = self.findDeflection(nums)

        # print("Test", deflectionPt)
        
        if nums[deflectionPt]==target:
            return deflectionPt

        firstHalf = self.binSearch(nums, left, deflectionPt-1, target)
        if firstHalf!=-1:
            return firstHalf
        return self.binSearch(nums, deflectionPt+1, right, target)
            