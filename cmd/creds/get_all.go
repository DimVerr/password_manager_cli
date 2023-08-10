package creds

import (
	"fmt"
	"os"
	"password_manager/cmd/auth"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

var getAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "Get all credentials from DB",
	Long: `Get all credentials from DB`,
	Run: func(cmd *cobra.Command, args []string) {
		ReadFromDB()
	},
}

func init() {
	CredsCmd.AddCommand(getAllCmd)
}

func ReadFromDB() {
userID := auth.CheckLogin()
db := utils.ConnectToDB()

creds := []utils.Credential{}

resultCreds := db.Select("id", "user_id", "cred_name", "domain", "login", "password").Where("user_id = ?", userID).Find(&creds)
if resultCreds.Error != nil{
    fmt.Println("Error ocured while searching creds.")
	os.Exit(1)
}

if len(creds) == 0{
    fmt.Println("User has no credentials.")
    os.Exit(1)
}

for _, cred := range creds{
	fmt.Printf("ID: %v\n UserID: %v\n CredName: %v\n Domain: %v\n Login: %v\n Password: %v\n\n", cred.ID, cred.UserID, cred.CredName, cred.Domain, cred.Login, cred.Password)
}
	
}