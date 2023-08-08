package creds

import (
	"fmt"
	"os"
	"password_manager/cmd/auth"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

var getOneCmd = &cobra.Command{
	Use:   "getOne",
	Short: "Get one credentials from DB",
	Long: `Get one credentials from DB`,
	Run: func(cmd *cobra.Command, args []string) {
		SearchInDb()	
	},
}

func init() {
	CredsCmd.AddCommand(getOneCmd)

	getOneCmd.Flags().UintVarP(&credID, "credID", "i", credID, "cred ID for searching")
	
	getOneCmd.MarkFlagRequired("credID")
}

func SearchInDb(){
userID := auth.CheckLogin()
db := utils.ConnectToDB()

cred := utils.Credential{}

resultSearch := db.Select("id", "user_id", "cred_name", "domain", "login", "password").Where("user_id = ? AND id = ?", userID, credID).First(&cred)
if resultSearch.Error != nil {
	fmt.Println("Wrong credID")
	os.Exit(1)
}

fmt.Printf("ID: %v\n UserID: %v\n CredName: %v\n Domain: %v\n Login: %v\n Password: %v\n\n", cred.ID, cred.UserID, cred.CredName, cred.Domain, cred.Login, cred.Password)

}