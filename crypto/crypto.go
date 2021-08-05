package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func Encrypt(key, data string) ([]byte, error) {
	blockCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	return ciphertext, nil
}

func Decrypt(key, data string) ([]byte, error) {
	blockCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, ciphertext := []byte(data[:gcm.NonceSize()]), []byte(data[gcm.NonceSize():])
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
