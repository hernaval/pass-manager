package structs

type PasswordData struct {
	Id         int16
	Name       string
	Ciphertext string
}

type PasswordStorage struct {
	Data []PasswordData
}
