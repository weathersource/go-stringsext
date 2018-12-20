package stringsext

import (
	"crypto/sha256"
	"encoding/hex"
)

// hash returns a salted sha256 hash of the str parameter
func Hash(str string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(str + salt))
	return hex.EncodeToString(hash.Sum(nil)[:])
}
