package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/s-vvardenfell/tegrum/clouds/gdrive"
	"github.com/s-vvardenfell/tegrum/clouds/yadisk"
	"github.com/s-vvardenfell/tegrum/records/csv_record"
	"github.com/s-vvardenfell/tegrum/telegram"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Gets backup archives stored earlier from specified storages",
	Run: func(cmd *cobra.Command, _ []string) {
		//select storage type
		//if not specified, error
		var rr RecorderRetriever
		csvFlag, err := cmd.Flags().GetBool(csv)
		if err != nil {
			logrus.Fatal(err)
		}

		if csvFlag {
			rr = &csv_record.CsvRecorderRetriever{}
		} else {
			logrus.Fatal("storage not specified")
		}

		// getting destination dir
		dstDir, err := cmd.Flags().GetString(destinationDir)
		if err != nil || dstDir == "" || strings.Contains(dstDir, "-") {
			logrus.Fatal("destination dir cannot be empty or begins with '-'(dash)")
		}

		// select storages for download
		storages := make([]Downloader, 0)
		if g, err := cmd.Flags().GetBool(googelDrive); err == nil && g {
			storages = append(storages, gdrive.NewGDrive(filepath.Join(resourceDir, gConfig)))
		} else if err != nil {
			logrus.Warningf("cannot init GDrive, %v", err)
		}

		if y, err := cmd.Flags().GetBool(yandexDisk); err == nil && y {
			storages = append(storages, yadisk.NewYaDisk(filepath.Join(resourceDir, yaConfig)))
		} else if err != nil {
			logrus.Warningf("cannot init YaDisk, %v", err)
		}

		if t, err := cmd.Flags().GetBool(telega); err == nil && t {
			storages = append(storages, telegram.NewTelegram(filepath.Join(resourceDir, tgConfig)))
		} else if err != nil {
			logrus.Warningf("cannot init Telegram, %v", err)
		}

		for _, storage := range storages {
			logrus.Info("Загружаю архив из %s\n", storage.Extension())

			fileId, err := fileIdByExt(rr, storage.Extension())
			if err != nil {
				logrus.Warningf("cannot get last file id for %s, %v\n", storage.Extension(), err)
			}

			if err := storage.DownLoadFile(fileId, dstDir); err != nil {
				logrus.Warningf("error occured while downloading %s archive to %s, %v\n", storage.Extension(), dstDir, err)
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
		defer func() { _ = file.Close() }()

		results, err := rr.Retrieve(file, ext)
		if err != nil {
			return "", fmt.Errorf("%v, %v", v, err)
		}
		return results[0], nil
		// case *otherTypes:
	}
	return "", fmt.Errorf("record with index %s not found", ext)
}
