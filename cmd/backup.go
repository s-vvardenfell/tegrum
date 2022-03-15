package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/s-vvardenfell/Backuper/archiver"
	"github.com/s-vvardenfell/Backuper/clouds"
	"github.com/s-vvardenfell/Backuper/mailing"
	"github.com/spf13/cobra"
)

var (
	dirSrc       string
	dstDir       string
	archiverType string
)

type DirsToBackup struct {
	Dirs []string `json:"dirs"`
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backups files immediately to specified storages",
	Long:  `long descr: backups files immediately`,
	Run: func(cmd *cobra.Command, _ []string) {
		g, _ := cmd.Flags().GetBool("gdrive")
		y, _ := cmd.Flags().GetBool("yadisk")
		t, _ := cmd.Flags().GetBool("telegram")
		e, _ := cmd.Flags().GetBool("email")

		if archiverType == "zip" {
			arch := &archiver.Zip{}
			archiveDirs(arch)

		} else if archiverType == "tar" {
			arch := &archiver.Tar{}
			archiveDirs(arch)

		} else {
			log.Fatal("Wrong archiver type (zip and tar supported)")
		}

		if g {
			fmt.Println("Works gdrive")
			cl := clouds.NewGDrive()
			send(cl, ".gitkeep")
		}

		if y {
			cl := clouds.NewYaDisk()
			send(cl, ".gitkeep")
		}

		if t {
			fmt.Println("Works telegram")
		}

		if e {
			fmt.Println("Works email")
			cl := &mailing.Mail{}
			send(cl, ".gitkeep")

		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().BoolP("gdrive", "g", false, "Upload backup archive to Google Drive")
	backupCmd.Flags().BoolP("yadisk", "y", false, "Upload backup archive to Yandex Disk")
	backupCmd.Flags().BoolP("telegram", "t", false, "Sends backup archive to Telegram chat/channel")
	backupCmd.Flags().BoolP("email", "e", false, "Sends backup archive via email")

	backupCmd.Flags().StringVarP(&dirSrc, "dirSrc", "s", "", "Config path")
	backupCmd.MarkFlagRequired("dirSrc")

	backupCmd.Flags().StringVarP(&dstDir, "dstDir", "d", "", "Result path")
	backupCmd.MarkFlagRequired("dstDir")

	backupCmd.Flags().StringVarP(&archiverType, "archiver", "a", "", "Use zip / tar")
	backupCmd.MarkFlagRequired("archiver")
}

func archiveDirs(arch archiver.ArchiverExtracter) {
	f, err := os.Open(dirSrc)

	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatal(err)
	}

	var dirs DirsToBackup
	json.Unmarshal([]byte(byteValue), &dirs)

	for _, dir := range dirs.Dirs {
		arch.Archive(dir, dstDir+"/"+filepath.Base(dir)+".zip")
	}
}

//скорее всего надо внутри функции сделать цикл по архивам в результ-папке, собрать в 1 архив, добавить дату и отправить
func send(cl clouds.Uploader, filename string) {
	cl.UploadFile(filename)
}
