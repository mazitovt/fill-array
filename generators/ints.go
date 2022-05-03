package generators

import (
	"math/rand"
	"time"
)

func NewInt64Randomizer(min, max int64) (int64, func() int64) {
	left := max - min + 1
	rnd := *rand.New(rand.NewSource(time.Now().UnixNano()))
	return left, func(m1, m2 int64, r rand.Rand) func() int64 {
		return func() int64 {
			return r.Int63n(m1) + m2
		}
	}(left, min, rnd)
}
