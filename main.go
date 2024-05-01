package main

import (
	"fmt"
	"pass-manager/pass-manager/encrypt"
)

func main() {
	key := []byte("mysecuredkeymysecuredkeymysecure")
	masterPassword := []byte("mymasterpass")
	cipher := encrypt.Encrypt(key, masterPassword)
	fmt.Println("encrypted ", string(cipher[:]))

	//create a master password

	//create devired key for the password

	//give the password to encrypt

}
