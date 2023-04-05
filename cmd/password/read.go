/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package password

import (
	"fmt"
	"io/ioutil"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read password from file",
	Long: `Read password from file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("read called")
	},
}

func init() {
	PasswordCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ReadPasswordFromFile() {
	var filename string

	fmt.Println("Enter name of file with the password:")
	fmt.Scan(&filename)
	bs, _ := ioutil.ReadFile(filename)
	password := string(bs)

	fmt.Printf("Your password is : %v ", password) 
}