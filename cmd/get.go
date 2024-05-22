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

// getCmd represents the get command
// TODO flag to display the encrypted pass
var getCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Get password by name",
	Long: `Retrieve a stored password by its name will return the decrypted password if exist, 
	otherwise raise an error. 
	Example usage : get googlepass
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := encrypt.EncKey([]byte(masterPassword))
		name := args[0]
		password, err := feature.GetByName(name, key)
		if err != nil {
			fmt.Printf("error getting password: %s", err)
			os.Exit(1)
		}
		fmt.Printf(password.Ciphertext)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
