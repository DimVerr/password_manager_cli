package creds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	credID uint
	credName string
	domain string
	login string
	password string
	newCredName string
	newDomain string
	newLogin string
	newPassword string
)

var CredsCmd = &cobra.Command{
	Use:   "creds",
	Short: "Creds commands (create, update, delete, getAll, getOne)",
	Long: `Creds commands (create, update, delete, getAll, getOne)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Call one of following commands: `create`, `update`, `delete`, `getAll` or `getOne`.")
	},
}