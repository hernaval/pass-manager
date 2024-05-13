package storage

import (
	"pass-manager/pass-manager/structs"
	"testing"
)

func TestFindByName(t *testing.T) {
	// given a password storage
	passwords := structs.PasswordStorage{
		Data: []structs.PasswordData{
			{Id: 1, Name: "googlepass", Ciphertext: "googlepassword"},
			{Id: 2, Name: "facepass", Ciphertext: "facepassword"},
		},
	}

	// when retrieving googlepass
	googlepass, _ := FindByName(passwords, "googlepass")

	if googlepass.Name != "googlepass" {
		t.Errorf("Should get googlepass")
	}

	// when retrieving non existing password
	_, err := FindByName(passwords, "nopass")
	if err == nil {
		t.Errorf("Shoud get error loading non existing password")
	}

}
