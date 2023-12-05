package securer

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"hash"
)

type RSA struct {
	bobPrivateKey       *rsa.PrivateKey
	bobPublicKey        *rsa.PublicKey
	RSAPublicKeyContent string
	PublicKeyContent    string
	hash                hash.Hash
	label               []byte
}

func NewRSA() *RSA {
	r := &RSA{}
	bits := 1024
	bobPrivateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	bobPublicKey := &bobPrivateKey.PublicKey
	rsaPublicKey := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(bobPublicKey))
	pubBytes, _ := x509.MarshalPKIXPublicKey(bobPublicKey)
	publicKey := base64.StdEncoding.EncodeToString(pubBytes)
	hash := sha256.New()
	label := []byte("")

	r.bobPrivateKey = bobPrivateKey
	r.bobPublicKey = bobPublicKey
	r.RSAPublicKeyContent = rsaPublicKey
	r.PublicKeyContent = publicKey
	r.hash = hash
	r.label = label
	return r
}

func (r *RSA) Encrypt(message []byte) (ciphertext []byte, err error) {
	ciphertext, err = rsa.EncryptOAEP(r.hash, rand.Reader, r.bobPublicKey, message, r.label)
	return
}

func (r *RSA) Decrypt(ciphertext []byte) (plaintext []byte, err error) {
	plaintext, err = rsa.DecryptOAEP(r.hash, rand.Reader, r.bobPrivateKey, ciphertext, r.label)
	return
}

func (r *RSA) EncryptString(message string) (encryptString string, err error) {
	var ciphertext []byte
	msgBytes := []byte(message)
	ciphertext, err = r.Encrypt(msgBytes)
	encryptString = string(ciphertext)
	return
}

func (r *RSA) DecryptString(message string) (decryptString string, err error) {
	var plaintext []byte
	msgBytes := []byte(message)
	plaintext, err = r.Decrypt(msgBytes)
	decryptString = string(plaintext)
	return
}

func (r *RSA) EncryptToBase64(message string) (encryptB64 string, err error) {
	var ciphertext []byte
	msgBytes := []byte(message)
	ciphertext, err = r.Encrypt(msgBytes)
	encryptB64 = base64.StdEncoding.EncodeToString(ciphertext)
	return
}

func (r *RSA) DecryptFromBase64(b64message string) (decryptString string, err error) {
	var msgBytes []byte
	var plaintext []byte
	msgBytes, err = base64.StdEncoding.DecodeString(b64message)

	if err == nil {
		plaintext, err = r.Decrypt(msgBytes)
		decryptString = string(plaintext)
	}
	return
}

func (r *RSA) EncodeToPemWithRSAPrivateKey() (privatekeyPem string) {
	privatekeyPem = string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(r.bobPrivateKey)}))
	return
}

func (r *RSA) EncodeToPemWithRSAPublicKey() (pubkeyPem string) {
	pubkeyPem = string(pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(r.bobPublicKey)}))
	return
}

func (r *RSA) EncodeToPemWithPrivateKey() (privatekeyPem string) {
	bytes, err := x509.MarshalPKCS8PrivateKey(r.bobPrivateKey)

	if err == nil {
		privatekeyPem = string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: bytes}))
	}

	return
}

func (r *RSA) EncodeToPemWithPublicKey() (pubkeyPem string) {
	bytes, err := x509.MarshalPKIXPublicKey(r.bobPublicKey)

	if err == nil {
		pubkeyPem = string(pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: bytes}))
	}

	return
}
