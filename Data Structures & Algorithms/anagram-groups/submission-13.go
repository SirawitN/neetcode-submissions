import (
	"slices"
	"maps"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, word := range strs {
		key := [26]int{}

		for _, ch := range word {
			key[unicode.ToLower(ch) - 'a'] += 1
		}
		groups[key] = append(groups[key], word)
	}

	return slices.Collect(maps.Values(groups))
}
