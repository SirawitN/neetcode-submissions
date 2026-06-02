const ALPHABETS = 26

func hashSetConv(str string) [ALPHABETS]int{
	out := [ALPHABETS]int{}
	for _, char := range str {
		out[char - 'a'] += 1
	}
	return out
}


func groupAnagrams(strs []string) [][]string {
	charGroups := make(map[[ALPHABETS]int][]string, len(strs))
	for _, str := range strs{
		hashSet := hashSetConv(str)

		if words, ok := charGroups[hashSet]; ok {
			charGroups[hashSet] = append(words, str)
		} else {
			charGroups[hashSet] = []string{str}
		}
	}

	results := make([][]string, 0, len(charGroups))
	for group := range charGroups {
		results = append(results, charGroups[group])
	}
	return results
}
