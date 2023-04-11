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

// saveCmd represents the save command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add credentials to file",
	Long: `add credentials to file`,
	Run: func(cmd *cobra.Command, args []string) {
		EncodeJson(name, domain, login, password)	
	},
}

func init() {
	addCmd.Flags().StringVarP(&name, "name", "n", "", "name for credentials")
	addCmd.Flags().StringVarP(&domain, "domain", "d", "", "user domain")
	addCmd.Flags().StringVarP(&login, "login", "l", "", "user login")
	addCmd.Flags().StringVarP(&password, "password", "p", "", "user password")

	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("domain")
	addCmd.MarkFlagRequired("login")
	addCmd.MarkFlagRequired("password")
	
	PasswordCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func EncodeJson(name string, domain string, login string, password string) {
	
	data, err1 := os.ReadFile("credentials.json")
	if err1 != nil {
		panic(err1)
	}

	var creds utils.Storage
	err := json.Unmarshal(data, &creds)

	if err != nil {
		panic(err)
	}

	for _,v := range creds.Credentials{
		if v.Name == name {
			fmt.Println("Name is already used")
			os.Exit(1)
		}
	}

	userCredentials := utils.Credential{
		Name: name, Domain: domain, Login: login, Password: password,
	}

	creds.Credentials = append(creds.Credentials, userCredentials)
	
	finalJson, err := json.MarshalIndent(creds, "", "\t")
	if err != nil {
		panic(err)
	}
	
	os.WriteFile("credentials.json", finalJson, 0666)

}
