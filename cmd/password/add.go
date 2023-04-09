/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package password

import (
	"encoding/json"
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
		EncodeJson(utils.Domain, utils.Login, utils.Password)	
	},
}

func init() {
	addCmd.Flags().StringVarP(&utils.Domain, "domain", "d", "", "user domain")
	addCmd.Flags().StringVarP(&utils.Login, "login", "l", "", "user login")
	addCmd.Flags().StringVarP(&utils.Password, "password", "p", "", "user password")
	
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


func EncodeJson(domain string, login string, password string) {

	data, err1 := os.ReadFile("credentials.json")
	if err1 != nil {
		panic(err1)
	}

	var creds utils.Storage
	err := json.Unmarshal(data, &creds)

	if err != nil {
		panic(err)
	}

	userCredentials := utils.Credential{
		Domain: domain, Login: login, Password: password,
	}

	creds.Credentials = append(creds.Credentials, userCredentials)
	
	finalJson, err := json.MarshalIndent(creds, "", "\t")
	if err != nil {
		panic(err)
	}
	
	os.WriteFile("credentials.json", finalJson, 0666)
	
	// finalJson, err := json.MarshalIndent(userCredentials, "", "\t")
	// if err != nil {
	// 	panic(err)
	// }
	
	// os.WriteFile("credentials.json", finalJson, 0666)

}
