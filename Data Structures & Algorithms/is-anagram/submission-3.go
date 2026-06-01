func isAnagram(s string, t string) bool {
    if len(s)!=len(t) {
        return false
    }

    seenCharacters := make(map[rune]int, len(s))
    for _, char := range s {
        seenCharacters[char] += 1
    }

    for _, char := range t {
        count, exist := seenCharacters[char]

        if !exist || count <= 0 {
            return false
        } else {
            seenCharacters[char] -= 1
        }
    }
    return true
}
