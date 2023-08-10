package creds

import (
	"fmt"
	"os"
	"password_manager/cmd/auth"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete credentials from DB",
	Long: `Delete credentials from DB`,
	Run: func(cmd *cobra.Command, args []string) {
		DeleteFromDB()
	},
}

func init() {
	CredsCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().UintVarP(&credID, "credID", "i", credID, "cred ID for searching")
	
	deleteCmd.MarkFlagRequired("credID")
}

func DeleteFromDB() {
	userID := auth.CheckLogin()
	db := utils.ConnectToDB()
	
	var creds = utils.Credential{}

	result := db.Unscoped().Where("user_id = ? AND id = ?", userID, credID).Delete(creds)
	if result.RowsAffected == 0 {
		fmt.Println("No credentials with this credential ID were found.")
		os.Exit(1)
	}
	fmt.Println("Your creds were deleted successfully.")
}
