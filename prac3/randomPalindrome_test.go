package randomPalindrome

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := range 1000 {
		p := RandomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("%d, IsPalindrome(%q) = true", i, p)
		}
	}
}
