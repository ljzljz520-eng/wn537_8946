package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func HashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:])
}
func CheckPassword(hash, password string) bool {
	return strings.EqualFold(hash, HashPassword(password))
}
func StrongPassword(password string) bool {
	return len(password) >= 8 && strings.IndexFunc(password, func(r rune) bool { return r >= '0' && r <= '9' }) >= 0
}
