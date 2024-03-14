package util

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func Endcode(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
