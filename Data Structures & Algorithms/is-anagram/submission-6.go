import (
	"maps"
)

func isAnagram(s string, t string) bool {
	if len(s)!=len(t) {
		return false
	}

	charSet1 := make(map[rune]int)
	charSet2 := make(map[rune]int)

	for _, r := range s {
		charSet1[r] += 1
	}

	for _, r := range t {
		charSet2[r] += 1
	}

	return maps.Equal(charSet1, charSet2)
}
