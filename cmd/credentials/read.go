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

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read all credentials from DB",
	Long: `Read all credentials from DB`,
	Run: func(cmd *cobra.Command, args []string) {
		ReadFromDB()
	},
}

func init() {
	CredentialsCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func ReadPasswordFromFile() {

// 	data, err1 := os.ReadFile("credentials.json")
// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	var creds utils.User

// 	err := json.Unmarshal(data, &creds)
// 	if err != nil {
// 		panic(err1)
// 	}

// 	for _, v := range creds.Credential {
// 			fmt.Printf("Name : %s\nDomain : %s\nLogin : %s\nPassword : %s\n\n", v.CredName, v.Domain, v.Login, v.Password)
// 	}
// }

func ReadFromDB() {
userID := auth.CheckLogin()
db := utils.ConnectToDB()

creds := []utils.Credential{}

resultCreds := db.Select("id", "user_id", "cred_name", "domain", "login", "password").Where("user_id = ?", userID).Find(&creds)
  if resultCreds.Error != nil{
    panic("Error ocured while searching creds")
  }

  if len(creds) == 0{
    fmt.Println("User has no credentials")
    os.Exit(1)
  }

for _, cred := range creds{
	fmt.Printf("ID: %v\n UserID: %v\n CredName: %v\n Domain: %v\n Login: %v\n Password: %v\n\n", cred.ID, cred.UserID, cred.CredName, cred.Domain, cred.Login, cred.Password)
	}

}