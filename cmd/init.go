/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"pass-manager/pass-manager/encrypt"
	"pass-manager/pass-manager/structs"
	"pass-manager/pass-manager/utils"

	"github.com/spf13/cobra"
)

const (
	// fileConfigExt = "*.yaml"
	storageName = "passdb.psm"
)

// flags variable
var filename string

// initCmd represents the init command

// TODO
// verify existing config | storage
// ask for override | backup for storage
// other command should not be accessible if not initialized
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a storage to manage password",
	Long: `This command will setup a storage for the password manager. 
	It will first create a file based storage to save password and then
	create configuration file. You can choose the path for each file.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// var storageConf map[string]any
		fmt.Println(
			`
|  _ \  / ___|   |  \/  |
| |_) | \___ \   | |\/| |
|  __/   ___) |  | |  | |
|_|     |____/   |_|  |_|
v2024.0.0.1-alpha
		`)
		var msg string
		if filename == "" || masterPassword == "" {
			fmt.Println("Filename and master password are required")
			cmd.Help()
			os.Exit(0)
		}
		// find config files
		steps, err := GetInitSteps()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if len(steps) > 0 {

			for _, step := range steps {
				switch step {
				case DoConfigFile:
					createNewConfigFile()

				case DoDatabasee:
					createEmptyDatabse()
				}
			}

			msg = `
				New storage is created.
				New configuration file is created.
				Password manager is ready to use. E N J O Y !!!
			`
		} else {
			msg = `
				Password manager is already configured.
				Start using it.
			`
		}
		// storageConf = viper.GetStringMap("storage")

		fmt.Println(msg)
	},
}

func createNewConfigFile() {
	err := utils.CreateNewFile(filename)
	if err != nil {
		fmt.Println("Error creating a file based storage for password")
		os.Exit(1)
	}
}

func createEmptyDatabse() {
	passData := structs.PasswordData{}
	key := encrypt.EncKey([]byte(masterPassword))
	contents, _ := utils.ToJson(passData)
	err := utils.EncryptWrite(storageName, contents, key)
	if err != nil {
		fmt.Printf("Error init databse %s", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)

	// flags
	initCmd.Flags().StringVarP(&filename, "filename", "f", "pass-manager.yaml", "Name of the config file")
}
