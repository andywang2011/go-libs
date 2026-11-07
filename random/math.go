package random

import (
	"math/rand/v2"
)

// Shuffle permutes sli in place using the auto-seeded math/rand/v2 generator.
// It is NOT cryptographically secure; do not use it for secrets.
func Shuffle[T any](sli []T) {
	rand.Shuffle(len(sli), func(i, j int) {
		sli[i], sli[j] = sli[j], sli[i]
	})
}
