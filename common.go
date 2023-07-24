package wrench

import (
	"math/rand"
	"time"
)

func SliceContain[T comparable](s T, in []T) bool {
	for _, e := range in {
		if e == s {
			return true
		}
	}
	return false
}

func Max[T Number](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Number](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func Shuffle[T any](in []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(in), func(i, j int) {
		in[i], in[j] = in[j], in[i]
	})
}
