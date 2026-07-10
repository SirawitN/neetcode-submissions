func checkIdenticalMaps(m1, m2 []int) bool {
	for i := range 26 {
		if m1[i]!=m2[i] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	len1, len2 := len(s1), len(s2)
	if len1 > len2 {
		return false
	}

	map1 := make([]int, 26)
	map2 := make([]int, 26)

	for _, r := range s1 {
		map1[r - 'a'] += 1
	}

	var left int
	for i, r := range s2 {
		map2[r-'a']+=1

		if i-left+1 > len1 {
			map2[s2[left]-'a'] -= 1
			left += 1
		}
		if checkIdenticalMaps(map1, map2) {
			return true
		}
	}
	
	return false
}
