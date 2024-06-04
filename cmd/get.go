/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"pass-manager/pass-manager/encrypt"
	"pass-manager/pass-manager/feature"

	"github.com/spf13/cobra"
)

var copy bool

var getCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Get password by name",
	Long: `Retrieve a stored password by its name will return the decrypted password if exist, 
	otherwise raise an error. 
	Example usage : get googlepass
	`,
	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		CheckInitialized()
	},
	Run: func(cmd *cobra.Command, args []string) {
		key := encrypt.EncKey([]byte(masterPassword))
		name := args[0]
		password, err := feature.GetByName(name, key)
		if err != nil {
			switch err {
			case encrypt.ErrAuthentication:
				fmt.Println("Unauthorized: missing or wrong password provided.")

			default:
				fmt.Printf("Error get password")
			}
			os.Exit(1)
		}

		fmt.Printf(password.Ciphertext)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().BoolVarP(&copy, "copy", "c", false, "Copy password to clipboard")
}
