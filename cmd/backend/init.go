/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package backend

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"password_manager/cmd/utils"
)



// intCmd represents the int command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "file for your credentials",
	Long: `Json file for your credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		createExistingFile()	
	},
}

func init() {

	initCmd.Flags().StringVarP(&utils.FileName, "filename", "f", "", "file name for your credentials")
	initCmd.Flags().StringVarP(&utils.Author, "author", "a", "", "author of the file")
	initCmd.Flags().StringVarP(&utils.Create, "create", "c", "", "insert 1 to the flag argument to create file in any way")


	// Here you will define your flags and configuration settings.
	BackendCmd.AddCommand(initCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// intCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// intCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createFile() {

	userCredentials := utils.Storage{}
	
	userCredentials.Author = utils.Author

	finalJson, err := json.MarshalIndent(userCredentials, "", "\t")
	if err != nil {
		panic(err)
	}

	if utils.FileName != "" {
		os.WriteFile(utils.FileName + ".json", finalJson, 0666)
		}else{
		os.WriteFile("credentials.json", finalJson, 0666)
	}

}

func createExistingFile() {
	_ , err := os.ReadFile("credentials.json")
	if err != nil {
		createFile()
	}else if err == nil && utils.Create == "1"{
		createFile()
	}else {
		fmt.Println("File is already existing")
		os.Exit(1)
	}
}


