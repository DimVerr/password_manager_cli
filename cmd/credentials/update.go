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

// editCmd represents the edit command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update credentials in DB",
	Long: `update credentials in DB`,
	Run: func(cmd *cobra.Command, args []string) {
		UpdateCreds()
	},
}

func init() {
	CredentialsCmd.AddCommand(updateCmd)

	updateCmd.Flags().UintVarP(&credID, "credID", "i", credID, "cred ID for searching")
	// editCmd.Flags().StringVarP(&domain, "domain", "d", "", "user domain")
	// editCmd.Flags().StringVarP(&login, "login", "l", "", "user login")
	// editCmd.Flags().StringVarP(&password, "password", "p", "", "user password")
	updateCmd.Flags().StringVarP(&newCredName, "newCredName", "n", "", "new name for credentials")
	updateCmd.Flags().StringVarP(&newDomain, "newDomain", "d", "", "new user domain")
	updateCmd.Flags().StringVarP(&newLogin, "newLogin", "l", "", "new user login")
	updateCmd.Flags().StringVarP(&newPassword, "newPassword", "p", "", "new user password")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func editCredentials(name string, domain string, login string , password string, newCredName string, newdomain string, newlogin string, newpassword string) {
// 	var Creds utils.User

// 	data, err1 := os.ReadFile("credentials.json")
// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	err := json.Unmarshal(data, &Creds)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i := 0; i < len(Creds.Credential); i++ {
// 		if newname == Creds.Credential[i].CredName{
// 			fmt.Println("Name is already used")
// 			os.Exit(1)
// 		}else {
// 			continue
// 		}
// 	}


// 	for i := 0; i < len(Creds.Credential); i++ {
// 		if Creds.Credential[i].CredName == name && newname != ""{
// 			Creds.Credential[i].CredName = newname			
// 		}else {
// 			continue
// 		}
// 	}

// 	for i := 0; i < len(Creds.Credential); i++ {
// 		if Creds.Credential[i].CredName == name && newdomain != "" && domain !=""{
// 			Creds.Credential[i].Domain = newdomain			
// 		}else {
// 			continue
// 		}
// 	}

// 	for i := 0; i < len(Creds.Credential); i++ {
// 		if Creds.Credential[i].CredName == name && newlogin != "" && login !=""{
// 			Creds.Credential[i].Login = newlogin		
// 		}else {
// 			continue
// 		}
// 	}

// 	for i := 0; i < len(Creds.Credential); i++ {
// 		if Creds.Credential[i].CredName == name && newpassword != "" && password != ""{
// 			Creds.Credential[i].Password = newpassword			
// 		}else {
// 			continue
// 		}
// 	}
	
// 	finalJson, err := json.MarshalIndent(Creds, "", "\t")
// 	if err != nil {
// 		panic(err)
// 	}

// 	os.WriteFile("credentials.json", finalJson, 0666)
// }

func UpdateCreds(){
	userID := auth.CheckLogin()
	db := utils.ConnectToDB()
	
	if credID == 0{
		fmt.Println("Please add credID as flag. You can find all your cred IDs with `read` command.")
		os.Exit(1)
	}
	
	if newCredName != "" {
		db.Model(utils.Credential{}).Where("user_id = ? AND id = ?", userID, credID).Update("cred_name", newCredName)
	}

	if newDomain != "" {
		db.Model(utils.Credential{}).Where("user_id = ? AND id = ?", userID, credID).Update("domain", newDomain)
	}

	if newLogin != "" {
		db.Model(utils.Credential{}).Where("user_id = ? AND id = ?", userID, credID).Update("login", newLogin)	
		}
	
	if newPassword != "" {
		db.Model(utils.Credential{}).Where("user_id = ? AND id = ?", userID, credID).Update("password", newPassword)
	}
	
	fmt.Println("Your credentials were updated successfully")
}