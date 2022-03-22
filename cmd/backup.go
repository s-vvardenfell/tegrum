package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/s-vvardenfell/Backuper/archiver"
	"github.com/s-vvardenfell/tegrum/clouds"
	"github.com/s-vvardenfell/tegrum/email"
	"github.com/s-vvardenfell/tegrum/telegram"
	"github.com/s-vvardenfell/tegrum/types"
	"github.com/spf13/cobra"
)

var (
	dirSrc       string
	dirDst       string
	archiverType string
)

const resources = "W:/Golang/src/Backuper/resources" //TODO specify in common config, the same for other files locations
const tgConfig = "telegram.json"
const gConfig = "credentials.json"
const yaConfig = ""
const emailConfig = "email.json"

var archiveName = ""
var archToRemove = ""

type DirsToBackup struct {
	Dirs []string `json:"dirs"`
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backups files immediately to specified storages",
	Long:  `long descr: backups files immediately`, //REMOVE
	Run: func(cmd *cobra.Command, _ []string) {
		o, _ := cmd.Flags().GetBool("one")
		m, _ := cmd.Flags().GetBool("multiple")
		var arch archiver.ArchiverExtracter

		if archiverType == "zip" {
			arch = &archiver.Zip{}
		} else if archiverType == "tar" {
			arch = &archiver.Tar{}
		} else {
			log.Fatal("Wrong archiver type (<zip> and <tar>(will be .tar.gz) supported)")
		}

		if o {
			archiveDir(arch)
		} else if m {
			archiveDirs(arch)
		} else {
			log.Fatal("Single/multiple file mod not selected (use -o for single file or -m for multiple files listed in config)")
		}

		storages := make([]types.Uploader, 0)

		if g, err := cmd.Flags().GetBool("gdrive"); err == nil && g {
			storages = append(storages, clouds.NewGDrive(filepath.Join(resources, gConfig)))
		}
		if y, err := cmd.Flags().GetBool("yadisk"); err == nil && y {
			storages = append(storages, clouds.NewYaDisk(filepath.Join(resources, yaConfig)))
		}
		if t, err := cmd.Flags().GetBool("telegram"); err == nil && t {
			storages = append(storages, telegram.NewTelegram(filepath.Join(resources, tgConfig)))
		}
		if e, err := cmd.Flags().GetBool("email"); err == nil && e {
			storages = append(storages, email.NewMail(filepath.Join(resources, emailConfig)))
		}

		//TODO сюда можно горутины! сделать бенчмарк
		//TODO сохранять fileId в цикле, обрабатывать тут все ошибки, не дб фатала, тк другие способы отправки могут сработать, если 1 не сработал
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

	backupCmd.Flags().StringVarP(&dirDst, "dstDir", "d", "", "Result path")
	backupCmd.MarkFlagRequired("dstDir")

	backupCmd.Flags().StringVarP(&archiverType, "archiver", "a", "", "Use zip / tar")
	backupCmd.MarkFlagRequired("archiver")
}

// Traverses a list of files from dirSrc and archives it
func archiveDirs(arch archiver.ArchiverExtracter) string {
	f, err := os.Open(dirSrc)
	if err != nil {
		log.Fatal(err)
	}

	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var dirs DirsToBackup
	if err := json.Unmarshal([]byte(byteValue), &dirs); err != nil {
		log.Fatal(err)
	}

	tempDir, err := archiver.TempDir(dirDst)
	if err != nil {
		log.Fatal(err)
	}

	if err := archiver.PackArchives(arch, dirs.Dirs, dirDst, tempDir); err != nil {
		log.Fatal(err)
	}

	//gzipping if tar selected
	switch v := arch.(type) {
	case *archiver.Tar:
		archName := tempDir + "." + archiverType
		if err := archiver.Gzip(archName, dirDst); err != nil {
			log.Fatalf("error gziping file(%v), %v", err, v)
		}
		return archName + ".gz"
	}
	return tempDir + "." + archiverType
}

func archiveDir(arch archiver.ArchiverExtracter) string {
	if err := arch.Archive(dirSrc, dirDst); err != nil {
		log.Fatalf("error while single-file archive processed: %v", err)
	}

	//gzipping if tar selected
	switch v := arch.(type) {
	case *archiver.Tar:
		archName := strings.TrimSuffix(dirSrc, filepath.Ext(dirSrc)) + "." + archiverType
		archName = filepath.Join(dirDst, filepath.Base(archName))
		if err := archiver.Gzip(archName, dirDst); err != nil {
			log.Fatalf("error gziping file(%v), %v", err, v)
		}
		return archName + ".gz"
	}
	return dirDst + "." + archiverType
}
