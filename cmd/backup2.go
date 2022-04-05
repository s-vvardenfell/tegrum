package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/s-vvardenfell/tegrum/archiver"
	"github.com/s-vvardenfell/tegrum/archiver/tarring"
	"github.com/s-vvardenfell/tegrum/archiver/zipping"
	"github.com/spf13/cobra"
)

// var (
// 	srcDir string
// 	dstDir string
// )

const (
	gdrive         = "gdrive"
	yadisk         = "yadisk"
	telegram       = "telegram"
	email          = "email"
	oneFileMode    = "one"
	multiFileMode  = "multiple"
	sourceDir      = "srcDir"
	destinationDir = "dstDir"
	tar            = "tar"
	zip            = "zip"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backups files immediately to specified storages",
	Long:  `long descr: backups files immediately`, //TODO examples
	Run: func(cmd *cobra.Command, _ []string) {

		// select archiver type (tar/zip)
		var arch archiver.Archiver
		tr, err := cmd.Flags().GetBool(tar)
		if err != nil {
			log.Fatal(err)
		}
		zp, err := cmd.Flags().GetBool(zip)
		if err != nil {
			log.Fatal(err)
		}

		if tr && zp {
			log.Fatalf("cannot use %s and %s at the same time\n", tar, zip)
		} else if tr {
			arch = &tarring.Tar{}
			fmt.Printf("tar mode, %T\n", arch) // test output remove
		} else if zp {
			arch = &zipping.Zip{}
			fmt.Printf("zip mode, %T\n", arch) // test output remove
		} else {
			log.Fatal("Tar/zip mode not selected (use --zip for zip-archiving or --tar for tar + gzip)")
		}

		// getting source and destination dirs/files
		srcDir, err := cmd.Flags().GetString(sourceDir)
		if err != nil || srcDir == "" || strings.Contains(srcDir, "-") {
			log.Fatal("source dir cannot be empty or begins with '-'(dash)")
		}
		dstDir, err := cmd.Flags().GetString(destinationDir)
		if err != nil || dstDir == "" || strings.Contains(dstDir, "-") {
			log.Fatal("destination dir cannot be empty or begins with '-'(dash)")
		}
		fmt.Printf("srcDir is %s\n", srcDir) // test output remove
		fmt.Printf("dstDir is %s\n", dstDir) // test output remove

		// getting one- or multi-file mode
		o, err := cmd.Flags().GetBool(oneFileMode)
		if err != nil {
			log.Fatal(err)
		}
		m, err := cmd.Flags().GetBool(multiFileMode)
		if err != nil {
			log.Fatal(err)
		}

		if o && m {
			log.Fatalf("cannot use %s and %s file modes at the same time\n", oneFileMode, multiFileMode)
		} else if o {
			fmt.Printf("one-file mode\n") // test output remove
		} else if m {
			fmt.Printf("multi-file mode\n") // test output remove
		} else {
			log.Fatal("Single/multiple file mode not selected (use -o for single file or -m for multiple files listed in config)")
		}

		// select storages for upload
		storages := make([]Uploader, 0)
		if g, err := cmd.Flags().GetBool(gdrive); err == nil && g {
			fmt.Println("Selected gdrive", storages) // test output remove
			// storages = append(storages, clouds.NewGDrive(filepath.Join(resources, gConfig)))
		} else if err != nil {
			log.Println(err) //show an error with logrus but not fail
		}

		if y, err := cmd.Flags().GetBool(yadisk); err == nil && y {
			fmt.Println("Selected yadisk", storages) // test output remove
			// storages = append(storages, clouds.NewYaDisk(filepath.Join(resources, yaConfig)))
		} else if err != nil {
			log.Println(err)
		}

		if t, err := cmd.Flags().GetBool(telegram); err == nil && t {
			fmt.Println("Selected telegram", storages) // test output remove
			// storages = append(storages, telegram.NewTelegram(filepath.Join(resources, tgConfig)))
		} else if err != nil {
			log.Println(err)
		}

		if e, err := cmd.Flags().GetBool(email); err == nil && e {
			fmt.Println("Selected email", storages) // test output remove
			// storages = append(storages, email.NewMail(filepath.Join(resources, emailConfig)))
		} else if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().BoolP(gdrive, "g", false, "Upload backup archive to Google Drive")
	backupCmd.Flags().BoolP(yadisk, "y", false, "Upload backup archive to Yandex Disk")
	backupCmd.Flags().BoolP(telegram, "t", false, "Sends backup archive to Telegram chat/channel")
	backupCmd.Flags().BoolP(email, "e", false, "Sends backup archive via email")

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
}
