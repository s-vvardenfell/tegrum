package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/s-vvardenfell/tegrum/archiver/tarring"
	"github.com/s-vvardenfell/tegrum/archiver/zipping"
	"github.com/s-vvardenfell/tegrum/clouds/gdrive"
	"github.com/s-vvardenfell/tegrum/clouds/yadisk"
	"github.com/s-vvardenfell/tegrum/email"
	"github.com/s-vvardenfell/tegrum/records/csv_record"
	"github.com/s-vvardenfell/tegrum/telegram"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// var (
// 	srcDir string
// 	dstDir string
// )

var resourceDir string
var csvDataFile string

// var resourceDir = "W:/Golang/src/Backuper/resources"
// var csvDataFile = "W:/Golang/src/Backuper/resources/data.csv"

const (
	tgConfig    = "telegram.json"
	gConfig     = "credentials.json"
	yaConfig    = "yandex.json"
	emailConfig = "email.json"
)

const (
	googelDrive    = "gdrive"
	yandexDisk     = "yadisk"
	telega         = "telegram"
	mail           = "email"
	oneFileMode    = "one"
	multiFileMode  = "multiple"
	sourceDir      = "srcDir"
	destinationDir = "dstDir"
	tar            = "tar"
	zip            = "zip"
	csv            = "csv"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backups files immediately to specified storages",
	Run: func(cmd *cobra.Command, _ []string) {

		// select archiver type (tar/zip)
		var arch Archiver
		tr, err := cmd.Flags().GetBool(tar)
		if err != nil {
			logrus.Fatal(err)
		}
		zp, err := cmd.Flags().GetBool(zip)
		if err != nil {
			logrus.Fatal(err)
		}

		if tr && zp {
			logrus.Fatalf("cannot use %s and %s at the same time", tar, zip)
		} else if tr {
			arch = tarring.NewTar()
		} else if zp {
			arch = zipping.NewZip()
		} else {
			logrus.Fatal("Tar/zip mode not selected (use --zip for zip-archiving or --tar for tar + gzip)")
		}

		//select storage type
		//if not specified, not store
		var rr RecorderRetriever
		csvFlag, err := cmd.Flags().GetBool(csv)
		if err != nil {
			logrus.Fatal(err)
		}

		if csvFlag {
			rr = &csv_record.CsvRecorderRetriever{}
		} else {
			rr = nil
		}

		// getting source and destination dirs/files
		srcDir, err := cmd.Flags().GetString(sourceDir)
		if err != nil || srcDir == "" || strings.Contains(srcDir, "-") {
			logrus.Fatal("source dir cannot be empty or begins with '-'(dash)")
		}
		dstDir, err := cmd.Flags().GetString(destinationDir)
		if err != nil || dstDir == "" || strings.Contains(dstDir, "-") {
			logrus.Fatal("destination dir cannot be empty or begins with '-'(dash)")
		}

		// getting one- or multi-file mode
		o, err := cmd.Flags().GetBool(oneFileMode)
		if err != nil {
			logrus.Fatal(err)
		}
		m, err := cmd.Flags().GetBool(multiFileMode)
		if err != nil {
			logrus.Fatal(err)
		}

		var archiveName string
		if o && m {
			logrus.Fatalf("cannot use %s and %s file modes at the same time", oneFileMode, multiFileMode)
		} else if o {
			archiveName, err = arch.Archive(srcDir, dstDir)
			if err != nil {
				logrus.Fatalf("error during one-file mode archiving, %v", err)
			}
		} else if m {
			archiveName, err = archiveDirs(arch, srcDir, dstDir)
			if err != nil {
				logrus.Fatalf("error during multi-file mode archiving, %v", err)
			}
		} else {
			logrus.Fatal("Single/multiple file mode not selected (use -o for single file or -m for multiple files listed in config)")
		}

		// select storages for upload
		storages := make([]Uploader, 0)
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

		if e, err := cmd.Flags().GetBool(mail); err == nil && e {
			storages = append(storages, email.NewMail(filepath.Join(resourceDir, emailConfig)))
		} else if err != nil {
			logrus.Warningf("cannot init email, %v", err)
		}

		for _, storage := range storages {
			logrus.Infof("Загружаю %s в %s", archiveName, storage.Extension())

			fileId, err := storage.UploadFile(archiveName)
			if err != nil {
				logrus.Warningf("error occured while uploading archive to %T, %v", storage, err)
				continue //not fail because other storages may work propely
			}

			//if some flag like scv is raised
			if rr != nil {
				if err := storeUploadedFilesValues(rr, fileId, storage.Extension()); err != nil {
					logrus.Warning(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	wd, err := os.Getwd()
	if err != nil {
		logrus.Fatal(err)
	}
	resourceDir = filepath.Join(wd, "resources")
	csvDataFile = filepath.Join(wd, "resources/data.csv")

	backupCmd.Flags().BoolP(googelDrive, "g", false, "Upload backup archive to Google Drive")
	backupCmd.Flags().BoolP(yandexDisk, "y", false, "Upload backup archive to Yandex Disk")
	backupCmd.Flags().BoolP(telega, "t", false, "Sends backup archive to Telegram chat/channel")
	backupCmd.Flags().BoolP(mail, "e", false, "Sends backup archive via email")

	backupCmd.Flags().BoolP(oneFileMode, "o", false, "One file from arg")
	backupCmd.Flags().BoolP(multiFileMode, "m", false, "Multiple files listed in *.json-arg file")

	backupCmd.Flags().StringP(sourceDir, "s", "", "File to backup path")
	backupCmd.MarkFlagRequired(sourceDir)
	backupCmd.Flags().StringP(destinationDir, "d", "", "Result dir with backup archive path")
	backupCmd.MarkFlagRequired(destinationDir)

	// with global variables
	// backupCmd.Flags().StringVarP(&srcDir, "srcDir", "s", "", "File to backup path")
	// backupCmd.MarkFlagRequired("dirSrc")
	// backupCmd.Flags().StringVarP(&dstDir, "dstDir", "d", "", "Result dir with backup archive path")
	// backupCmd.MarkFlagRequired("dstDir")

	backupCmd.Flags().Bool(tar, false, "Use tar/gz")
	backupCmd.Flags().Bool(zip, false, "Use zip")
	backupCmd.Flags().Bool(csv, false, "Use csv to store uploaded archives data")
}

// Traverses a list of files from dirSrc to new dir with timestamp name and archives it
func archiveDirs(arch Archiver, srcDir, dstDir string) (string, error) {
	f, err := os.Open(srcDir)
	if err != nil {
		return "", err
	}
	defer func() { _ = f.Close() }()
	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	var dirs DirsToBackup
	if err := json.Unmarshal([]byte(byteValue), &dirs); err != nil {
		return "", err
	}

	tempDir, err := tempDir(dstDir)
	if err != nil {
		return "", err
	}

	for _, dir := range dirs.Dirs {
		if _, err := arch.Archive(dir, tempDir); err != nil {
			return "", err
		}
	}

	if _, err := arch.Archive(tempDir, dstDir); err != nil {
		return "", err
	}

	if err := os.RemoveAll(tempDir); err != nil {
		return "", err
	}
	return tempDir + fmt.Sprintf(".%s", arch.Extension()), nil
}

// store given file id using given RecorderRetriever type
func storeUploadedFilesValues(rr RecorderRetriever, fileId, storageName string) error {
	// if _, ok := rr.(*csv_record.CsvRecorderRetriever); ok {
	// 	file, err := os.OpenFile(csvDataFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer func() { _ = file.Close() }()
	// 	//possible refactor: RecorderRetriever classes method Record() should open it Writers himself
	// 	data := []string{fileId, storageName, time.Now().Format("01.02.2006 15:04:05")}
	// 	if err := rr.Record(file, data); err != nil {
	// 		return err
	// 	}
	// }
	// return nil

	//TODO refactor: RecorderRetriever clasS method Record() should open it Writers himself
	//func body should be only rr.Store()
	switch v := rr.(type) {
	case *csv_record.CsvRecorderRetriever:
		file, err := os.OpenFile(csvDataFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return fmt.Errorf("%v, %v", v, err)
		}
		defer func() { _ = file.Close() }()

		data := []string{fileId, storageName, time.Now().Format("01.02.2006 15:04:05")}
		if err := rr.Record(file, data); err != nil {
			return fmt.Errorf("%v, %v", v, err)
		}
	default:
		return fmt.Errorf("not suitable or not processed RecorderRetriever: %T", rr)
		// case *otherTypes:
	}
	return nil
}

// creates temporary dir to collect all previos archived files from list in -m mode
func tempDir(dst string) (string, error) {
	archiveDir := time.Now().Format("02-01-2006_15-04-05")
	p := filepath.Join(dst, archiveDir)
	err := os.Mkdir(p, 0644)
	if err != nil {
		return "", err
	}
	return p, nil
}
