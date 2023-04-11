/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package password

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name string
	domain string
	login string
	password string
	newName string
	newDomain string
	newLogin string
	newPassword string
)

// passwordCmd represents the password command
var PasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "Password",
	Long: `Password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("password called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passwordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passwordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
