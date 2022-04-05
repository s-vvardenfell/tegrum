package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Gets backup archives stored earlier from specified storages",
	Long:  `long descr: backups files immediately`,
	Run: func(cmd *cobra.Command, _ []string) {
		g, _ := cmd.Flags().GetBool("gdrive")
		y, _ := cmd.Flags().GetBool("yadisk")
		t, _ := cmd.Flags().GetBool("telegram")

		if !g && !y && !t {
			log.Fatal("No source specified")
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

		//возможно, нужно исп if-else и если нет
	},
}

func init() {
	rootCmd.AddCommand(retrieveCmd)
	retrieveCmd.Flags().BoolP("gdrive", "g", false, "Download last backup archive from Google Drive")
	retrieveCmd.Flags().BoolP("yadisk", "y", false, "Download last backup archive from Yandex Disk")
	retrieveCmd.Flags().BoolP("telegram", "t", false, "Download last backup archive from Telegram")
}
