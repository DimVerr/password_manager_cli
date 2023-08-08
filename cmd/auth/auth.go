package auth

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	userName string
	userPassword string
	userID uint
)
  
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorization commands are set here",
	Long: `Authorization commands are set here (login and signup)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Call one of following commands: `signup` or `login`")
	},
}

func init() {

}
