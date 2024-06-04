package utils

import (
	"testing"
)

func TestFindFilePattern(t *testing.T) {
	pattern := "*.go"
	dir := ".\\"
	files, err := FindFilesMatch(pattern, dir)
	if err != nil {
		t.Errorf("This function should not raise error")
	}
	if len(files) == 0 {
		t.Errorf("There is existing files match %s", pattern)
	}
}

func TestCreateFile(t *testing.T) {
	err := CreateNewFile("database.txt")

	if err != nil {
		t.Errorf("File should be created without error")
	}
}
