package storage

// This is a file storage based
// The file content is encrypted
// JSON format is prefered for further use of Querying
import (
	"fmt"
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
		fmt.Println("storage error ", err)
		return err
	}
	passwords.Data = append(passwords.Data, data)
	storageValue, err := utils.ToJson(passwords)
	if err != nil {
		fmt.Println("storage error ", err)
		return err
	}

	err = utils.EncryptWrite("passdb.txt", storageValue, key)
	if err != nil {
		return err
	}

	return nil
}

// Convert the fileContents in JSON format to PasswordStorage
func LoadPassword(key []byte) (structs.PasswordStorage, error) {
	var passwords structs.PasswordStorage

	json, err := utils.ReadDecrypt("passdb.txt", key)
	if err != nil {
		fmt.Println("storage error ", err)
		return structs.PasswordStorage{}, err
	}

	err = utils.FromJson(json, &passwords)
	if err != nil {
		fmt.Println("Error loading", err)
	}

	return passwords, nil
}

// get the password by name in datasource
// datasource should be retrieved beforehand
func FindByName(datasource structs.PasswordStorage, name string) (structs.PasswordData, error) {
	filterName := func(pass structs.PasswordData) bool { return pass.Name == name }
	filtered := utils.Filter(datasource.Data, filterName)
	if len(filtered) == 0 {
		return structs.PasswordData{}, fmt.Errorf("no password found with name %s", name)
	}

	return filtered[0], nil //TODO shoud be refactored || name is unique
}
