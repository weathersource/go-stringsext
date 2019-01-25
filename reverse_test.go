package stringsext

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {

	str := "abcdefg"
	assert.Equal(t, "gfedcba", Reverse(str))
}
