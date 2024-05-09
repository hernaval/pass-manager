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
