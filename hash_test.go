package stringsext

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {

	hash := Hash("foo", "bar")
	assert.Equal(t, "c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2", hash)
}
