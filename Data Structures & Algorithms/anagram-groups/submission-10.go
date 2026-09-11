import (
    "maps"
    "slices"
)

func groupAnagrams(strs []string) [][]string {
    groupAnagram := make(map[[26]int][]string)

    for _, word := range strs {
        charFreq := [26]int{}
        for _, r := range word {
            charFreq[r - 'a'] += 1
        }
        groupAnagram[charFreq] = append(groupAnagram[charFreq], word)
    }

    return slices.Collect(maps.Values(groupAnagram))
}
