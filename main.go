package main

import (
	"github.com/spf13/cobra"
	"os/user"
)

var username, password, host string
var configPath string

var rootCmd = &cobra.Command{
	Use:   "rsynker",
	Short: "rsynker is a friendly rsync command line tool",
	Long:  `rsynker is a friendly rsync command line tool`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	// Be careful not to use the duplicated shorthand flag to the subcommand
	// This will cause go panic
	// For example, use "P" for password and "p" for port
	currentUser, _ := user.Current()

	rootCmd.PersistentFlags().StringVarP(&configPath, "conf", "c", "/etc/rsyncd.conf", "config of rsync server")
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", currentUser.Username, "username for rsync")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "P", "", "password for rsync")
	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "localhost", "rsync server host")

	rootCmd.AddCommand(cmdServer)
	rootCmd.AddCommand(cmdAddModule)
	rootCmd.AddCommand(cmdDelModule)
	rootCmd.Execute()
}
