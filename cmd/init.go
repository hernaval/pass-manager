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
	fileConfigExt = "*.yaml"
	storageName   = "passdb.txt"
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
		var msg string
		if filename == "" || masterPassword == "" {
			fmt.Println("Filename and master password are required")
			cmd.Help()
			os.Exit(0)
		}
		// find config files
		rootDir := utils.CurrentDir()
		fileConfigs, err := utils.FindFilesMatch(fileConfigExt, rootDir)
		if err != nil {
			fmt.Printf("error while finding file configs %s %s", err, rootDir)
			os.Exit(1)
		}

		// load or create a config file with minimum template
		if len(fileConfigs) == 0 {
			err := utils.CreateNewFile(filename)
			if err != nil {
				fmt.Println("Error creating a config file")
				os.Exit(1)
			}
			// create password storage file
			err = utils.CreateNewFile(storageName)
			if err != nil {
				fmt.Println("Error creating a file based storage for password")
				os.Exit(1)
			}

			// init an empty storage
			passData := structs.PasswordData{}
			key := encrypt.EncKey([]byte(masterPassword))
			contents, _ := utils.ToJson(passData)
			err = utils.EncryptWrite(storageName, contents, key)
			if err != nil {
				fmt.Printf("Error init databse %s", err)
				os.Exit(1)
			}

			msg = `
				New storage is created.
				New configuration file is created.
				Password manager is ready to use. E N J O Y !!!
			`
		} else {
			// storageConf = viper.GetStringMap("storage")
			msg = `
			Password manager is already configured.
			Start using it.
			
			`
		}

		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	initCmd.Flags().StringVarP(&filename, "filename", "f", "pass-manager.yaml", "Name of the config file")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
