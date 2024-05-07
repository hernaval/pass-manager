package main

import (
	"fmt"
	"pass-manager/pass-manager/encrypt"
	"pass-manager/pass-manager/storage"
	"pass-manager/pass-manager/structs"
	"testing"
)

func TestEncryptionDecryption(t *testing.T) {
	var passwords structs.PasswordStorage
	//create a master password
	masterPassword := "adminadmin"

	//create KDF key for the master password
	key := encrypt.EncKey([]byte(masterPassword))

	//create data
	passData := structs.PasswordData{Id: 1, Name: "cbs", Ciphertext: "password123"}

	storage.StorePassword(passData, key)

	passwords, _ = storage.LoadPassword(key)

	fmt.Println(passwords)

	// if err != nil {
	// 	t.Errorf("Failing code")
	// }
}
