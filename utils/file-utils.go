package utils

import (
	"os"
	"pass-manager/pass-manager/encrypt"
)

// read the whole file contents
func Read(filepath string) ([]byte, error) {
	f, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// truncate file contents and write
func Write(filepath string, values []byte) error {
	err := os.WriteFile(filepath, values, 0666)
	if err != nil {
		return err
	}
	return nil
}

// Read File content
// Decrypt the content to get JSON FORMAT
func ReadDecrypt(filepath string, key []byte) ([]byte, error) {
	fileContents, err := Read(filepath)
	if err != nil {
		return nil, err
	}

	json, err := encrypt.Decrypt(fileContents, key)
	if err != nil {
		return nil, err
	}
	return json, nil
}

// Encrypt the values to save in file
// Write the content to file
func EncryptWrite(filepath string, values []byte, key []byte) error {
	encrypted, err := encrypt.Encrypt(key, values)
	if err != nil {
		return err
	}
	err = Write(filepath, encrypted)
	if err != nil {
		return err
	}
	return nil
}
