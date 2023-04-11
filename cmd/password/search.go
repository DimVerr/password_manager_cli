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


// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search credentials in file",
	Long: `Search credentials in file`,
	Run: func(cmd *cobra.Command, args []string) {
		searchCredentials(domain , login, password)	
	},
}

func init() {
	PasswordCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&name, "name", "n", "", "name for credentials")
	searchCmd.Flags().StringVarP(&domain, "domain", "d", "", "user domain")
	
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func searchCredentials(domain string, login string , password string) {
	var creds utils.Storage

	data, err1 := os.ReadFile("credentials.json")
	if err1 != nil {
		panic(err1)
	}

	err := json.Unmarshal(data, &creds)
	if err != nil {
		panic(err)
	}

	for _, v := range creds.Credentials {
		if v.Name == name || v.Domain == domain {
			fmt.Printf("Name : %s\nDomain : %s\nLogin : %s\nPassword : %s\n\n", v.Name, v.Domain, v.Login, v.Password)
		}
	}

}