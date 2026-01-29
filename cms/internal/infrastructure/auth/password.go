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
	ph.sha256.Write([]byte(password))
	return fmt.Sprintf("%x", ph.sha256.Sum(nil)), nil
}
