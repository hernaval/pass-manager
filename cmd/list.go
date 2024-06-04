/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"pass-manager/pass-manager/encrypt"
	"pass-manager/pass-manager/feature"
	"strings"

	"github.com/spf13/cobra"
)

// list of available flag of this command
var show bool

// listCmd represents the list command
// TODO flag for showing specified field
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all password",
	Long: `List all password and display all info about it, by default password text is hidden. 
	To show password use --show -s flag. 
	Example list --show
	`,
	PreRun: func(cmd *cobra.Command, args []string) {
		CheckInitialized()
	},
	Run: func(cmd *cobra.Command, args []string) {
		key := encrypt.EncKey([]byte(masterPassword))
		data, err := feature.List(key)
		if err != nil {
			switch err {
			case encrypt.ErrAuthentication:
				fmt.Println("Unauthorized: missing or wrong password provided.")

			default:
				fmt.Printf("error loading password %s", err)
			}
			os.Exit(1)
		}

		fmt.Println("ID		NAME	PASSWORD")
		for _, password := range data.Data {
			textPas := password.Ciphertext
			if !show {
				textPas = strings.Repeat("*", 7)
			}
			fmt.Printf("%d ----------> %s -----------> %s", password.Id, password.Name, textPas)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// flags
	listCmd.Flags().BoolVarP(&show, "show", "s", false, "Show password as plaintext")
}
