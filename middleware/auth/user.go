package auth

import (
	"crypto/sha1"
	"encoding/hex"
)

type User struct {
	Email string `json:"email"`
}

type UserResponse struct {
	AccessKey string `json:"accesskey"`
}

func GenerateUserAccessKey(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
