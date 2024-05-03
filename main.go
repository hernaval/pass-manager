package main

import (
	"fmt"
	"pass-manager/pass-manager/encrypt"
	"pass-manager/pass-manager/utils"
)

func main() {
	//create a master password
	masterPassword := "adminadmin"

	//create hashed key for the master password
	key, err := encrypt.HashPassword(masterPassword)
	key = encrypt.EncKey(key)
	if err != nil {
		fmt.Printf("Error hashing password : %s", err)
	}

	//encrypt the plaintext password along with the key
	password := []byte("password123we")
	cipherText, err := encrypt.Encrypt(key, password)
	if err != nil {
		fmt.Printf("Error encrypting plaintext : %s", err)
	}
	fmt.Printf("Encrypted password : %s\n", string(cipherText))

	// write to file as database
	utils.Write("pass.txt", cipherText)

	//decrypt the ciphertext along with the key
	cf, err := utils.Read("pass.txt")
	if err != nil {
		fmt.Printf("Error reading file")
	}
	plaintext, err := encrypt.Decrypt(cf, key)
	if err != nil {
		fmt.Printf("Error decrypting  : %s\n", err)
	}
	fmt.Printf("Decrypted password : %s", plaintext)

	//authentication
	auth := encrypt.VerifyPassword(key, masterPassword)
	if auth == nil {
		fmt.Println("You are successfuly authenticated")
	} else {
		fmt.Println("Wrong credentials")
	}

}
