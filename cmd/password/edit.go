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
		editCredentials(utils.Domain, utils.Login, utils.Password, utils.NewDomain, utils.NewLogin, utils.NewPassword)
	},
}

func init() {
	PasswordCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&utils.Domain, "domain", "d", "", "user domain")
	editCmd.Flags().StringVarP(&utils.Login, "login", "l", "", "user login")
	editCmd.Flags().StringVarP(&utils.Password, "password", "p", "", "user password")
	editCmd.Flags().StringVarP(&utils.NewDomain, "newdomain", "a", "", "new user domain")
	editCmd.Flags().StringVarP(&utils.NewLogin, "newlogin", "b", "", "new user login")
	editCmd.Flags().StringVarP(&utils.NewPassword, "newpassword", "c", "", "new user password")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func editCredentials(domain string, login string , password string, newdomain string, newlogin string, newpassword string) {
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
		if Creds.Credentials[i].Domain == domain && Creds.Credentials[i].Login == login && Creds.Credentials[i].Password == password{
			Creds.Credentials[i].Domain = newdomain
			Creds.Credentials[i].Login = newlogin
			Creds.Credentials[i].Password = newpassword
		}else {
			fmt.Println("Wrong credentials")
			break
		}
	}

	// for _,v := range creds.Credentials{
	// 	if v.Domain == domain {
	// 		v.Domain = newdomain
	// 	}else {
	// 		fmt.Println("Domain does not exist")			
	// 	}
	// }

	// for i := 0; i < len(creds.Credentials); i++ {
	// 	if creds.Credentials[i].Login == login {
	// 		creds.Credentials[i].Login = newlogin
	// 	}
	// }

	// for i := 0; i < len(creds.Credentials); i++ {
	// 	if creds.Credentials[i].Password == password{
	// 		creds.Credentials[i].Password = newpassword
	// 	}
	// }
	
	finalJson, err := json.MarshalIndent(Creds, "", "\t")
	if err != nil {
		panic(err)
	}

	os.WriteFile("credentials.json", finalJson, 0666)
}