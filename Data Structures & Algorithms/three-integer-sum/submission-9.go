import (
    "slices"
    "maps"
)

func twoSum(nums []int, target int, start, stop int) [][2]int {
    ans_pairs := make([][2]int, 0, len(nums))
    i, j := start, stop

    for i<j {
        sum := nums[i] + nums[j]

        if sum==target {
            ans_pairs = append(ans_pairs, [2]int{nums[i], nums[j]})
            i += 1
            j -= 1
            continue
        }

        if sum < target{
            i += 1
        } else {
            j -= 1
        }
    }

    return ans_pairs
}

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    n := len(nums)

    ans := make(map[[3]int]struct{}, n)
    for i:=0; i<n-2; i++ {
        target := -1 * nums[i]
        pairs := twoSum(nums, target, i+1, n-1)

        if len(pairs)>0 {
            for _, p := range pairs{
                ans[[3]int{nums[i], p[0], p[1]}] = struct{}{}
            }
        }
    }

    // 1. Collect the [3]int keys from your map
    keys := slices.Collect(maps.Keys(ans)) // type is [][3]int

    // 2. Convert [][3]int to [][]int
    result := make([][]int, len(keys))
    for i, triplet := range keys {
        result[i] = triplet[:] // slice expression turns [3]int into []int
    }

    return result
}
