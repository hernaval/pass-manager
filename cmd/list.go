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
	Run: func(cmd *cobra.Command, args []string) {
		key := encrypt.EncKey([]byte(masterPassword))
		data, err := feature.List(key)
		if err != nil {
			fmt.Printf("error loading password %s", err)
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

	// Here you will define your flags and configuration settings.
	listCmd.Flags().BoolVarP(&show, "show", "s", false, "Show password as plaintext")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
