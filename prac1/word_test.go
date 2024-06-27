package word

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// use a test table is common in golang test
	tests := []struct {
		testName string
		input    string
		want     bool
	}{
		{
			testName: "successful test with ascii letters",
			input:    "xxx",
			want:     true,
		},
		{
			testName: "successful test with ascii letters and numbers",
			input:    "1xxx1",
			want:     true,
		},
		{
			testName: "nonsuccessful test with non-ascii letters",
			input:    "我我",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := IsPalindrome(tt.input); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v", tt.input, tt.want)
			}
		})
	}
}

func TestNonPalindrome(t *testing.T) {
	got := IsPalindrome("xxy")
	want := false
	if got != want {
		t.Errorf("want %v, but got %v", want, got)
	}
}
