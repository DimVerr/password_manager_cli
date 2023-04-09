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
		searchCredentials(utils.Domain , utils.Login, utils.Password)	
	},
}

func init() {
	PasswordCmd.AddCommand(searchCmd)


	searchCmd.Flags().StringVarP(&utils.Domain, "domain", "d", "", "user domain")
	searchCmd.Flags().StringVarP(&utils.Login, "login", "l", "", "user login")
	searchCmd.Flags().StringVarP(&utils.Password, "password", "p", "", "user password")
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
		if v.Domain == utils.Domain || v.Login == utils.Login || v.Password == utils.Password {
			fmt.Printf("%v\n", v)
		}
	}


// 	for i := 0; i < len(creds.Credentials); i++ {
// 		if creds.Credentials[i].Domain == utils.Domain || creds.Credentials[i].Login == utils.Login ||creds.Credentials[i].Password == utils.Password {
// 			fmt.Println(creds.Credentials[i])
// 		}
// }
}