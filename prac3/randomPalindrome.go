package randomPalindrome

import (
	"math/rand"
	"unicode"
)

// this function creates random Palindrome of size 24 runes
func RandomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)

	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}
func IsPalindrome(s string) bool {
	runes := []rune{}
	for _, r := range s {
		if unicode.IsLetter(r) {
			runes = append(runes, r)
		}
	}

	for i := range runes {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}

	return true
}
