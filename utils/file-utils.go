package utils

import (
	"io/fs"
	"os"
	"pass-manager/pass-manager/encrypt"
	"path/filepath"
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

// Get a root path as input
// Loop the children of the root
// Find all file matching the pattern (use shell based file pattern)
func FindFilesMatch(pattern string, root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func CreateNewFile(filename string) error {
	dir, _ := os.Getwd()
	separator := string(os.PathSeparator)
	_, err := os.Create(dir + separator + filename)

	return err
}

func CurrentDir() string {
	dir, _ := os.Getwd()

	return dir
}

func Separator() string {
	return string(os.PathSeparator)
}
