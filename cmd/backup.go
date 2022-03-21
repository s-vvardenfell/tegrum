package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/s-vvardenfell/Backuper/archiver"
	"github.com/s-vvardenfell/Backuper/clouds"
	"github.com/spf13/cobra"
)

var (
	dirSrc       string
	dstDir       string
	archiverType string
)

const tgConfig = "resources/telegram.json"
const gConfig = "resources/credentials.json"
const yaConfig = ""
const emailConfig = "resources/email.json"

const gdrive = "gdrive"

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
			archiveDir(arch)
		} else if m {
			archiveDirs(arch)
		} else {
			log.Fatal("Single/multiple file mod not selected (use -o for single file or -m for multiple files listed in config)")
		}

		g, _ := cmd.Flags().GetBool("gdrive") //TODO обработка ошибок / const gdrive = "gdrive" - протестить
		y, _ := cmd.Flags().GetBool("yadisk")
		t, _ := cmd.Flags().GetBool("telegram")
		e, _ := cmd.Flags().GetBool("email")

		storages := make([]clouds.Uploader, 0)

		if g {
			storages = append(storages, clouds.NewGDrive(gConfig))
		}

		if y {
			storages = append(storages, clouds.NewYaDisk(yaConfig))
		}

		if t {
			fmt.Println("Tg works")
			// storages = append(storages, telegram.NewTelegram(tgConfig))
		}

		if e {
			fmt.Println("Email works")
			// storages = append(storages, email.NewMail(emailConfig))
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

	backupCmd.Flags().BoolP("one", "o", false, "One file from flag arg")
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

	archiveDir := time.Now().Format("02.Jan.2006_15:04:05_backup")
	for _, dir := range dirs.Dirs {
		arch.Archive(dir, dstDir+"/"+archiveDir+"/")
	}
}

func archiveDir(arch archiver.ArchiverExtracter) {
	if err := arch.Archive(dirSrc, dstDir); err != nil {
		log.Fatalf("error while single-file archive processed: %v", err)
	}
}
