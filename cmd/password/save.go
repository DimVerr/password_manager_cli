/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package password

import (
	"fmt"
	"io/ioutil"
	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save password to file",
	Long: `Save password to file`,
	Run: func(cmd *cobra.Command, args []string) {
		savePasswordToFile()	
	},
}

func init() {
	PasswordCmd.AddCommand(saveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func savePasswordToFile() {
	var password string
	var fileName string

	fmt.Println("Enter your password:")
	fmt.Scan(&password)

	fmt.Println("Enter file name:")
	fmt.Scan(&fileName)

	ioutil.WriteFile(fileName, []byte(password), 0666)
}
