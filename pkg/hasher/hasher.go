package hasher

import (
	"crypto/sha1"
	"fmt"
)

type Hasher interface {
	Hash(s string) string
}

type SHA1Hasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{
		salt: salt,
	}
}

func (h *SHA1Hasher) Hash(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt)))
}
