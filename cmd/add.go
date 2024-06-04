/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"pass-manager/pass-manager/encrypt"
	"pass-manager/pass-manager/feature"
	"pass-manager/pass-manager/structs"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [name] [password]",
	Short: "Add a new password to encrypt",
	Long: `A plaintext password can be saved to a secure and encrypted database 
	using this command. Give a significant name to get it seamlessly next time you
	want to use it.
	`,
	Args: cobra.RangeArgs(0, 3),

	PreRun: func(cmd *cobra.Command, args []string) {
		CheckInitialized()
	},
	Run: func(cmd *cobra.Command, args []string) {
		key := encrypt.EncKey([]byte(masterPassword))
		// q := []tui.Question{tui.NewQuestion("Enter a name", ""), tui.NewQuestion("Enter a password", "test")}
		// container := tui.New(q)

		// p := tea.NewProgram(*container)
		// if _, err := p.Run(); err != nil {
		// 	log.Fatal(err)
		// }

		_, err := feature.AddPassword(structs.PasswordData{Id: 1, Name: args[0], Ciphertext: args[1]}, key)
		if err != nil {
			fmt.Printf("error while adding a new password %s", err)
			os.Exit(1)
		}
		fmt.Printf("%s password successfully added.", args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
