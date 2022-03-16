package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Runs as a service/daemon and backups files by schedule",
	Long:  `long descr: runs as a service/daemon and backups files by schedule`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Daemon works")
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
