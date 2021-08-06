package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"log"
)

type Cryptor struct {
	gcm cipher.AEAD
}

func NewCryptor(key string) *Cryptor {
	blockCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		log.Fatal(err)
	}
	return &Cryptor{gcm: gcm}
}

func (c *Cryptor) Encrypt(data string) ([]byte, error) {
	nonce := make([]byte, c.gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := c.gcm.Seal(nonce, nonce, []byte(data), nil)
	return ciphertext, nil
}

func (c *Cryptor) Decrypt(data string) ([]byte, error) {
	nonce, ciphertext := []byte(data[:c.gcm.NonceSize()]), []byte(data[c.gcm.NonceSize():])
	plaintext, err := c.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
