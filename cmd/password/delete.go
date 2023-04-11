/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package password

import (
	"encoding/json"
	"fmt"
	"os"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete file with password",
	Long: `Delete file with password`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteCredential(name)
	},
}

func init() {
	PasswordCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&name, "name", "n", "", "name of your credentials")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteCredential(name string) {
	var creds utils.Storage
	
	data, err1 := os.ReadFile("credentials.json")
	if err1 != nil {
		panic(err1)
	}

	err := json.Unmarshal(data, &creds)
	if err != nil {
		panic(err)
	}

	for i :=0; i < len(creds.Credentials); i++ {
		if creds.Credentials[i].Name == name {
			creds.Credentials = append(creds.Credentials[:i],creds.Credentials[i+1:]... )
		} else if len(creds.Credentials) == 0{
			fmt.Println("There are no credentials in file")
			os.Exit(1)
		}else{
			fmt.Println("Wrong credentials name")
			os.Exit(1)
		}
	}

	finalJson, err := json.MarshalIndent(creds, "", "\t")
	if err != nil {
		panic(err)
	}
	
	os.WriteFile("credentials.json", finalJson, 0666)
}

