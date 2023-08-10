package auth

import (
	"fmt"
	"os"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Sign up procedure to create User.",
	Long: `Sign up procedure to create User.`,
	Run: func(cmd *cobra.Command, args []string) {
		SignUp()
	},
}

func init() {
	AuthCmd.AddCommand(signupCmd)
	signupCmd.Flags().StringVarP(&userName, "name", "n", "", "name of your user")
	signupCmd.Flags().StringVarP(&userPassword, "password", "p", "", "your password")
	signupCmd.MarkFlagRequired("name")
	signupCmd.MarkFlagRequired("password")
}

func SignUp() {
	db := utils.ConnectToDB()
	if userName == "" || userPassword ==""{
		fmt.Println("Insert name and password as flags.")
		os.Exit(1)
	}
	resultSignUp := db.Create(&utils.User{Name: userName, Password: userPassword})
	if resultSignUp.Error != nil {
		fmt.Println("Username is already used.")
		os.Exit(1)
	}
	fmt.Println("You are successfulyy signed up. Run `login` command to be able to interact with another commands.")
}