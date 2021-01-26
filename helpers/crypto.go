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

// NewSHA256FromString helper
func NewSHA256FromString(data string) []byte {
	return NewSHA256([]byte(data))
}

// NewSHA256String helper
func NewSHA256String(data string) string {
	hash := NewSHA256FromString(data)
	return hex.EncodeToString(hash)
}
