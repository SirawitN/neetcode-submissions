func isPalindrome(s string) bool {
    i, j := 0, len(s)-1

    for i<j {
        i_rune, j_rune := rune(s[i]), rune(s[j])
        if !unicode.IsLetter(i_rune) && !unicode.IsDigit(i_rune) {
            i += 1
            continue
        }
        if !unicode.IsLetter(j_rune) && !unicode.IsDigit(j_rune) {
            j -= 1
            continue
        }

        if unicode.ToLower(i_rune) != unicode.ToLower(j_rune) {
            return false
        }

        i, j = i+1, j-1
    }

    return true
}
