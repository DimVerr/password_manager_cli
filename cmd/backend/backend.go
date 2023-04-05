/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package backend

import (
	"fmt"

	"github.com/spf13/cobra"
)

// backendCmd represents the backend command
var BackendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Backend stuff",
	Long: `Backend stuff`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("backend called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
