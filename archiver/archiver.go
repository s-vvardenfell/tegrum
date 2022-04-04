package archiver

import (
	"os"
	"path/filepath"
	"time"

	"github.com/s-vvardenfell/tegrum/cmd"
)

func TempDir(dst string) (string, error) {
	archiveDir := time.Now().Format("02-01-2006_15-04-05")
	p := filepath.Join(dst, archiveDir)
	err := os.Mkdir(p, 0644)
	if err != nil {
		return "", err
	}
	return p, nil
}

func PackArchives(a cmd.ArchiverExtracter, dirList []string, dst, target string) error {
	for _, dir := range dirList {
		if err := a.Archive(dir, target); err != nil {
			return err
		}
	}

	if err := a.Archive(target, dst); err != nil {
		return err
	}

	if err := os.RemoveAll(target); err != nil {
		return err
	}
	return nil
}
