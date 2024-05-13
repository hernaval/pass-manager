package feature

import (
	"pass-manager/pass-manager/encrypt"
	"pass-manager/pass-manager/storage"
	"pass-manager/pass-manager/structs"
)

func AddPassword(passData structs.PasswordData) (structs.PasswordData, error) {
	//create a master password
	masterPassword := "adminadmin"

	//create KDF key for the master password
	key := encrypt.EncKey([]byte(masterPassword))

	// save
	err := storage.StorePassword(passData, key)
	if err != nil {
		return structs.PasswordData{}, err
	}

	return passData, err
}

func GetByName(name string) (structs.PasswordData, error) {
	//create a master password
	masterPassword := "adminadmin"

	//create KDF key for the master password
	key := encrypt.EncKey([]byte(masterPassword))

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

func List() (structs.PasswordStorage, error) {
	//create a master password
	masterPassword := "adminadmin"

	//create KDF key for the master password
	key := encrypt.EncKey([]byte(masterPassword))

	// retrieve password list (datasource)
	passwords, err := storage.LoadPassword(key)
	if err != nil {
		return structs.PasswordStorage{}, err
	}

	return passwords, nil
}
