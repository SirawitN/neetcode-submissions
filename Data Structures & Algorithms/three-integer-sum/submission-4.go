import "slices"

func twoSum(nums []int, target int) *[][2]int {
    a, b := 0, len(nums)-1
    ans := [][2]int{}

    for a<b {
        if nums[a]+nums[b]==target {
            ans = append(ans, [2]int{nums[a], nums[b]})
            a, b = a+1, b-1
            continue
        }

        if nums[a]+nums[b]<target {
            a += 1
        } else {
            b -= 1
        }
    }

    return &ans
}

func threeSum(nums []int) [][]int {
    distinctTriplets := make(map[[3]int]struct{})
    
    slices.Sort(nums)
    for i, n := range nums[:len(nums)-2] {
        twoElements := twoSum(nums[i+1:], -n)
        if twoElements != nil {
            for _, t := range *twoElements{
                distinctTriplets[[3]int{n, t[0], t[1]}] = struct{}{}
            }
        }
    }
    ans := make([][]int, len(distinctTriplets))
    ansIdx := 0
    for triplet := range distinctTriplets {
        ans[ansIdx] = triplet[:]
        ansIdx += 1
    }

    return ans
}
