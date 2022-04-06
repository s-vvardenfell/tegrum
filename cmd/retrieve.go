package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/s-vvardenfell/tegrum/clouds/gdrive"
	"github.com/s-vvardenfell/tegrum/clouds/yadisk"
	"github.com/s-vvardenfell/tegrum/records/csv_record"
	"github.com/s-vvardenfell/tegrum/telegram"
	"github.com/spf13/cobra"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Gets backup archives stored earlier from specified storages",
	Long:  `long descr: backups files immediately`,
	Run: func(cmd *cobra.Command, _ []string) {
		//select storage type
		//if not specified, error
		var rr RecorderRetriever
		csvFlag, err := cmd.Flags().GetBool(csv)
		if err != nil {
			log.Fatal(err)
		}

		if csvFlag {
			rr = &csv_record.CsvRecorderRetriever{}
		} else {
			log.Fatal("storage not specified")
		}

		// getting destination dir
		dstDir, err := cmd.Flags().GetString(destinationDir)
		if err != nil || dstDir == "" || strings.Contains(dstDir, "-") {
			log.Fatal("destination dir cannot be empty or begins with '-'(dash)")
		}

		// select storages for download
		storages := make([]Downloader, 0)
		if g, err := cmd.Flags().GetBool(googelDrive); err == nil && g {
			storages = append(storages, gdrive.NewGDrive(filepath.Join(resourceDir, gConfig)))
		} else if err != nil {
			log.Println(err) //TODO show an error with logrus but not fail
		}

		if y, err := cmd.Flags().GetBool(yandexDisk); err == nil && y {
			storages = append(storages, yadisk.NewYaDisk(filepath.Join(resourceDir, yaConfig)))
		} else if err != nil {
			log.Println(err)
		}

		if t, err := cmd.Flags().GetBool(telega); err == nil && t {
			storages = append(storages, telegram.NewTelegram(filepath.Join(resourceDir, tgConfig)))
		} else if err != nil {
			log.Println(err)
		}

		for _, storage := range storages {
			fmt.Printf("Загружаю архив из %s\n", storage.Extension())

			fileId, err := fileIdByExt(rr, storage.Extension())
			if err != nil {
				log.Printf("cannot get last file id for %s, %v\n", storage.Extension(), err)
			}

			if err := storage.DownLoadFile(fileId, dstDir); err != nil {
				log.Printf("error occured while downloading %s archive to %s, %v\n", storage.Extension(), dstDir, err)
				continue //not fail because other storages may work propely
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(retrieveCmd)
	retrieveCmd.Flags().BoolP(googelDrive, "g", false, "Download last backup archive from Google Drive")
	retrieveCmd.Flags().BoolP(yandexDisk, "y", false, "Download last backup archive from Yandex Disk")
	retrieveCmd.Flags().BoolP(telega, "t", false, "Download last backup archive from Telegram")

	retrieveCmd.Flags().StringP(destinationDir, "d", "", "Result dir with backup archive path")
	retrieveCmd.MarkFlagRequired(destinationDir)

	retrieveCmd.Flags().Bool(csv, false, "Use csv-file to read uploaded archives data")
}

func fileIdByExt(rr RecorderRetriever, ext string) (string, error) {
	switch v := rr.(type) {
	case *csv_record.CsvRecorderRetriever:
		file, err := os.Open(csvDataFile)
		if err != nil {
			return "", fmt.Errorf("%v, %v", v, err)
		}

		results, err := rr.Retrieve(file, ext)
		if err != nil {
			return "", fmt.Errorf("%v, %v", v, err)
		}
		return results[0], nil
	}
	return "", fmt.Errorf("record with index %s not found", ext)
}
