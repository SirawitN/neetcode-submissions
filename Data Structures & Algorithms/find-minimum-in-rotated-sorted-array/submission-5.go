func findMin(nums []int) int {
    left, right := 0, len(nums)-1
    possMin := nums[left]
    var mid int

    for left <= right {
        mid = (left+right)/2

        if  nums[left] < nums[mid] && nums[left] < nums[right] {
            possMin = min(possMin, nums[left])
            right = mid-1
        } else if nums[right] < nums[left] && nums[right] < nums[mid] {
            possMin = min(possMin, nums[right])
            left = mid+1
        } else {
            possMin = min(possMin, nums[mid])
            right = mid-1
        }
    }

    return possMin
}
