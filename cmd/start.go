package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mds1455975151/gin-demo/server"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Print the version number of gin-demo",
	Long:  `All software has versions, This is gin-demo's'`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}