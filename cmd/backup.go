package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/s-vvardenfell/Backuper/archiver"
	"github.com/s-vvardenfell/Backuper/clouds"
	"github.com/s-vvardenfell/Backuper/email"
	"github.com/s-vvardenfell/Backuper/telegram"
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
		o, _ := cmd.Flags().GetBool("one")
		m, _ := cmd.Flags().GetBool("multiple")
		var arch archiver.ArchiverExtracter

		if archiverType == "zip" {
			arch = &archiver.Zip{}
		} else if archiverType == "tar" {
			arch = &archiver.Tar{}
		} else {
			log.Fatal("Wrong archiver type (zip and tar supported)")
		}

		if o {
			if err := arch.Archive(dirSrc, dstDir); err != nil {
				log.Fatalf("error while single-file archive processed: %v", err)
			}
		} else if m {
			archiveDirs(arch)
		} else {
			log.Fatal("Single/multiple file mod not selected (use -o for single file or -m for multiple files listed in config)")
		}

		g, _ := cmd.Flags().GetBool("gdrive") //TODO обработка ошибок
		y, _ := cmd.Flags().GetBool("yadisk")
		t, _ := cmd.Flags().GetBool("telegram")
		e, _ := cmd.Flags().GetBool("email")

		storages := make([]clouds.Uploader, 0)

		if g {
			storages = append(storages, clouds.NewGDrive())
		}

		if y {
			storages = append(storages, clouds.NewYaDisk())
		}

		if t {
			storages = append(storages, &telegram.Telegram{})
		}

		if e {
			storages = append(storages, &email.Mail{})
		}

		//TODO сюда можно горутины! сделать бенчмарк
		for i := range storages {
			// storages[i].UploadFile("resources/map.json.zip")
			fmt.Printf("Загружаю в %T", storages[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().BoolP("gdrive", "g", false, "Upload backup archive to Google Drive")
	backupCmd.Flags().BoolP("yadisk", "y", false, "Upload backup archive to Yandex Disk")
	backupCmd.Flags().BoolP("telegram", "t", false, "Sends backup archive to Telegram chat/channel")
	backupCmd.Flags().BoolP("email", "e", false, "Sends backup archive via email")

	backupCmd.Flags().BoolP("one", "o", true, "One file from flag arg")
	backupCmd.Flags().BoolP("multiple", "m", false, "Multiple files listed in *.json file")

	backupCmd.Flags().StringVarP(&dirSrc, "dirSrc", "s", "", "Config path")
	backupCmd.MarkFlagRequired("dirSrc")

	backupCmd.Flags().StringVarP(&dstDir, "dstDir", "d", "", "Result path")
	backupCmd.MarkFlagRequired("dstDir")

	backupCmd.Flags().StringVarP(&archiverType, "archiver", "a", "", "Use zip / tar")
	backupCmd.MarkFlagRequired("archiver")
}

//TODO исп-ть type assertion или что-то такое чтобы определить тип архиватора и выполнить gzip для tar
//Объединить неск архивов в один
// Traverses a list of files from dirSrc and archives it
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
		arch.Archive(dir, dstDir)
	}
}
