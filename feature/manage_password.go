package feature

import (
	"pass-manager/pass-manager/storage"
	"pass-manager/pass-manager/structs"
)

func AddPassword(passData structs.PasswordData, key []byte) (structs.PasswordData, error) {
	// save
	err := storage.StorePassword(passData, key)
	if err != nil {
		return structs.PasswordData{}, err
	}

	return passData, err
}

func GetByName(name string, key []byte) (structs.PasswordData, error) {
	// retrieve password list (datasource)
	datasource, err := storage.LoadPassword(key)
	if err != nil {
		return structs.PasswordData{}, err
	}

	password, err := storage.FindByName(datasource, name)
	if err != nil {
		return structs.PasswordData{}, err
	}

	return password, nil
}

func List(key []byte) (structs.PasswordStorage, error) {
	// retrieve password list (datasource)
	passwords, err := storage.LoadPassword(key)
	if err != nil {
		return structs.PasswordStorage{}, err
	}

	return passwords, nil
}
