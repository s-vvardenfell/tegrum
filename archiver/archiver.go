package archiver

import (
	"os"
	"path/filepath"
	"time"
)

type Archiver interface {
	Archive(source, dst string) (string, error)
	Extension() string
}

type Extracter interface {
	Extract(archive, dst string) error
	Extension() string
}

func TempDir(dst string) (string, error) {
	archiveDir := time.Now().Format("02-01-2006_15-04-05")
	p := filepath.Join(dst, archiveDir)
	err := os.Mkdir(p, 0644)
	if err != nil {
		return "", err
	}
	return p, nil
}

// arch, dirs.Dirs, dirDst, tempDir
// func PackArchives(a Archiver, dirList []string, dst, target string) error {
// 	for _, dir := range dirList {
// 		if err := a.Archive(dir, target); err != nil {
// 			return err
// 		}
// 	}

// 	if err := a.Archive(target, dst); err != nil {
// 		return err
// 	}

// 	if err := os.RemoveAll(target); err != nil {
// 		return err
// 	}
// 	return nil
// }
