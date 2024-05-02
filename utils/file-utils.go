package utils

import (
	"os"
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
