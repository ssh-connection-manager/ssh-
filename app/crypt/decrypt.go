package crypt

import (
	"crypto/aes"
	"crypto/cipher"
)

func decrypt(key []byte, ciphertext []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return plaintext
}

func Decrypt(ciphertext string) string {
	decrypted := decrypt(GetKey(), []byte(ciphertext))

	return string(decrypted)
}
