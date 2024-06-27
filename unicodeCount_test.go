package unicodecount

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

// mock input
var mockInput = `Hello, 世界! 你好，世界! こんにちは、世界`

func TestCountUnicodeRune(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantMap     map[rune]int
		wantUtfLen  [utf8.UTFMax + 1]int
		wantInvalid int
	}{
		{
			name:        "simple success test",
			input:       mockInput,
			wantMap:     map[int32]int{32: 3, 33: 2, 44: 1, 72: 1, 101: 1, 108: 2, 111: 1, 12289: 1, 12371: 1, 12385: 1, 12395: 1, 12399: 1, 12435: 1, 19990: 3, 20320: 1, 22909: 1, 30028: 3, 65292: 1},
			wantUtfLen:  [5]int{0, 11, 0, 15, 0},
			wantInvalid: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMap, gotUtfLen, gotInvalid := CountUnicodeRune(tt.input)
			assert.Equal(t, gotMap, tt.wantMap)
			assert.Equal(t, gotUtfLen, tt.wantUtfLen)
			assert.Equal(t, gotInvalid, tt.wantInvalid)
		})
	}
}
