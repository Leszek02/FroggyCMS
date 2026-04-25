package auth

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

type PasswordHasher struct {
	sha256 hash.Hash
}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{
		sha256: sha256.New(),
	}
}

func (ph *PasswordHasher) HashPassword(password string) (string, error) {
	h := sha256.New()
	h.Write([]byte(password))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
