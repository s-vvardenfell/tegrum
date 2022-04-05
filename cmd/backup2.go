package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/s-vvardenfell/tegrum/archiver"
	"github.com/s-vvardenfell/tegrum/archiver/taring"
	"github.com/s-vvardenfell/tegrum/archiver/zipping"
	"github.com/spf13/cobra"
)

// var (
// 	srcDir string
// 	dstDir string
// )

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backups files immediately to specified storages",
	Long:  `long descr: backups files immediately`, //TODO examples
	Run: func(cmd *cobra.Command, _ []string) {

		//TODO error if two selected
		if o, err := cmd.Flags().GetBool("one"); err == nil && o {
			fmt.Printf("one-file mode\n")
		} else if m, err := cmd.Flags().GetBool("multiple"); err == nil && m {
			fmt.Printf("multi-file mode\n")
		} else {
			log.Fatal("Single/multiple file mode not selected (use -o for single file or -m for multiple files listed in config)")
		}

		var arch archiver.Archiver
		//TODO error if two selected
		if tar, err := cmd.Flags().GetBool("tar"); err == nil && tar {
			arch = &taring.Tar{}
			fmt.Printf("tar mode, %T\n", arch)
		} else if zip, err := cmd.Flags().GetBool("zip"); err == nil && zip {
			arch = &zipping.Zip{}
			fmt.Printf("zip mode, %T\n", arch)
		} else {
			log.Fatal("Tar/zip mode not selected (use --zip for zip-archiving or --tar for tar + gzip)")
		}

		srcDir, err := cmd.Flags().GetString("srcDir") //fatal if is empty
		if err != nil || srcDir == "" || strings.Contains(srcDir, "-") {
			log.Fatal("source dir cannot be empty or begins with '-'(dash)")
		}
		dstDir, err := cmd.Flags().GetString("dstDir")
		if err != nil || dstDir == "" || strings.Contains(dstDir, "-") {
			log.Fatal("destination dir cannot be empty or begins with '-'(dash)")
		}
		fmt.Printf("srcDir is %s\n", srcDir)
		fmt.Printf("dstDir is %s\n", dstDir)

		storages := make([]Uploader, 0)
		if g, err := cmd.Flags().GetBool("gdrive"); err == nil && g { //TODO ERRORS HANDLING - IF err !=nil?
			fmt.Println("Selected gdrive", storages)
			// storages = append(storages, clouds.NewGDrive(filepath.Join(resources, gConfig)))
		} else if err != nil {
			log.Println(err) //show an error with logrus but not fail
		}

		if y, err := cmd.Flags().GetBool("yadisk"); err == nil && y {
			fmt.Println("Selected yadisk", storages)
			// storages = append(storages, clouds.NewYaDisk(filepath.Join(resources, yaConfig)))
		} else if err != nil {
			log.Println(err)
		}

		if t, err := cmd.Flags().GetBool("telegram"); err == nil && t {
			fmt.Println("Selected telegram", storages)
			// storages = append(storages, telegram.NewTelegram(filepath.Join(resources, tgConfig)))
		} else if err != nil {
			log.Println(err)
		}

		if e, err := cmd.Flags().GetBool("email"); err == nil && e {
			fmt.Println("Selected email", storages)
			// storages = append(storages, email.NewMail(filepath.Join(resources, emailConfig)))
		} else if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().BoolP("gdrive", "g", false, "Upload backup archive to Google Drive")
	backupCmd.Flags().BoolP("yadisk", "y", false, "Upload backup archive to Yandex Disk")
	backupCmd.Flags().BoolP("telegram", "t", false, "Sends backup archive to Telegram chat/channel")
	backupCmd.Flags().BoolP("email", "e", false, "Sends backup archive via email")

	backupCmd.Flags().BoolP("one", "o", false, "One file from arg")
	backupCmd.Flags().BoolP("multiple", "m", false, "Multiple files listed in *.json-arg file")

	backupCmd.Flags().StringP("srcDir", "s", "", "File to backup path")
	backupCmd.MarkFlagRequired("srcDir")
	backupCmd.Flags().StringP("dstDir", "d", "", "Result dir with backup archive path")
	backupCmd.MarkFlagRequired("dstDir")

	// with global variables
	// backupCmd.Flags().StringVarP(&srcDir, "srcDir", "s", "", "File to backup path")
	// backupCmd.MarkFlagRequired("dirSrc")
	// backupCmd.Flags().StringVarP(&dstDir, "dstDir", "d", "", "Result dir with backup archive path")
	// backupCmd.MarkFlagRequired("dstDir")

	backupCmd.Flags().Bool("tar", false, "Use tar/gz")
	backupCmd.Flags().Bool("zip", false, "Use zip")
}
