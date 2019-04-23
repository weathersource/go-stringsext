package stringsext

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestInSlice(t *testing.T) {
	s := []string{"test", "test2","test3"}

	assert.Equal(t, InSlice("test",s), true)

	assert.Equal(t, InSlice("aaaaaaaaaa",s), false)
}
