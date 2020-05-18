package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

// NewSHA256 ...
func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func NewSHA256FromString(data string) []byte {
	return NewSHA256([]byte(data))
}

func NewSHA256String(data string) string {
	hash := NewSHA256FromString(data)
	return hex.EncodeToString(hash)
}
