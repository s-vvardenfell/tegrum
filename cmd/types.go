package cmd

import "io"

type Uploader interface {
	UploadFile(filename string) (string, error)
}

type Downloader interface {
	DownLoadFile(fileId, dst string) error
}

type Record interface {
	Store(w io.Writer, data []string) error
	Retrieve(r io.Reader, index string) ([]string, error)
}

type DirsToBackup struct {
	Dirs []string `json:"dirs"`
}
