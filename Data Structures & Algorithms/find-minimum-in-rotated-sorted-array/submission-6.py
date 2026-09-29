class Solution:
    def findMin(self, nums: List[int]) -> int:
        left, right = 0, len(nums)-1
        possMin = nums[left]

        while left<=right:
            mid = (left+right) // 2

            if (nums[left]<nums[mid]) and (nums[left]<nums[right]):
                possMin = min(possMin, nums[left])
                right = mid-1
            elif (nums[right]<nums[left]) and (nums[right]<nums[mid]):
                possMin = min(possMin, nums[right])
                left = mid+1
            else:
                possMin = min(possMin, nums[mid])
                right = mid-1
        
        return possMin
        