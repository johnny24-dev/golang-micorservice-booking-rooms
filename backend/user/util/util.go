package util

import (
	"math/rand"
)

func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}
