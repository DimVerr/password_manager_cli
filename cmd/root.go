package cmd

import (
	"fmt"
	"os"
	"password_manager/cmd/creds"
	"password_manager/cmd/auth"
	"password_manager/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "password_manager",
	Short: "Password manager will help you to store your users and their credentials in easy way",
	Long: `Password manager allows you to create, update, delete and show your credentials. For example:
	Call "password_manager help" command to see available blocks`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(creds.CredsCmd)
	rootCmd.AddCommand(utils.UtilsCmd)
	rootCmd.AddCommand(auth.AuthCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".password_manager")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
