
func isIdentical(listA, listB [26]int) bool {
	for i:=0; i<26; i++ {
		if listA[i]!=listB[i] {
			return false
		}
	}

	return true
}

func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	chFreq1 := [26]int{}
	for _, ch := range s1 {
		chFreq1[ch-'a'] += 1
	}

	left, right := 0, n1-1
	s2_chars := []rune(s2)
	chFreq2 := [26]int{}
	
	for i:=left; i<=right; i++ {
		chFreq2[s2_chars[i]-'a'] += 1
	}
	
	for right<n2 {
		if isIdentical(chFreq1, chFreq2) {
			return true
		}
		// fmt.Println(chFreq2, s2[left], s2[right])
		chFreq2[s2_chars[left]-'a'] -= 1
		left += 1
		right += 1
		if right < n2 {
			chFreq2[s2_chars[right]-'a'] += 1
		}
	}

	return false
}
