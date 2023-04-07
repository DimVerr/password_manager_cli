/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package backend

import (
	"os"
	"github.com/spf13/cobra"
	"encoding/json"
)

var (
	fileName string
	author string
)

type credential struct {
	Domain string `json:"domain"`
	Login string `json:"login" `
	Password string `json:"password"`
}

type storage struct {
	Author string `json:"author"`
	Credentials []credential `json:"credentials"`
}

// intCmd represents the int command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "file for your credentials",
	Long: `Json file for your credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		createFile()	
	},
}

func init() {

	initCmd.Flags().StringVarP(&fileName, "filename", "f", "", "file name for your credentials")
	initCmd.Flags().StringVarP(&author, "author", "a", "", "author of the file")


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

	userCredentials := storage{}
	
	userCredentials.Author = author

	finalJson, err := json.MarshalIndent(userCredentials, "", "\t")
	if err != nil {
		panic(err)
	}

	if fileName != "" {
		os.WriteFile(fileName + ".json", finalJson, 0666)
		}else{
		os.WriteFile("credentials.json", finalJson, 0666)
	}

}



