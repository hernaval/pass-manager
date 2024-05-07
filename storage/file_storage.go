package storage

// This is a file storage based
// The file content is encrypted
// JSON format is prefered for further use of Querying
import (
	"fmt"
	"pass-manager/pass-manager/encrypt"
	"pass-manager/pass-manager/structs"
	"pass-manager/pass-manager/utils"
)

// Load all passwords
// Add the the new password
// Convert it to JSON format
// Save the encrypted json to file
func StorePassword(data structs.PasswordData, key []byte) error {
	passwords, err := LoadPassword(key)
	if err != nil {
		fmt.Println("misy olana", err)
		return err
	}
	passwords.Data = append(passwords.Data, data)
	storageValue, err := utils.ToJson(passwords)
	if err != nil {
		fmt.Println("misy olana", err)
		return err
	}
	encrypted, err := encrypt.Encrypt(key, storageValue)
	if err != nil {
		return err
	}
	err = utils.Write("pass.txt", encrypted)
	if err != nil {
		return err
	}

	return nil
}

// Read File content
// Decrypt the content to get JSON FORMAT
// Convert JSON to PasswordStorage
// return
func LoadPassword(key []byte) (structs.PasswordStorage, error) {
	var passwords structs.PasswordStorage

	fileContents, err := utils.Read("pass.txt")
	if err != nil {
		return structs.PasswordStorage{}, err
	}

	json, err := encrypt.Decrypt(fileContents, key)
	if err != nil {
		fmt.Println("misy olana", err)
		return structs.PasswordStorage{}, err
	}

	err = utils.FromJson(json, &passwords)

	if err != nil {
		fmt.Println("Error loading", err)
	}

	return passwords, nil
}
