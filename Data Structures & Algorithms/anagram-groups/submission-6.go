import (
    "maps"
    "slices"
)

func groupAnagrams(strs []string) [][]string {
    results := make(map[[26]int][]string)

    for _, str := range strs {
        charFreq := [26]int{}
        for _, r := range str {
            charFreq[r - 'a'] += 1
        }

        results[charFreq] = append(results[charFreq], str)
    }

    return slices.Collect(maps.Values(results))
}
