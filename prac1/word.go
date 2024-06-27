package word

import "unicode"

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
