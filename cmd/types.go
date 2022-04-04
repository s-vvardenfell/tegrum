package cmd

import "io"

type Uploader interface {
	UploadFile(filename string) (string, error)
}

type Downloader interface {
	DownLoadFile(fileId, dst string) error
}

type ArchiverExtracter interface {
	Archive(source, target string) error
	Extract(archive, target string) error
}

type Record interface {
	Store(w io.Writer, data []string) error
	Retrieve(r io.Reader, index string) ([]string, error)
}
