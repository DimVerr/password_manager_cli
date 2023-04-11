/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package utils

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Credential struct {
	Name string `json:"name"`
	Domain string `json:"domain"`
	Login string `json:"login" `
	Password string `json:"password"`
}

type Storage struct {
	Author string `json:"author"`
	Credentials []Credential `json:"credentials"`
}

// utilsCmd represents the utils command
var UtilsCmd = &cobra.Command{
	Use:   "utils",
	Short: "All global variables are set here",
	Long: `All global variables are set here`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("utils called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// utilsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// utilsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
