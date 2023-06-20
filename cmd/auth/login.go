/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Login() {
	db := utils.ConnectToDB()
	resultName := db.First(&utils.User{}, "name = ?", userName)
	if resultName.Error != nil {
	  fmt.Println("Invalid username")
	  os.Exit(1)
	} 
  
	resultPassword := db.First(&utils.User{}, "password = ?", userPassword).Where("name =?", userName)
	if resultPassword.Error != nil {
	  fmt.Println("Invalid password")
	  os.Exit(1)
	}
	db.Model(&utils.User{}).Where("name =? and password =?", userName, userPassword).Pluck("id", &userID)
	UserRefer, err1 := json.MarshalIndent(userID, "", "/t")
	if err1 != nil {
		panic(err1)
	}
	os.WriteFile("UserID", UserRefer, 0644)
	fmt.Println("You are sucessfully logged in")
}

func CheckLogin() uint{

	uid, err2 := os.ReadFile("UserID") 
	if err2 != nil {
		fmt.Println("You are not logged in. Please log in with `login` command")
		os.Exit(1)
	}else {
		err3 := json.Unmarshal(uid, &userID)
		if err3 != nil {
			panic("Impossible to extract UserID from file")
		}
	}
	return userID
}