func hashSetConv(str string) map[rune]int{
	out := make(map[rune]int, len(str))
	for _, char := range str {
		out[char] += 1
	}
	return out
}

func sortHashSet(hash map[rune]int) string{
	runes := make([]int, 0, len(hash))
	for char := range hash {
		runes = append(runes, int(char))
	}

	sort.Ints(runes)

	var out, str string
	for _, r := range runes {
		str = strconv.Itoa(r)
		out += fmt.Sprintf("%s:%v,", str, hash[rune(r)])
	}

	return out
}

func groupAnagrams(strs []string) [][]string {
	charGroups := make(map[string][]string, len(strs))
	for _, str := range strs{
		hashSet := hashSetConv(str)
		sortedHashSet := sortHashSet(hashSet)

		if words, ok := charGroups[sortedHashSet]; ok {
			charGroups[sortedHashSet] = append(words, str)
		} else {
			charGroups[sortedHashSet] = []string{str}
		}
	}

	results := make([][]string, 0, len(charGroups))
	for group := range charGroups {
		results = append(results, charGroups[group])
	}
	return results
}
