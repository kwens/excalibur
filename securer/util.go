package securer

import (
	"crypto/rand"
	"encoding/base64"
)

func NewRandString(len uint) string {
	bytes := make([]byte, len)
	rand.Read(bytes)
	return base64.StdEncoding.EncodeToString(bytes)[:len]
}
