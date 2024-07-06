package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func SimpleHash(str string) string {
	hash := sha256.Sum256([]byte(str))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func HmacHash(input string, key []byte) string {

	h := hmac.New(sha256.New, key)
	h.Write([]byte(input))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
