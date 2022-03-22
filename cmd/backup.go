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
	"github.com/s-vvardenfell/Backuper/clouds"
	"github.com/s-vvardenfell/Backuper/email"
	"github.com/s-vvardenfell/Backuper/telegram"
	"github.com/spf13/cobra"
)

var (
	dirSrc       string
	dirDst       string
	archiverType string
)

const resources = "W:/Golang/src/Backuper/resources"
const tgConfig = "telegram.json"
const gConfig = "credentials.json"
const yaConfig = ""
const emailConfig = "email.json"
const gdrive = "gdrive"

var archiveName = ""

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
			archiveName = archiveDir(arch)
		} else if m {
			archiveName = archiveDirs(arch)
		} else {
			log.Fatal("Single/multiple file mod not selected (use -o for single file or -m for multiple files listed in config)")
		}

		g, _ := cmd.Flags().GetBool("gdrive") //TODO обработка ошибок / const gdrive = "gdrive" - протестить
		y, _ := cmd.Flags().GetBool("yadisk")
		t, _ := cmd.Flags().GetBool("telegram")
		e, _ := cmd.Flags().GetBool("email")

		storages := make([]clouds.Uploader, 0)

		if g {
			storages = append(storages, clouds.NewGDrive(filepath.Join(resources, gConfig)))
		}

		if y {
			storages = append(storages, clouds.NewYaDisk(filepath.Join(resources, yaConfig)))
		}

		if t {
			t := telegram.NewTelegram(filepath.Join(resources, tgConfig))
			fileId, err := t.UploadFile(archiveName)
			if err != nil {
				log.Printf("%v, new fileId(%s) not stored\n", err, fileId)
			}
			//TODO store new fileId
		}

		if e {
			e := email.NewMail(filepath.Join(resources, emailConfig))
			if err := e.SendMsgWithAttachment(archiveName); err != nil {
				log.Printf("email sending error, %v", err.Error())
			}
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
