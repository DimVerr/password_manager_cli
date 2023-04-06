/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package backend

import (
	"os"
	"github.com/spf13/cobra"
)

var (
	fileName string
)



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
	os.Create("credentials.json")
		if fileName != "" {
			createFileWithCustomName(fileName)
		}
}

func createFileWithCustomName(fileName string) {
	os.Create(fileName+".json")
}	
