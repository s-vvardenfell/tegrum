package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tergum",
	Short: "go-written file backuper",
	Long:  `long descr: go-written file backuper`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("execute: %v", err)
	}
}
