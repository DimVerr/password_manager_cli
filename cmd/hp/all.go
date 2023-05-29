/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package hp

import (
	"encoding/json"
	"net/http"
	"os"
	"password_manager/cmd/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		allCredentials() 
	},
}

func init() {
	HpCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func allCredentials() {
    router := gin.Default()
    router.GET("/all", getCreds)
	// router.GET("/byname/:name", getCredsbyName)
    router.Run("localhost:8080")

}

func getCreds(c *gin.Context) {
	var creds utils.Storage

	data, err1 := os.ReadFile("credentials.json")
	if err1 != nil {
		panic(err1)
	}
	err := json.Unmarshal(data, &creds)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, creds.Credentials)
}

// func getCredsbyName(c *gin.Context) {
// 	name := c.Param("name")
// 	var creds utils.Storage

// 	data, err1 := os.ReadFile("credentials.json")
// 	if err1 != nil {
// 		panic(err1)
// 	}
// 	err := json.Unmarshal(data, &creds)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, v := range creds.Credentials {
// 		if v.Name == name {
// 			c.IndentedJSON(http.StatusOK, v)
// 			fmt.Printf("Name : %s\nDomain : %s\nLogin : %s\nPassword : %s\n\n", v.Name, v.Domain, v.Login, v.Password)
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "creds not found"})
// }
