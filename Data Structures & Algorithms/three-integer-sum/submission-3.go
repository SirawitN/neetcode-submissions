func twoSum(nums []int, target int) [][]int{
    i, j := 0, len(nums)-1
    var sum int
    var out [][]int

    for {
        if i >= j {
            break
        }

        sum = nums[i] + nums[j]
        if sum==target {
            out = append(out, []int{nums[i], nums[j]})
            i, j = i+1, j-1
            continue
        }

        if sum<target {
            i += 1
        } else {
            j -= 1
        }
    }
    return out
}

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    
    distinctTriplet := make(map[[3]int]struct{})
    output := [][]int{}

    var target int
    var tmp [3]int
    var triplets [][]int
    for i, num := range nums[:len(nums)-2] {
        target = -1*num
        triplets = twoSum(nums[i+1:], target)
        for _, triplet := range triplets {
            tmp[0] = num
            tmp[1] = triplet[0]
            tmp[2] = triplet[1]
            distinctTriplet[tmp] = struct{}{}
        }
    }

    for triplet := range distinctTriplet {
        output = append(output, triplet[:])
    }
    return output
}
