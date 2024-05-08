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
	googlepass := FindByName(passwords, "googlepass")

	if googlepass.Name != "googlepass" {
		t.Errorf("Should get googlepass")
	}

}
