package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// create hash password
func GenerateHash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

// check account password is match vs passwordhash or not
func CompareHash(password string, hash string) bool {
	newPasswordHash := GenerateHash(password)
	return newPasswordHash == hash
}
