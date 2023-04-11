/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package password

import (
	"encoding/json"
	"fmt"
	"os"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read password from file",
	Long: `Read password from file`,
	Run: func(cmd *cobra.Command, args []string) {
		ReadPasswordFromFile()
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

	data, err1 := os.ReadFile("credentials.json")
	if err1 != nil {
		panic(err1)
	}

	var creds utils.Storage

	err := json.Unmarshal(data, &creds)
	if err != nil {
		panic(err1)
	}

	for _, v := range creds.Credentials {
			fmt.Printf("Name : %s\nDomain : %s\nLogin : %s\nPassword : %s\n\n", v.Name, v.Domain, v.Login, v.Password)
	}
}