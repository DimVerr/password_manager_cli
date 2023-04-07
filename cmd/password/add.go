/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package password

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add credentials to file",
	Long: `add credentials to file`,
	Run: func(cmd *cobra.Command, args []string) {
		EncodeJson(domain, login, password)	
	},
}
var (
	domain string
	login string
	password string
)
type credential struct {
	Domain string `json:"domain"`
	Login string `json:"login" `
	Password string `json:"password"`
}

type storage struct {
	Author string `json:"author"`
	Credentials []credential `json:"credentials"`
}

func init() {
	addCmd.Flags().StringVarP(&domain, "domain", "d", "", "user domain")
	addCmd.Flags().StringVarP(&login, "login", "l", "", "user login")
	addCmd.Flags().StringVarP(&password, "password", "p", "", "user password")
	
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

	data, _ := os.ReadFile("credentials.json")

	var creds storage
	err := json.Unmarshal(data, &creds)

	if err != nil {
		panic(err)
	}

	userCredentials := credential{
		domain, login, password,
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
