package securer

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashSHA256(message string) string {
	sum := sha256.Sum256([]byte(message))
	return base64.StdEncoding.EncodeToString(sum[:])
}
