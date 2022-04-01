package archiver

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	ZIP = "zip"
	TAR = "tar.gz"
	GZ  = "gz"
)

func Archive(archType, source, target string) error {
	filename := filepath.Base(source)

	switch archType {
	case ZIP:
		fmt.Println("zip choosen")
		target = filepath.Join(target, fmt.Sprintf("%s.%s", strings.TrimSuffix(filename, filepath.Ext(filename)), ZIP))

		file, err := os.Create(target)
		if err != nil {
			return err
		}
		defer file.Close()

		archiver := tar.NewWriter(file)
		defer archiver.Close()

	case TAR:
		fmt.Println("tar choosen")
		target = filepath.Join(target, fmt.Sprintf("%s.%s", strings.TrimSuffix(filename, filepath.Ext(filename)), TAR))

		file, err := os.Create(target)
		if err != nil {
			return err
		}
		defer file.Close()

		archiver := tar.NewWriter(file)
		defer archiver.Close()

	default:
		fmt.Println("no such type")
		return fmt.Errorf("no such archiver type, %s", archType)
	}

	fmt.Println(target)

	// info, err := os.Stat(source)
	// if err != nil {
	// 	return nil
	// }

	// var baseDir string
	// if info.IsDir() {
	// 	baseDir = filepath.Base(source)
	// }

	return nil
}

func Extract(archiver io.WriteCloser, source, target string) error {
	return nil
}
