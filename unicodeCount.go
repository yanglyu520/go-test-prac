package unicodecount

import (
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

func CountUnicodeRune(s string) (map[rune]int, [utf8.UTFMax + 1]int, int) {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	reader := strings.NewReader(s)
	for {
		// ReadRune reads a single UTF-8 encoded Unicode character and returns the rune and its size in bytes.
		//If the encoded rune is invalid, it consumes one byte and returns unicode.ReplacementChar (U+FFFD) with a size of 1.
		r, n, err := reader.ReadRune()

		// case 1: break the loop if coming to the end of the input
		if err == io.EOF {
			break
		}

		// case 2: break if readrune coming to an unexpected error
		if err != nil {
			panic("err reading the rune")
		}

		// case 3: if a rune is invalid, then add to the invalid count
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++
	}

	return counts, utflen, invalid
}
