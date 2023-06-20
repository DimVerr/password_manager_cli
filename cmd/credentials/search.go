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

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search credentials in DB",
	Long: `Search credentials in DB`,
	Run: func(cmd *cobra.Command, args []string) {
		SearchInDb()	
	},
}

func init() {
	CredentialsCmd.AddCommand(searchCmd)
	searchCmd.Flags().UintVarP(&credID, "credID", "i", credID, "cred ID for searching")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func searchCredentials(name string, domain string) {
// 	var creds utils.User

// 	data, err1 := os.ReadFile("credentials.json")
// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	err := json.Unmarshal(data, &creds)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, v := range creds.Credential {
// 		if v.CredName == name || v.Domain == domain {
// 			fmt.Printf("Name : %s\nDomain : %s\nLogin : %s\nPassword : %s\n\n", v.CredName, v.Domain, v.Login, v.Password)
// 		}
// 	}

// }

// func SearchCredsFromDB(name string, domain string) {
// 	var creds utils.Credential
	
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 	  panic("failed to connect database")
// 	}
	
// 	db.Where("name = ? OR domain >= ?", name , domain).Find(&creds)
// 	fmt.Printf("Name : %s\nDomain : %s\nLogin : %s\nPassword : %s\n\n", creds.CredName, creds.Domain, creds.Login, creds.Password)
// }

func SearchInDb(){
userID := auth.CheckLogin()
db := utils.ConnectToDB()

if credID == 0{
	fmt.Println("Please add credID as flag. You can find all your cred IDs with `read` command.")
	os.Exit(1)
}

cred := utils.Credential{}

resultSearch := db.Select("id", "user_id", "cred_name", "domain", "login", "password").Where("user_id = ? AND id = ?", userID, credID).First(&cred)
if resultSearch.Error != nil {
	fmt.Println("Wrong credID")
	os.Exit(1)
}
fmt.Printf("ID: %v\n UserID: %v\n CredName: %v\n Domain: %v\n Login: %v\n Password: %v\n\n", cred.ID, cred.UserID, cred.CredName, cred.Domain, cred.Login, cred.Password)

}