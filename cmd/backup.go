package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var archiver string

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backups files immediately to specified storages",
	Long:  `long descr: backups files immediately`,
	Run: func(cmd *cobra.Command, _ []string) {
		g, _ := cmd.Flags().GetBool("gdrive")
		y, _ := cmd.Flags().GetBool("yadisk")
		t, _ := cmd.Flags().GetBool("telegram")
		e, _ := cmd.Flags().GetBool("email")

		if archiver == "zip" {
			fmt.Printf("Using %s archiver\n", archiver)
		} else if archiver == "tar" {
			fmt.Printf("Using %s archiver\n", archiver)
		} else {
			log.Fatal("Wrong archiver type (zip and tar supported)")
		}

		if g {
			fmt.Println("Works gdrive")
		}

		if y {
			// storage := clouds.NewYaDisk()
			// storage.DownLoadFile("1", "2")
			fmt.Println("Works yandex")
		}

		if t {
			fmt.Println("Works telegram")
		}

		if e {
			fmt.Println("Works email")
		}

	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().BoolP("gdrive", "g", false, "Upload backup archive to Google Drive")
	backupCmd.Flags().BoolP("yadisk", "y", false, "Upload backup archive to Yandex Disk")
	backupCmd.Flags().BoolP("telegram", "t", false, "Sends backup archive to Telegram chat/channel")
	backupCmd.Flags().BoolP("email", "e", false, "Sends backup archive via email")

	backupCmd.Flags().StringVarP(&archiver, "archiver", "a", "", "Use zip / tar")
	backupCmd.MarkFlagRequired("archiver")
}
