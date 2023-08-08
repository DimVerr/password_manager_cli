package creds

import (
	"fmt"
	"os"
	"password_manager/cmd/auth"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

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
	
	CredsCmd.AddCommand(createCmd)
}

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
