func searchDeflection(nums []int) int {
    left, right := 0, len(nums)-1
    var mid int
    possDeflectPt := left
    possDeflectVal := nums[left]

    for left<=right {
        mid = (left+right) / 2

        if nums[left]<nums[mid] && nums[left]<nums[right] {
            right = mid-1
            if nums[left] < possDeflectVal {
                possDeflectVal = nums[left]
                possDeflectPt = left
            } 
        } else if nums[right] < nums[left] && nums[right] < nums[mid] {
            left = mid+1
            if nums[right] < possDeflectVal {
                possDeflectVal = nums[right]
                possDeflectPt = right
            }
        } else {
            right = mid-1
            if nums[mid] < possDeflectVal {
                possDeflectVal = nums[mid]
                possDeflectPt = mid
            }
        }
    }

    return possDeflectPt
}

func binSearch(nums []int, left, right, target int) int {
    var mid int
    
    for left <= right {
        mid = (left+right) / 2

        if nums[mid]==target {
            return mid
        }

        if nums[mid]<target {
            left = mid+1
        } else {
            right = mid-1
        }
    }

    return -1
}

func search(nums []int, target int) int {
    deflectionPt := searchDeflection(nums)

    if nums[deflectionPt]==target {
        return deflectionPt
    }

    left, right := 0, len(nums)-1
    firstHalf := binSearch(nums, left, deflectionPt-1, target)
    if firstHalf != -1 {
        return firstHalf
    }
    return binSearch(nums, deflectionPt+1, right, target)
}
