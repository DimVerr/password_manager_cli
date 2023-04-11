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


// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit credentials in file",
	Long: `Edit credentials in file`,
	Run: func(cmd *cobra.Command, args []string) {
		editCredentials(name, domain, login, password, newName, newDomain, newLogin, newPassword)
	},
}

func init() {
	PasswordCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&name, "name", "n", "", "name for credentials")
	editCmd.Flags().StringVarP(&domain, "domain", "d", "", "user domain")
	editCmd.Flags().StringVarP(&login, "login", "l", "", "user login")
	editCmd.Flags().StringVarP(&password, "password", "p", "", "user password")
	editCmd.Flags().StringVarP(&newName, "newname", "", "", "new name for credentials")
	editCmd.Flags().StringVarP(&newDomain, "newdomain", "", "", "new user domain")
	editCmd.Flags().StringVarP(&newLogin, "newlogin", "", "", "new user login")
	editCmd.Flags().StringVarP(&newPassword, "newpassword", "", "", "new user password")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func editCredentials(name string, domain string, login string , password string, newname string, newdomain string, newlogin string, newpassword string) {
	var Creds utils.Storage

	data, err1 := os.ReadFile("credentials.json")
	if err1 != nil {
		panic(err1)
	}

	err := json.Unmarshal(data, &Creds)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(Creds.Credentials); i++ {
		if newname == Creds.Credentials[i].Name{
			fmt.Println("Name is already used")
			os.Exit(1)
		}else {
			continue
		}
	}


	for i := 0; i < len(Creds.Credentials); i++ {
		if Creds.Credentials[i].Name == name && newname != ""{
			Creds.Credentials[i].Name = newname			
		}else {
			continue
		}
	}

	for i := 0; i < len(Creds.Credentials); i++ {
		if Creds.Credentials[i].Name == name && newdomain != "" && domain !=""{
			Creds.Credentials[i].Domain = newdomain			
		}else {
			continue
		}
	}

	for i := 0; i < len(Creds.Credentials); i++ {
		if Creds.Credentials[i].Name == name && newlogin != "" && login !=""{
			Creds.Credentials[i].Login = newlogin		
		}else {
			continue
		}
	}

	for i := 0; i < len(Creds.Credentials); i++ {
		if Creds.Credentials[i].Name == name && newpassword != "" && password != ""{
			Creds.Credentials[i].Password = newpassword			
		}else {
			continue
		}
	}
	
	finalJson, err := json.MarshalIndent(Creds, "", "\t")
	if err != nil {
		panic(err)
	}

	os.WriteFile("credentials.json", finalJson, 0666)
}