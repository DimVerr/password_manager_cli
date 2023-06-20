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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete credentials from DB",
	Long: `Delete credentials from DB`,
	Run: func(cmd *cobra.Command, args []string) {
		DeleteFromDB()
	},
}

func init() {
	CredentialsCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().UintVarP(&credID, "credID", "i", credID, "cred ID for searching")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func deleteCredential(name string) {
// 	var creds utils.User
	
// 	data, err1 := os.ReadFile("credentials.json")
// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	err := json.Unmarshal(data, &creds)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i :=0; i < len(creds.Credential); i++ {
// 		if creds.Credential[i].CredName == name {
// 			creds.Credential = append(creds.Credential[:i],creds.Credential[i+1:]... )
// 		} else if len(creds.Credential) == 0{
// 			fmt.Println("There are no credentials in file")
// 			os.Exit(1)
// 		}else{
// 			fmt.Println("Wrong credentials name")
// 			os.Exit(1)
// 		}
// 	}

// 	finalJson, err := json.MarshalIndent(creds, "", "\t")
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	os.WriteFile("credentials.json", finalJson, 0666)
// }

func DeleteFromDB() {
	userID := auth.CheckLogin()
	db := utils.ConnectToDB()

	if credID == 0{
		fmt.Println("Please add credID as flag. You can find all your cred IDs with `read` command.")
		os.Exit(1)
	}
	
	var creds = utils.Credential{}
	resultDelete := db.Where("user_id = ? AND id = ?", userID, credID).First(&creds)
	if resultDelete.Error != nil {
		fmt.Println("Wrong credID")
		os.Exit(1)
	}

	db.Unscoped().Delete(creds)
	fmt.Println("Your creds were deleted successfully")
}
