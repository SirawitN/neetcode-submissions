type Tuple struct{
    Idx int
    Value int
}

func dailyTemperatures(temperatures []int) []int {
    stack := make([]Tuple, len(temperatures))
    ans := make([]int, len(temperatures))
    tos := -1

    for i, curT := range temperatures {
        for tos!=-1 && curT > stack[tos].Value {
            item := stack[tos]
            tos--

            ans[item.Idx] = i - item.Idx
        }

        tos++
        stack[tos] = Tuple{
            Idx: i,
            Value: curT,
        }
    }

    return ans
}
