package tarring

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const ext = "tar.gz"

type Tar struct {
	extension string
}

func NewTar() *Tar {
	return &Tar{ext}
}

func (t *Tar) Extension() string {
	return t.extension
}

func (t *Tar) Archive(source, target string) (string, error) {
	filename := filepath.Base(source)
	target = filepath.Join(target, fmt.Sprintf("%s.%s", strings.TrimSuffix(filename, filepath.Ext(filename)), t.Extension()))
	tarfile, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer tarfile.Close()

	gzw := gzip.NewWriter(tarfile)
	defer gzw.Close()

	tarball := tar.NewWriter(gzw)
	defer tarball.Close()

	info, err := os.Stat(source)
	if err != nil {
		return "", err
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	return target, filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			}

			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})
}

func (t *Tar) Extract(tarball, target string) error {
	reader, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer reader.Close()
	// tarReader := tar.NewReader(reader)

	gzw, _ := gzip.NewReader(reader)
	defer gzw.Close()

	tarReader := tar.NewReader(gzw)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}
	return nil
}
