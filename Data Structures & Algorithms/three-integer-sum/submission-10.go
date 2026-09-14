import (
    "slices"
)

func twoSum(nums []int, target int, start, stop int) [][2]int {
    var ans_pairs [][2]int
    i, j := start, stop

    for i<j {
        sum := nums[i] + nums[j]

        if sum==target {
            ans_pairs = append(ans_pairs, [2]int{nums[i], nums[j]})
            i += 1
            j -= 1
            
            for i<j && nums[i]==nums[i-1] {i++}
            for i<j && nums[j]==nums[j+1] {j--}
        } else if sum < target{
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

    result := make([][]int, 0, len(ans))
    for triplet := range ans {
        result = append(result, []int{triplet[0], triplet[1], triplet[2]})
    }

    return result
}
