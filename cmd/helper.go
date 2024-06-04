package cmd

import (
	"errors"
	"fmt"
	"os"
	"pass-manager/pass-manager/utils"
)

type Step int

const (
	fileConfigExt    = "*.yaml"
	databaseFilegExt = "*.psm"
)

// Initialization steps
const (
	DoConfigFile Step = iota
	DoDatabasee
)

func (s Step) String() string {
	names := [...]string{
		"DoConfigFile",
		"DoDatabse",
	}
	if s >= 0 && int(s) < len(names) {
		return names[s]
	}
	return ""
}

var (
	ErrLoadingFile = errors.New("error while finding file")

	// ErrConfigFileCorrupted = errors.New("Config file malformed")
)

func GetInitSteps() ([]Step, error) {
	rootDir := utils.CurrentDir()
	var results []Step
	fileConfigs, err := utils.FindFilesMatch(fileConfigExt, rootDir)
	if err != nil {
		return nil, ErrLoadingFile
	}
	if len(fileConfigs) == 0 {
		results = append(results, DoConfigFile)
	}
	// check file is well structured

	// check if database is present
	databaseFiles, err := utils.FindFilesMatch(databaseFilegExt, rootDir)
	if err != nil {
		return nil, ErrLoadingFile
	}
	if len(databaseFiles) == 0 {
		results = append(results, DoDatabasee)
	}
	// load or create a config file with minimum template
	return results, nil
}

func IsInitialized() bool {
	steps, err := GetInitSteps()

	return len(steps) == 0 && err == nil
}

func CheckInitialized() {
	if !IsInitialized() {
		fmt.Println("Password manager not initialized. Do it first!!!")
		initCmd.Usage()
		os.Exit(1)
	}
}
