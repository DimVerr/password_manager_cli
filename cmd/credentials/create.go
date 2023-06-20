/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package credentials

import (
	"fmt"
	"os"
	"password_manager/cmd/auth"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create your credentials",
	Long: `Create your credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		AddCredsToDB()	
	},
}

func init() {
	createCmd.Flags().StringVarP(&credName, "credName", "c", "", "name for credentials")
	createCmd.Flags().StringVarP(&domain, "domain", "d", "", "domain")
	createCmd.Flags().StringVarP(&login, "login", "l", "", "login")
	createCmd.Flags().StringVarP(&password, "password", "p", "", "password")

	createCmd.MarkFlagRequired("credName")
	createCmd.MarkFlagRequired("domain")
	createCmd.MarkFlagRequired("login")
	createCmd.MarkFlagRequired("password")
	
	CredentialsCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// func EncodeJson(name string, domain string, login string, password string) {
	
// 	data, err1 := os.ReadFile("credentials.json")
// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	var creds utils.Storage
// 	err := json.Unmarshal(data, &creds)

// 	if err != nil {
// 		panic(err)
// 	}

// 	for _,v := range creds.Credentials{
// 		if v.Name == name {
// 			fmt.Println("Name is already used")
// 			os.Exit(1)
// 		}
// 	}

// 	userCredentials := utils.Credential{
// 		Name: name, Domain: domain, Login: login, Password: password,
// 	}

// 	creds.Credentials = append(creds.Credentials, userCredentials)
	
// 	finalJson, err := json.MarshalIndent(creds, "", "\t")
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	os.WriteFile("credentials.json", finalJson, 0666)

// }

func AddCredsToDB() {
	userID := auth.CheckLogin() 
	db := utils.ConnectToDB()

	resultAdd := db.Create(&utils.Credential{UserID: userID, CredName: credName, Domain: domain, Login: login, Password: password})
	if resultAdd.Error != nil {
		fmt.Println("Impossible to create credentials")
		os.Exit(1)
	}
	
	fmt.Println("Your credentials were added successfully")
}
