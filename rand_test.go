package stringsext

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {

	str := Rand(10)
	assert.Equal(t, 10, len(str))

	str2 := Rand(10, "a")
	assert.Equal(t, "aaaaaaaaaa", str2)
}
