package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func Encrypt(key []byte, password []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("Cipher error : %s", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Printf("GCM error : %s", err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Printf("Nonce error : %s", err)
	}
	cipherText := gcm.Seal(nonce, nonce, []byte(password), nil)

	return cipherText
}
