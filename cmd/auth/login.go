package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login command to be able to interact with another commands.",
	Long: `Login command to be able to interact with another commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		Login()	
	},
}

func init() {
	AuthCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&userName, "name", "n", "", "name of your user")
	loginCmd.Flags().StringVarP(&userPassword, "password", "p", "", "your password")
	loginCmd.MarkFlagRequired("name")
	loginCmd.MarkFlagRequired("password")
}

func Login() {
	db := utils.ConnectToDB()
	resultName := db.First(&utils.User{}, "name = ?", userName)
	if resultName.Error != nil {
	  fmt.Println("Invalid username.")

	  os.Exit(1)
	} 
  
	resultPassword := db.First(&utils.User{}, "password = ?", userPassword).Where("name =?", userName)
	if resultPassword.Error != nil {
	  fmt.Println("Invalid password.")

	  os.Exit(1)
	}
	db.Model(&utils.User{}).Where("name =? and password =?", userName, userPassword).Pluck("id", &userID)
	UserRefer, err := json.MarshalIndent(userID, "", "/t")
	if err != nil {
		fmt.Println("Impossible to extract user ID.")
	}

	err1 := os.WriteFile("UserID", UserRefer, 0644)
	if err1 != nil{
		fmt.Println("Impossible to extract user ID.")
	}
	fmt.Println("You are sucessfully logged in.")

}

func CheckLogin() uint{

	uid, err := os.ReadFile("UserID") 
	if err != nil {
		fmt.Println("You are not logged in. Please log in with `login` command.")

		os.Exit(1)
	}else {
		err := json.Unmarshal(uid, &userID)
		if err != nil {
			fmt.Println("Impossible to extract UserID.")
			os.Exit(1)
		}
	}

	return userID
}