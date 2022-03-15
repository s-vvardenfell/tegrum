package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Runs as a service/daemon and backups files by schedule",
	Long:  `long descr: runs as a service/daemon and backups files by schedule`,
	Run: func(cmd *cobra.Command, args []string) {
		sum := 0
		for _, args := range args {
			num, err := strconv.Atoi(args)

			if err != nil {
				fmt.Println(err)
			}
			sum = sum + num
		}
		fmt.Println("result of addition is", sum)
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
