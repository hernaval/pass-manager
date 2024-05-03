package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// ensuring the encription key is 32-bit size supported by aes256
// convert it to byte slice
func EncKey(key []byte) []byte {
	sha := sha256.Sum256(key)
	return sha[:]
}

// Encrypts plaintext using AES and a given key
// create a cipher block using the key provided
// create a GCM instance (for encryption and authentication)
// generate the nonce (number used once) for the operation
// it prevents the risk of repeating the same cipher text
// use the GCM to encrypt the password (plaintext) along with nonce
// append the nonce to the return value to ensure same nonce will be used for decrypting
func Encrypt(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	cipherText := gcm.Seal(nil, nonce, plaintext, nil)

	return append(nonce, cipherText...), nil
}

// Hashes the master password
// It's a KDF (Key Derivation Function)
// convert the hashed value to a valid AES 32-bit key
func HashPassword(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hashed, nil
}

func VerifyPassword(hashed []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hashed, []byte(password))
}

// Decrypts the ciphtertext using key
// The ciphertext contains nonce and the ciphertext itself
// use GCM to decrypt the ciphertext using the extracted nonce
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("error after open")
		return nil, err
	}
	return plaintext, nil
}
