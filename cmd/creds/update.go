package creds

import (
	"fmt"
	"os"
	"password_manager/cmd/auth"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update credentials in DB",
	Long: `update credentials in DB`,
	Run: func(cmd *cobra.Command, args []string) {
		UpdateCreds()
	},
}

func init() {
	CredsCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&newCredName, "newCredName", "n", "", "new name for credentials")
	updateCmd.Flags().StringVarP(&newDomain, "newDomain", "d", "", "new user domain")
	updateCmd.Flags().StringVarP(&newLogin, "newLogin", "l", "", "new user login")
	updateCmd.Flags().StringVarP(&newPassword, "newPassword", "p", "", "new user password")
	updateCmd.Flags().UintVarP(&credID, "credID", "i", credID, "cred ID for searching")

	updateCmd.MarkFlagRequired("credID")
	updateCmd.MarkFlagRequired("newCredName")
	updateCmd.MarkFlagRequired("newDomain")
	updateCmd.MarkFlagRequired("newLogin")
	updateCmd.MarkFlagRequired("newPassword")
}

func UpdateCreds(){
	userID := auth.CheckLogin()
	db := utils.ConnectToDB()
	
	if credID == 0{
		fmt.Println("Please add credID as flag. You can find all your cred IDs with `read` command.")
		os.Exit(1)
	}

	newCreds := utils.Credential{CredName: newCredName, Domain: newDomain, Login: newLogin, Password: newPassword}
	
	result := db.Model(utils.Credential{}).Where("user_id = ? and id = ?",userID, credID).Updates(&newCreds)
	if result.RowsAffected == 0 {
		fmt.Println("No credentials with this credential ID were found.")
		os.Exit(1)
	}
	
	fmt.Println("Your credentials were updated successfully.")
}