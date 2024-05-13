/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"pass-manager/pass-manager/feature"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
// TODO flag for showing specified field
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all password",
	Long:  `List all password and display all info about it, including decrypted password and name`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := feature.List()
		if err != nil {
			fmt.Errorf("error loading password %s", err)
		}
		fmt.Println("ID		NAME	PASSWORD")
		for _, password := range data.Data {
			fmt.Printf("%d ----------> %s -----------> %s", password.Id, password.Name, password.Ciphertext)
			fmt.Println()
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
