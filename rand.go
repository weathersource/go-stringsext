package stringsext

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Rand creates a random string of a given length. If validChrs is not provided
// the set of [a-zA-Z0-9] will be used.
func Rand(length int, validChrs ...string) string {

	valid := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if len(validChrs) > 0 {
		valid = validChrs[0]
	}
	letterRunes := []rune(valid)

	rs := make([]rune, length)

	// populate each rune with a random letter
	for i := range rs {
		rs[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	// return string conversion
	return string(rs)
}
