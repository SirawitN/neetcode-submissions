func isPermutation(charFreq1, charFreq2 [26]int) bool {
    for i := range 26 {
        if charFreq1[i] != charFreq2[i] {
            return false
        }
    }

    return true
}

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    charFreqOfS1 := [26]int{}
    charFreqOfS2 := [26]int{}

    for _, r := range s1 {
        charFreqOfS1[r - 'a'] += 1
    }

    startWind := 0
    stopWind := min(startWind+len(s1)-1, len(s1)-1)
    for i:=startWind; i<stopWind; i++{
        charFreqOfS2[rune(s2[i]) - 'a'] += 1
    }
    
    for stopWind < len(s2) {
        charFreqOfS2[rune(s2[stopWind]) - 'a'] += 1
        if isPermutation(charFreqOfS1, charFreqOfS2) {
            return true
        }

        charFreqOfS2[rune(s2[startWind]) - 'a'] -= 1
        startWind++
        stopWind++
    }

    return false
}
