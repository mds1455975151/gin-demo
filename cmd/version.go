package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gin-demo",
	Long:  `All software has versions, This is gin-demo's'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gin-demo v1.0.0 -- develp(00a2ad5)")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}