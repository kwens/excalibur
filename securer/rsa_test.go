package securer

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	randStr := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(randStr)
	fmt.Println(randStr[:8])
	fmt.Println(randStr[8:])
	fmt.Println(HashSHA256("hello adf24af world"))
	s1 := HashSHA256("hello adf24af world")
	s2 := HashSHA256("hello adf24af world")
	for i := 0; i < 1000; i++ {
		s2 = HashSHA256("hello adf24af world")
	}
	fmt.Println(s1 == s2)
	t.Fail()
}
